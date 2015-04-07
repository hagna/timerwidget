// Copyright 2015 Zack Guo <gizak@icloud.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.


package main

import ("github.com/gizak/termui"
	tm "github.com/nsf/termbox-go"
	"github.com/hagna/timerwidget"
	"time"
	"flag"
	"log"
	"strconv"
	_ "net/http/pprof"
	"net/http"
)

var fuse = flag.Duration("t", 10 * time.Second, "timeout")

type stateargs struct {
	evt chan tm.Event
	widget *timerwidget.Widget
	pause bool
	nums string
}

type stateFn func(*stateargs) stateFn

func nums(s *stateargs) stateFn {
	w := s.widget
	select {
		case e := <-s.evt:
			if e.Type == tm.EventKey {
				switch e.Ch {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						s.nums += string(e.Ch)
						return nums
				}
				switch e.Key {
					case tm.KeyEnter:
						n, err := strconv.Atoi(s.nums)
						if err == nil {
							w.Len = time.Second * time.Duration(n)
							s.nums = ""
							w.Border.Label = w.Len.String()
						} else {
							w.Border.Label = err.Error()
						}
						return state1
				}
		}
	}
	return nums
}


func state1(s *stateargs) stateFn {
	w := s.widget
	select {
		case e := <-s.evt:
			if e.Type == tm.EventKey {
				switch e.Ch {
					case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
						s.nums += string(e.Ch)
						return nums	
					case 'q':
						return nil
					case 'p':
						s.pause = ! s.pause
						if s.pause == false {
							w.Retime()
						}
					case 'r':
						w.Rewind()
					case 'z':
						w.Increasing = ! w.Increasing
					case '+':
						w.Len += time.Second * 1
						w.Border.Label = w.Len.String()
					case '-':
						w.Len -= time.Second * 1
						w.Border.Label =  w.Len.String()
				}
			}

		default:
			termui.Render(w)
			if ! s.pause {
				w.Update()
			}
			time.Sleep(time.Second / 4)
		}
	return state1
}

func main() {
	flag.Parse()
go func() {
	log.Println(http.ListenAndServe(":6060", nil))
}()

	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	termui.UseTheme("helloworld")
	W, H := tm.Size()
	w := timerwidget.New(*fuse, W, H, "timer widget")

	evt := make(chan tm.Event)
	go func() {
		for {
			evt <- tm.PollEvent()
		}
	}()
	pause := false
	s := stateargs{pause:pause, evt: evt, widget: w}
	newstate := state1(&s)
	for newstate != nil {
		
		newstate = newstate(&s)
				
	}

}
