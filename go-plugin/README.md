# go-plugin
代码展示使用go实现动态加载插件，从而提高程序的可扩展性，并且实现对于接口实现的完全解藕

## 代码结构
这里也包含了最终生成的插件编译文件(*.so)，后续介绍so文件生成方式
```shell script
➜  go-plugin git:(main) ✗ tree
.
├── README.md
├── ext
│   └── ext.go
├── go.mod
├── main.go
└── plugins
    ├── ch
    │   └── ch_speak.go
    ├── ch_speak.so
    ├── en
    │   └── en_speak.go
    └── en_speak.so

4 directories, 8 files
➜  go-plugin git:(main) ✗ 

```

## 目录解释
- ext：定义统一插件接口，这里主要定义一个Say(msg string)作为演示
- plugins：对于插件的实现，不同目录存放各自的实现
- main.go：程序启动入口

## 代码描述
ext包和plugins包不做太多解释来，分别定义了插件定义，以及对应的中文(ch)和英文(en)插件实现

**main.go**

通过定义全局map来存放所有实现的插件，包括后续如果扩展插件实现，只需要在plugins目录下创建对应目录实现代码，并且编译生成新插件so文件，主程序代码不需要进行任何调整即可实现对于插件的支持。

>main.go整体思路
> 1. 通过对于插件存放目录./plugins目录下的所有插件实现so文件进行扫描
> 2. 然后通过go plugin机制将插件加载到全局map中
> 3. 程序只需要根据业务需要，通过key在全局map中获取自己需要使用的插件即可
> 4. 后期业务扩展只需要完全实现新的插件代码，然后重启服务之后，主程序就可以使用最新的插件，而不需要调整主程序代码

main.go中分别实现了使用方法和变量作为导出插件实现的逻辑。

>这里解释一下项目中handler方法是写死了文件名称对应的map key，实际项目中可以通过api传参灵活实现。
```go
func handlePlugin() {
	fmt.Printf("using plugins...\n")
	plugins := []string{"ch_speak", "en_speak"}

	for _, pluginType := range plugins {
		handler := (ext.Speak)(*PluginRegistry[pluginType])
		handler.Say(pluginType)
	}

	fmt.Printf("process plugins finished\n")
}
```


### 生成so文件
执行`go build -o plugins/en_speak.so -buildmode=plugin ./plugins/en/en_speak.go`来生成so插件文件

- -buildmode=plugin： 通过-buildmode=plugin来表明以插件方式构建
- -o plugins/en_speak.so：声明so文件输出位置
- ./plugins/en/en_speak.go 指定使用那个文件来生成插件so文件

实际执行效果
```shell script
➜  go-plugin git:(main) ✗ pwd
/Users/qujianfei/gitProject/go-demo/go-plugin
➜  go-plugin git:(main) ✗ go build -o plugins/en_speak.so -buildmode=plugin ./plugins/en/en_speak.go
➜  go-plugin git:(main) ✗ go build -o plugins/ch_speak.so -buildmode=plugin ./plugins/ch/ch_speak.go
➜  go-plugin git:(main) ✗ 

```

## 运行程序
```shell script
GOROOT=/usr/local/Cellar/go@1.20/1.20.9/libexec #gosetup
GOPATH=/Users/qujianfei/go #gosetup
/usr/local/Cellar/go@1.20/1.20.9/libexec/bin/go build -o /private/var/folders/7x/3t_1q_tj5k1g8kyrbsp7zsh40000gn/T/___go_build_github_swallretu_go_demo_go_plugin github/swallretu/go-demo/go-plugin #gosetup
/private/var/folders/7x/3t_1q_tj5k1g8kyrbsp7zsh40000gn/T/___go_build_github_swallretu_go_demo_go_plugin
ch_speak
./plugins/ch_speak.so 
Registered plugin: ch_speak
en_speak
./plugins/en_speak.so 
Registered plugin: en_speak
using plugins...
Running Plugin ChSpeak msg=ch_speak 
Running Plugin EnSpeak msg=en_speak 
process plugins finished

Process finished with exit code 0

```

https://pkg.go.dev/plugin@master