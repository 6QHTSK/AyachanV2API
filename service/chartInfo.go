package service

import (
	"container/heap"
	"github.com/6QHTSK/ayachan/manager"
	"github.com/6QHTSK/ayachan/model"
	"github.com/6QHTSK/ayachan/model/ayachanChart"
	"github.com/6QHTSK/ayachan/utils"
	"math"
)

// basicMetricsGetter 获得最基础的谱面信息(除了Irregular项)、Hit总数、HPS
func basicMetricsGetter(Chart ayachanChart.AyachanChart) (info model.ChartMetricsBasic, TotalHitCount int, TotalHPS float64, BPMInfo model.BpmInfo, err error) {
	var BPMList map[float64]float64
	var firstNoteTime float64
	BPMList = make(map[float64]float64)
	noteFlag := true // 检查前置区间内是否无note
	BPMInfo.BPMLow = math.MaxFloat64
	BPMInfo.BPMHigh = -1.0
	currentBPM := 120.0
	currentBPMStartTime := 0.0
	MainBPMTime := -1.0
	for _, note := range Chart {
		switch note.Type {
		case ayachanChart.NoteTypeSingle:
			if noteFlag {
				firstNoteTime = note.Time
				noteFlag = false
			}
			info.TotalNote++
			TotalHitCount++
			if note.Flick && note.Direction != 0 {
				info.SPRhythm = true
			}
		case ayachanChart.NoteTypeSlide:
			if noteFlag {
				firstNoteTime = note.Time
				noteFlag = false
			}
			if !note.Hidden {
				info.TotalNote++
				if note.Status == ayachanChart.SlideStart {
					TotalHitCount++
				}
			} else {
				info.SPRhythm = true
			}
		case ayachanChart.NoteTypeBpm:
			if !noteFlag {
				BPMInfo.BPMLow = math.Min(BPMInfo.BPMLow, currentBPM)
				BPMInfo.BPMHigh = math.Max(BPMInfo.BPMHigh, currentBPM)
				BPMList[currentBPM] += note.Time - currentBPMStartTime
				if BPMList[currentBPM] > MainBPMTime {
					BPMInfo.MainBPM = currentBPM
					MainBPMTime = BPMList[currentBPM]
				}
			}
			currentBPM = note.BPM
			currentBPMStartTime = note.Time
		}
	}
	// Append最后一个BPM数据
	BPMInfo.BPMLow = math.Min(BPMInfo.BPMLow, currentBPM)
	BPMInfo.BPMHigh = math.Max(BPMInfo.BPMHigh, currentBPM)
	BPMList[currentBPM] += Chart[len(Chart)-1].Time - currentBPMStartTime
	if BPMList[currentBPM] > MainBPMTime {
		BPMInfo.MainBPM = currentBPM
		MainBPMTime = BPMList[currentBPM]
	}
	info.TotalTime = math.Max(20.0, Chart[len(Chart)-1].Time-firstNoteTime)
	info.TotalNPS = float64(info.TotalNote) / info.TotalTime
	TotalHPS = float64(TotalHitCount) / info.TotalTime
	return info, TotalHitCount, TotalHPS, BPMInfo, nil
}

// noteCounter 谱面计数器，谱面Note计数
func noteCounter(Chart ayachanChart.AyachanChart) (NoteCount model.NoteCount) {
	for _, note := range Chart {
		switch note.Type {
		case ayachanChart.NoteTypeSingle:
			if note.Flick {
				if note.Direction > 0 {
					NoteCount.DirectionRight++
				} else if note.Direction < 0 {
					NoteCount.DirectionLeft++
				} else {
					NoteCount.Flick++
				}
			} else {
				NoteCount.Single++
			}
		case ayachanChart.NoteTypeSlide:
			switch note.Status {
			case ayachanChart.SlideStart:
				NoteCount.SlideStart++
			case ayachanChart.SlideTick:
				if note.Hidden {
					NoteCount.SlideHidden++
				} else {
					NoteCount.SlideTick++
				}
			case ayachanChart.SlideEnd:
				if note.Flick {
					NoteCount.SlideFlick++
				} else {
					NoteCount.SlideEnd++
				}
			}
		}
	}
	return NoteCount
}

// noteDistribution 谱面分布计数器，MaxScreenNPS
func noteDistribution(Chart ayachanChart.AyachanChart, totalTime float64) (MaxScreenNPS float64, Distribution model.Distribution) {
	Distribution.Note = make([]int, int(math.Ceil(totalTime+0.01)))
	Distribution.Hit = make([]int, int(math.Ceil(totalTime+0.01)))
	var MaxScreenNPSHeap utils.Float64Heap
	heap.Init(&MaxScreenNPSHeap)
	for _, note := range Chart {
		switch note.Type {
		case ayachanChart.NoteTypeSingle:
			Distribution.Note[int(math.Floor(note.Time))]++
			Distribution.Hit[int(math.Floor(note.Time))]++
		case ayachanChart.NoteTypeSlide:
			if note.Status == ayachanChart.SlideStart {
				Distribution.Note[int(math.Floor(note.Time))]++
				Distribution.Hit[int(math.Floor(note.Time))]++
			} else {
				if !note.Hidden {
					Distribution.Note[int(math.Floor(note.Time))]++
				}
			}
		}
	}
	for _, item := range Distribution.Note {
		//MaxScreenNPS = math.Max(MaxScreenNPS, float64(item))
		heap.Push(&MaxScreenNPSHeap, float64(item))
	}
	return MaxScreenNPSHeap.GetTopRankAverage(), Distribution
}

// reciprocal 去除0的倒数为inf的问题
func reciprocal(num float64) (r float64) {
	if num == 0.0 {
		return r
	}
	return 1.0 / num
}

// extendMetricsGetter 获得扩展谱面信息
func extendMetricsGetter(ParsedChart ayachanChart.ParsedChart) *model.ChartMetricsExtend {
	var leftCount, RightCount int
	var MaxSpeed, FingerMaxHPSLeft, FingerMaxHPSRight, FlickNoteInterval, NoteFlickInterval utils.Float64Heap
	heap.Init(&MaxSpeed)
	heap.Init(&FingerMaxHPSLeft)
	heap.Init(&FingerMaxHPSRight)
	heap.Init(&FlickNoteInterval)
	heap.Init(&NoteFlickInterval)
	var lastLeftHandHit *ayachanChart.ParsedNote
	var lastRightHandHit *ayachanChart.ParsedNote
	for i, note := range ParsedChart {
		if note.Hand == ayachanChart.LeftHand {
			if lastLeftHandHit != nil && !(note.Type == ayachanChart.NoteTypeSlide && note.Status != ayachanChart.SlideStart) {
				singleHPS := reciprocal(note.GetInterval(lastLeftHandHit))
				heap.Push(&FingerMaxHPSLeft, singleHPS)
			}
			leftCount++
			lastLeftHandHit = &note
		} else {
			if lastRightHandHit != nil && !(note.Type == ayachanChart.NoteTypeSlide && note.Status != ayachanChart.SlideStart) {
				singleHPS := reciprocal(note.GetInterval(lastRightHandHit))
				heap.Push(&FingerMaxHPSRight, singleHPS)
			}
			RightCount++
			lastRightHandHit = &note
		}
		if i != 0 {
			heap.Push(&MaxSpeed, math.Abs(note.GetGapFront())/note.GetIntervalFront())
		}
		if note.Type == ayachanChart.NoteTypeSingle && note.Flick {
			heap.Push(&FlickNoteInterval, reciprocal(note.GetIntervalBack()))
			heap.Push(&NoteFlickInterval, reciprocal(note.GetIntervalFront()))
		}
	}
	totalCount := leftCount + RightCount
	return &model.ChartMetricsExtend{
		LeftPercent:       float64(leftCount) / float64(totalCount),
		FingerMaxHPS:      math.Max(FingerMaxHPSLeft.GetTopRankAverage(), FingerMaxHPSRight.GetTopRankAverage()),
		MaxSpeed:          MaxSpeed.GetTopRankAverage(),
		FlickNoteInterval: FlickNoteInterval.GetTopRankAverage(),
		NoteFlickInterval: NoteFlickInterval.GetTopRankAverage(),
	}
}

// standardMetricsGetter 获得标准谱面信息,除了Irregular项
func standardMetricsGetter(Chart ayachanChart.AyachanChart) (StandardMetrics model.ChartMetricsStandard) {
	StandardMetrics.ChartMetricsBasic, StandardMetrics.TotalHitNote, StandardMetrics.TotalHPS, StandardMetrics.BpmInfo, _ = basicMetricsGetter(Chart)
	StandardMetrics.NoteCount = noteCounter(Chart)
	StandardMetrics.MaxScreenNPS, StandardMetrics.Distribution = noteDistribution(Chart, Chart[len(Chart)-1].Time)
	return StandardMetrics
}

// ChartAnalyze 获得全部的谱面信息
func ChartAnalyze(inputChart model.AnalyzableChart, diff int) (model.ChartMetrics, error) {
	suc, err := inputChart.ChartCheck()
	if !suc {
		return model.ChartMetrics{}, err
	}
	Chart := inputChart.DecodeToAyachan()
	standardMetrics := standardMetricsGetter(Chart)
	standardDifficulty, diff := standardDifficultyGetter(standardMetrics, diff)

	var ParsedChart ayachanChart.ParsedChart
	ParsedChart, standardMetrics.ChartMetricsBasic.IrregularInfo = parseChart(Chart)

	var extendMetrics *model.ChartMetricsExtend
	var extendDifficulty *model.ChartDifficultyExtend
	if standardMetrics.Irregular == model.RegularTypeRegular {
		extendMetrics = extendMetricsGetter(ParsedChart)
		extendDifficulty = extendDifficultyGetter(*extendMetrics, diff, standardDifficulty.Difficulty)
	}

	return model.ChartMetrics{
		Metrics:          &standardMetrics,
		MetricsExtend:    extendMetrics,
		Difficulty:       &standardDifficulty,
		DifficultyExtend: extendDifficulty,
	}, nil
}

func BandoriChartAnalyze(chartID int, diff int) (MapInfo model.ChartMetrics, err error) {
	BestdoriV2Map, err := manager.GetBandoriChart(chartID, diff)
	if err != nil {
		return MapInfo, err
	}
	return ChartAnalyze(BestdoriV2Map, diff)
}

func BestdoriChartAnalyze(chartID int) (MapInfo model.ChartMetrics, err error) {
	BestdoriV2Map, diff, err := manager.GetBestdoriChart(chartID)
	if err != nil {
		return MapInfo, err
	}
	return ChartAnalyze(BestdoriV2Map, diff)
}
