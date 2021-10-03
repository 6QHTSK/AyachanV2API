package Models

type RegularType int

const (
	RegularTypeUnknown RegularType = iota
	RegularTypeRegular
	RegularTypeIrregular
)

type DifficultyDescription int

const (
	DifficultyUnknown DifficultyDescription = iota // 该项难度未知
	DifficultyLow                                  // 该项难度偏低
	DifficultyNormal                               // 该项难度正常
	DifficultyHigh                                 // 该项难度偏高
)

type BpmInfo struct {
	BPMLow  float64
	BPMHigh float64
	MainBPM float64
}

type IrregularInfo struct {
	Irregular     RegularType // 存在多压/交叉（出张）0 失败 1 标准 2 非标准
	IrregularInfo string      // 无法分析的第一个错误情况
}

type NoteCount struct {
	SPRhythm    bool
	Single      int
	Flick       int
	SlideStart  int
	SlideTick   int
	SlideEnd    int
	SlideHidden int
	Direction   struct {
		Total int
		Left  int
		Right int
	}
}

type Distribution struct {
	Note []float64
	Hit  []float64
}

// MapInfoBasic 将会放入数据库存档的数据部分
type MapInfoBasic struct {
	BpmInfo
	IrregularInfo
	TotalNote int
	TotalTime float64
	TotalNPS  float64
}

// MapInfoStandard 基础部分，不要求正常谱面
type MapInfoStandard struct {
	MapInfoBasic

	TotalHitNote int
	MaxScreenNPS float64
	TotalHPS     float64

	NoteCount    NoteCount
	Distribution Distribution
}

// MapInfoExtend 扩展部分，要求正常谱面，非正常时为nil
type MapInfoExtend struct {
	LeftPercent       float64
	MaxSpeed          float64
	FingerMaxHPS      float64
	FlickNoteInterval float64
	NoteFlickInterval float64
}

// MapDifficultyStandard 基础部分，不要求正常谱面
type MapDifficultyStandard struct {
	TotalNPS            float64
	TotalHPS            float64
	MaxScreenNPS        float64
	Difficulty          float64
	BlueWhiteDifficulty float64
}

// MapDifficultyExtend 扩展部分，要求正常谱面，非正常时为nil
type MapDifficultyExtend struct {
	MaxSpeed          DifficultyDescription
	FingerMaxHPS      DifficultyDescription
	FlickNoteInterval DifficultyDescription
	NoteFlickInterval DifficultyDescription
}

type MapInfo struct {
	MapInfoStandard       MapInfoStandard
	MapInfoExtend         interface{} // MapInfoExtend
	MapDifficultyStandard MapDifficultyStandard
	MapDifficultyExtend   interface{} // MapDifficultyExtend
}
