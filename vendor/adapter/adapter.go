package adapter

func New() (func() (out <-chan interface{}), func(string)) {
	var outDone chan interface{}

	Pipeline := func() (out <-chan interface{}) {
		outDone = make(chan interface{})
		return outDone
	}
	SendData := func(data string) {
		outDone <- data
	}
	return Pipeline, SendData
}
