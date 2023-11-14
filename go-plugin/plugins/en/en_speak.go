package en

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
)

type EnSpeak struct{}

func (p *EnSpeak) Say(msg string) {
	fmt.Printf("Running Plugin EnSpeak msg=%s \n", msg)
}

func NewPlugin() ext.Speak {
	return &EnSpeak{}
}
