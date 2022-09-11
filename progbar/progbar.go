package progbar

import (
	"fmt"
	"time"
)

const (
	None = iota
	Percent
	Fraction
	Count
	Label
	Rate
)

type Coord struct {
	X, Y int
}

type ProgBar struct {
	Label    string
	Capacity int
	Count    int
	Modes    []uint
	Dims     Coord
	Started  time.Time
	DrawBar  func([]uint) string
	Erase    func() string
	Offset   Coord
	Spinner  bool
	IsOffset bool
}

func (p *ProgBar) Draw() {
	p.Hide()
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
