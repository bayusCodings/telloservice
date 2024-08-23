package seeders

import (
	"fmt"

	"github.com/bayuscodings/telloservice"
)

type Seeders struct {
	app *telloservice.ApplicationHandler
}

func (seeders *Seeders) Logger(message string) {
	fmt.Println(`Seeders :: ` + message)
}

func Init(app *telloservice.ApplicationHandler) {
	s := &Seeders{}
	s.app = app
}
