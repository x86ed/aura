package progbar

import (
	"fmt"
	"time"

	"github.com/x86ed/aura/util"
)

const (
	None = iota
	Percent
	Fraction
	Count
	Label
	Rate
)

type ProgBar struct {
	util.RenderObj
	Capacity int
	Count    int
	Modes    []uint
	Started  time.Time
	DrawBar  func([]uint) string
	Erase    func() string
	Spinner  bool
}

func (p *ProgBar) Draw(o ...util.Coord) {
	p.Hide()
	if len(o) > 0 {
		p.IsOffset = true
		p.Offset = o[0]
	}
	fmt.Println(p.DrawBar(p.Modes))
}

func (p *ProgBar) Hide() {
	fmt.Println(p.Erase())
}

func (p *ProgBar) Interrupt() error {
	p.Count = p.Capacity
	return nil
}

func (p *ProgBar) Start() {
	p.Started = time.Now()
}

func (p *ProgBar) ChannelUpdate(c chan int) {
	go func() {
		for p.Capacity < p.Count {
			select {
			case msg := <-c:
				p.Count += msg
				p.Draw()
			}
		}
	}()
}
