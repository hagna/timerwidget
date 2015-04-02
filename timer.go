package timerwidget

import (
	"github.com/gizak/termui"
	"time"
	//"log"
	"math"
)

type Widget struct {
	*termui.Gauge
	starttime time.Time
	len       time.Duration
	elapsed   time.Duration
}

func New(t time.Duration, w, h int, label string) *Widget {
	w0 := new(Widget)
	w0.Gauge = termui.NewGauge()
	w0.Percent = 100
	w0.Width = w
	w0.Height = h
	w0.Border.Label = label
	var err error
	w0.len = t
	w0.starttime = time.Now()
	if err != nil {
		panic(err)
	}
	return w0
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}

func (w *Widget) Update() {
	t := time.Since(w.starttime)
	w.elapsed += t
	if w.elapsed >= w.len {
		w.Percent = 0
		return
	}
	p := int(Round(100.0 - float64(w.elapsed*100)/float64(w.len)))
	//log.Println(w.Border.Label, "w.len is", w.len, "w.elapsed is", w.elapsed, "percent is", p)
	w.Percent = p
	w.starttime = time.Now()

}

func (w *Widget) Rewind() {
	w.Restart()
	w.Percent = 100
	w.elapsed = 0
}

func (w *Widget) Restart() {
	w.starttime = time.Now()
}
