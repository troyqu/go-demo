package main

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
	"os"
	"plugin"
	"strings"
)

var PluginRegistry = make(map[string] *ext.Speak)

func RegisterPluginFunc(name string, pluginPath string) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Printf("Error opening plugin %s: %s\n", name, err)
		return
	}


	//使用方法来作为查找插件的符号
	sym, err := p.Lookup("NewPlugin")
	if err != nil {
		fmt.Printf("Error looking up NewPlugin in %s: %s\n", name, err)
		return
	}

	newPlugin, ok := sym.(func() ext.Speak)
	if !ok {
		fmt.Printf("Error converting NewPlugin to func() ext.Speak in %s\n", name)
		return
	}

	pluginInstance := newPlugin()
	PluginRegistry[name] = &pluginInstance

	fmt.Printf("Registered plugin: %s\n", name)
}

func RegisterPluginVariable(name string, pluginPath string) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Printf("Error opening plugin %s: %s\n", name, err)
		return
	}

	//使用变量来作为查找插件的符号
	sym, err := p.Lookup("Speak")
	if err != nil {
		fmt.Printf("Error looking up NewPlugin in %s: %s\n", name, err)
		return
	}

	var speak ext.Speak
	speak, ok := sym.(ext.Speak)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// 4. use the module
	//speak.Say("hello")

	PluginRegistry[name] = &speak

	fmt.Printf("Registered plugin: %s\n", name)
}

func main() {
	// 扫描所有插件
	pluginsDir := "./plugins"
	files, err := os.ReadDir(pluginsDir)
	if err != nil {
		fmt.Println("Error reading plugins directory:", err)
		return
	}

	// 动态加载插件
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".so") {
			fmt.Println(strings.TrimSuffix(file.Name(), ".so"))
			fmt.Printf(fmt.Sprintf("%s/%s \n", pluginsDir, file.Name()))
			//RegisterPluginVariable(strings.TrimSuffix(file.Name(), ".so"), fmt.Sprintf("%s/%s", pluginsDir, file.Name()))
			RegisterPluginFunc(strings.TrimSuffix(file.Name(), ".so"), fmt.Sprintf("%s/%s", pluginsDir, file.Name()))
		}
	}

	handlePlugin()
}

func handlePlugin() {
	fmt.Printf("using plugins...\n")
	plugins := []string{"ch_speak", "en_speak"}

	for _, pluginType := range plugins {
		handler := (ext.Speak)(*PluginRegistry[pluginType])
		handler.Say(pluginType)
	}

	fmt.Printf("process plugins finished\n")
}
