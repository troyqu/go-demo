package main

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
	"github/swallretu/go-demo/go-plugin/pkg/database"
)

type ChSpeak struct{}

func (p ChSpeak) Say(msg string) {
	mockparam := "ch"
	result := database.MockDBOperator(mockparam)
	fmt.Printf("Running Plugin ChSpeak msg=%s result=%s\n", msg, result)
}

//使用变量导出插件包，配合使用变量来作为查找插件的符号
var Speak ChSpeak

//使用方法导出插件包，配合使用方法来作为查找插件的符号
func NewPlugin() ext.Speak {
	return ChSpeak{}
}

