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

	w := timerwidget.New(*fuse, 50, 10, "timer widget")

	evt := make(chan tm.Event)
	go func() {
		for {
			evt <- tm.PollEvent()
		}
	}()
	for {
		select {
		case e := <-evt:
			if e.Type == tm.EventKey && e.Ch == 'q' {
				return
			}
		default:
			termui.Render(w)
			w.Update()
			time.Sleep(time.Second / 2)
		}
	}

}
