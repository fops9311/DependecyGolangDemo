package adapter

type Adapter struct {
	name string
}

func New(name string) *Adapter {
	port := Adapter{name: "The adapter"}
	return &port
}
func (p Adapter) GetData() (data string) {
	return "some data from " + p.name
}
