package port

type Port interface {
	Pipeline(<-chan interface{}) <-chan interface{}
}

type ImplPort struct {
	Transform func(string) string
}

func (h *ImplPort) Pipeline(in <-chan interface{}, inDone <-chan interface{}) (out <-chan interface{}, done <-chan interface{}) {
	outC := make(chan interface{})
	outDone := make(chan interface{})
	go func() {
		for {
			select {
			case rec := <-in:
				outC <- h.Transform((rec).(string))
			case <-inDone:
				close(outDone)
				return
			}
		}
	}()
	return outC, outDone
}
