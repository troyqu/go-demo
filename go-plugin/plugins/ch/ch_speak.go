package main

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
)

type ChSpeak struct{}

func (p ChSpeak) Say(msg string) {
	fmt.Printf("Running Plugin ChSpeak msg=%s \n", msg)
}

//使用变量导出插件包，配合使用变量来作为查找插件的符号
var Speak ChSpeak

//使用方法导出插件包，配合使用方法来作为查找插件的符号
func NewPlugin() ext.Speak {
	return ChSpeak{}
}

