package util

type Renderable interface {
	Draw()
	Hide()
	Interrupt() error
}
