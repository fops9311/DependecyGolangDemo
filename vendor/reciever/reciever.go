package reciever

type Reciever struct {
	outC chan string
}

func (p *Reciever) Pipeline(in <-chan interface{}, inDone <-chan interface{}) (done <-chan interface{}) {
	outDone := make(chan interface{})
	go func() {
		for {
			select {
			case rec := <-in:
				println((rec).(string))
			case <-inDone:
				println("eop")
				close(outDone)
				return
			}
		}
	}()
	return outDone
}
