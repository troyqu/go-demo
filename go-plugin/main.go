package main

import (
	"fmt"
	"github/swallretu/go-demo/go-plugin/ext"
	"net/http"
	"os"
	"plugin"
	"strings"
)

var PluginRegistry = make(map[string]*ext.Speak)

func RegisterPlugin(name string, pluginPath string) {
	p, err := plugin.Open(pluginPath)
	if err != nil {
		fmt.Printf("Error opening plugin %s: %s\n", name, err)
		return
	}

	sym, err := p.Lookup("NewPlugin")
	if err != nil {
		fmt.Printf("Error looking up NewPlugin in %s: %s\n", name, err)
		return
	}

	newPlugin, ok := sym.(func() *ext.Speak)
	if !ok {
		fmt.Printf("Error converting NewPlugin to func() ext.Speak in %s\n", name)
		return
	}

	PluginRegistry[name] = newPlugin()
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
			RegisterPlugin(strings.TrimSuffix(file.Name(), ".so"), fmt.Sprintf("%s/%s", pluginsDir, file.Name()))
		}
	}

	// 启动 HTTP 服务
	http.HandleFunc("/call/plugin", handlePlugin)
	http.ListenAndServe(":8080", nil)
}

func handlePlugin(w http.ResponseWriter, r *http.Request) {
	pluginType := r.URL.Query().Get("type")

	handler := (ext.Speak)(*PluginRegistry[pluginType])

	handler.Say(pluginType)

	w.Write([]byte(fmt.Sprintf("Running plugin: %s", pluginType)))
}
