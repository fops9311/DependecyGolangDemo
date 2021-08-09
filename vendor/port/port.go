package port

import (
	ad "adapter"
	"strings"
)

type Port struct {
	adapter *ad.Adapter
}

func New() *Port {
	adapter := ad.New("P1")
	port := Port{adapter: adapter}
	return &port
}
func (a Port) GetModifiedData() (data string) {
	return strings.ToUpper(a.adapter.GetData()) + " THROUGH THE PORT"
}
