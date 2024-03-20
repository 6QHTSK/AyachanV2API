package model

import "github.com/6QHTSK/ayachan/model/ayachanChart"

type AnalyzableChart interface {
	ChartCheck() (bool, error)
	DecodeToAyachan() ayachanChart.AyachanChart
}
