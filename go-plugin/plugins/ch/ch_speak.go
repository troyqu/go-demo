package ch

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
)

type ChSpeak struct{}

func (p *ChSpeak) Say(msg string) {
	fmt.Printf("Running Plugin ChSpeak msg=%s \n", msg)
}

func NewPlugin() ext.Speak {
	return &ChSpeak{}
}
