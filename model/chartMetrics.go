package model

type RegularType int

const (
	RegularTypeUnknown RegularType = iota
	RegularTypeRegular
	RegularTypeIrregular
)

type DifficultyDescription int

const (
	DifficultyLow    = -1 // 该项难度偏低
	DifficultyNormal = 0  // 该项难度正常
	DifficultyHigh   = 1  // 该项难度偏高
)

type BpmInfo struct {
	BPMLow  float64 `json:"bpm_low"`
	BPMHigh float64 `json:"bpm_high"`
	MainBPM float64 `json:"main_bpm"`
}

type IrregularInfo struct {
	Irregular     RegularType `json:"irregular"`      // 存在多压/交叉（出张）0 失败 1 标准 2 非标准
	IrregularInfo string      `json:"irregular_info"` // 无法分析的第一个错误情况
}

type NoteCount struct {
	Single         int `json:"single"`
	Flick          int `json:"flick"`
	SlideStart     int `json:"slide_start"`
	SlideTick      int `json:"slide_tick"`
	SlideEnd       int `json:"slide_end"`
	SlideFlick     int `json:"slide_flick"`
	SlideHidden    int `json:"slide_hidden"`
	DirectionLeft  int `json:"direction_left"`
	DirectionRight int `json:"direction_right"`
}

type Distribution struct {
	Note []int `json:"note"`
	Hit  []int `json:"hit"`
}

// ChartMetricsBasic 将会放入数据库存档的数据部分
type ChartMetricsBasic struct {
	IrregularInfo
	TotalNote int     `json:"total_note"`
	TotalTime float64 `json:"total_time"`
	TotalNPS  float64 `json:"total_nps"`
	SPRhythm  bool    `json:"sp_rhythm"`
}

// ChartMetricsStandard 基础部分，不要求正常谱面
type ChartMetricsStandard struct {
	ChartMetricsBasic

	BpmInfo
	TotalHitNote int     `json:"total_hit_note"`
	MaxScreenNPS float64 `json:"max_screen_nps"`
	TotalHPS     float64 `json:"total_hps"`

	NoteCount    NoteCount
	Distribution Distribution
}

// ChartMetricsExtend 扩展部分，要求正常谱面，非正常时为nil
type ChartMetricsExtend struct {
	LeftPercent       float64 `json:"left_percent"`
	MaxSpeed          float64 `json:"max_speed"`
	FingerMaxHPS      float64 `json:"finger_max_hps"`
	FlickNoteInterval float64 `json:"flick_note_interval"`
	NoteFlickInterval float64 `json:"note_flick_interval"`
}

// ChartDifficultyStandard 基础部分，不要求正常谱面
type ChartDifficultyStandard struct {
	TotalNPS     float64 `json:"total_nps"`
	TotalHPS     float64 `json:"total_hps"`
	MaxScreenNPS float64 `json:"max_screen_nps"`
	Difficulty   float64 `json:"difficulty"`
}

// ChartDifficultyExtend 扩展部分，要求正常谱面，非正常时为nil
type ChartDifficultyExtend struct {
	MaxSpeed          DifficultyDescription `json:"max_speed"`
	FingerMaxHPS      DifficultyDescription `json:"finger_max_hps"`
	FlickNoteInterval DifficultyDescription `json:"flick_note_interval"`
	NoteFlickInterval DifficultyDescription `json:"note_flick_interval"`
}

type ChartMetrics struct {
	Metrics          *ChartMetricsStandard    `json:"metrics"`
	MetricsExtend    *ChartMetricsExtend      `json:"metrics_extend,omitempty"`
	Difficulty       *ChartDifficultyStandard `json:"difficulty"`
	DifficultyExtend *ChartDifficultyExtend   `json:"difficulty_extend,omitempty"`
}
