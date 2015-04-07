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
	Len       time.Duration
	elapsed   time.Duration
	Increasing bool
}


func (w0 *Widget) setStart() {
	if w0.Increasing {
		w0.Percent = 0
	} else {
		w0.Percent = 100
	}
}

func (w0 *Widget) setEnd() {
	if w0.Increasing {
		w0.Percent = 100
	} else {
		w0.Percent = 0
	}
}

func New(t time.Duration, w, h int, label string) *Widget {
	w0 := new(Widget)
	w0.Gauge = termui.NewGauge()
	w0.setStart()
	w0.Width = w
	w0.Height = h
	w0.Border.Label = label
	w0.Len = t
	w0.starttime = time.Now()
	return w0
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}

func (w *Widget) Update() {
	if w.elapsed >= w.Len {
		w.setEnd()
		return
	}

	t := time.Since(w.starttime)
	w.elapsed += t
	p := Round(float64(w.elapsed*100)/float64(w.Len))
	//log.Println(w.Border.Label, "w.Len is", w.Len, "w.elapsed is", w.elapsed, "percent is", p)
	if w.Increasing == false {
		w.Percent = int(math.Max(0, 100 - p))
	}  else {
		w.Percent = int(math.Min(100, p))
	}
	w.starttime = time.Now()

}

func (w *Widget) Rewind() {
	w.Retime()
	w.setStart()
	w.elapsed = 0
}

func (w *Widget) Retime() {
	w.starttime = time.Now()
}
