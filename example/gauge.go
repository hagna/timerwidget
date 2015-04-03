// Copyright 2015 Zack Guo <gizak@icloud.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.


package main

import ("github.com/gizak/termui"
	tm "github.com/nsf/termbox-go"
	"github.com/hagna/timerwidget"
	"time"
	"flag"
)

var fuse = flag.Duration("t", 10 * time.Second, "timeout")

func main() {
	flag.Parse()
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
FOR:
	for {
		select {
		case e := <-evt:
			if e.Type == tm.EventKey {
				switch e.Ch {
					case 'q':
						break FOR
					case 'p':
						pause = ! pause
						if pause == false {
							w.Retime()
						}
					case 'r':
						w.Rewind()
					case 'z':
						w.Decreasing = ! w.Decreasing
				}
			}

		default:
			termui.Render(w)
			if ! pause {
				w.Update()
			}
			time.Sleep(time.Second / 2)
		}
	}

}
