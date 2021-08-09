package businesslogic

import (
	"fmt"
	port "port"
)

type App struct {
	port *port.Port
}

func NewApp() *App {
	port := port.New()
	app := App{port: port}
	return &app
}

func (a App) Run() {
	fmt.Println(a.port.GetModifiedData())
}
