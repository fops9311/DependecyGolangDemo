package main

import (
	ad "adapter"
	port "port"
	"strings"
)

func main() {
	done := make(chan interface{})

	port1 := port.ImplPort{}

	port2 := port.ImplPort{}

	port3 := port.ImplPort{}

	adapterPipeline, sendData := ad.New()

	port1.Transform = func(in string) string {
		return strings.ToLower(in + " port 1")
	}
	port2.Transform = func(in string) string {
		return strings.ToUpper(in + " port 2")
	}
	port3.Transform = func(in string) string {
		return strings.ToUpper(in + " end of pipeline")
	}

	dataC, endofpipeline := port3.Pipeline(port2.Pipeline(port1.Pipeline(adapterPipeline(), done)))
	go func() {
		for {
			select {
			case str := <-dataC:
				println(str.(string))
			case <-endofpipeline:
				return
			}
		}
	}()
	go func() {
		sendData("tEst1")
		sendData("teSt2")
		sendData("tesT3")
		sendData("Test4")
		close(done)
	}()
	<-endofpipeline
}
