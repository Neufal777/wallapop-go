package src

import (
	"github.com/walla-chollo/src/downloader"
	"github.com/walla-chollo/src/wallapop"
)

type Engine struct {
	Downloader downloader.Downloader `json:"downloader"`
	Content    wallapop.Content      `json:"content"`
}

func NewEngine() *Engine {
	return &Engine{}
}

func (eng *Engine) SetDownloader(downloader downloader.Downloader) *Engine {
	eng.Downloader = downloader
	return eng
}

func (eng *Engine) SetContent(content wallapop.Content) *Engine {
	eng.Content = content
	return eng
}
