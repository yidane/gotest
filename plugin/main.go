package main

import (
	"fmt"
	"os"
	"plugin"
	"runtime"
	"time"
)

func main() {
	p, err := plugin.Open("plugin1.so")
	if err != nil {
		panic(err)
	}
	sb, err := p.Lookup("Say")
	if err != nil {
		panic(err)
	}
	sbs := sb.(func(string) string)

	if len(os.Args) > 2 {
		str, err := p.Lookup("Addr")
		if err != nil {
			panic(err)
		}
		*str.(*string) = os.Args[2]
	}

	fmt.Println(sbs(os.Args[1]))

	fmt.Println("PID = ", os.Getpid())

	go func() {
		var tick = time.NewTicker(time.Second)
		for {
			select {
			case <-tick.C:
				var memStats runtime.MemStats
				runtime.ReadMemStats(&memStats)
				fmt.Println(memStats.Sys)
			}
		}
	}()

	time.Sleep(time.Minute)
}
