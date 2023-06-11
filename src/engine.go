package src

type Engine struct {
	Downloader Downloader `json:"downloader"`
	Content    Content    `json:"content"`
}

func NewEngine() *Engine {
	return &Engine{}
}

func (eng *Engine) SetDownloader(downloader Downloader) *Engine {
	eng.Downloader = downloader
	return eng
}

func (eng *Engine) SetContent(content Content) *Engine {
	eng.Content = content
	return eng
}
