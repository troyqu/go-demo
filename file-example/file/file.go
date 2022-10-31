package file

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// 读取配置文件中的目录，替换gitlab domain
type File struct {
	TargetPath string `yaml:"targetPath"`
	Gitlab Gitlab `yaml:"gitlab"`
}

type Gitlab struct {
	ObsoleteDomain string `yaml:"obsoleteDomain"`
	NewDomain string `yaml:"newDomain"`
}

/**
初始化读取配置
 */
func Init() File{
	var file File
	fileYaml, err := ioutil.ReadFile("/Users/goproject/file-example/config/file/file.yaml")
	fmt.Println("fileYaml:", string(fileYaml))
	if err != nil{
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(fileYaml, &file)
	if err != nil {
		fmt.Println(err.Error())
	}
	return file;
}

func ReplaceGitConfig(gitlabYaml File) {
	fmt.Println("targetPath=",gitlabYaml.TargetPath)
	fmt.Println("Gitlab.ObsoleteDomain=",gitlabYaml.Gitlab.ObsoleteDomain)
	fmt.Println("Gitlab.NewDomain=",gitlabYaml.Gitlab.NewDomain)

	files := getAllPath(gitlabYaml.TargetPath)

	for _, file := range files {
		//fmt.Println(file)
		replaceGitDomain(file, gitlabYaml.Gitlab)
	}
}

/**
遍历目标路径下，所有以.git/config结尾的path
 */
func getAllPath(path string) ([] string){
	var files []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".git/config"){
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files;
}

/**
根据yaml的配置，替换目标路径下.git/config中的gitlab domain
 */
func replaceGitDomain(path string, gitlab Gitlab){
	fmt.Println("replace git domain path=", path)
	file, err := os.OpenFile(path, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("open file filed.", err)
		return
	}

	defer file.Close() //延时关闭文件

	reader := bufio.NewReader(file)
	pos := int64(0) //记录匹配的行位置
	for {
		line, err := reader.ReadString('\n') //读取每一行内容
		if err != nil {
			if err == io.EOF { //读到末尾
				fmt.Println("end of file!")
				break
			} else {
				fmt.Println("read file error!", err)
				return
			}
		}

		if strings.Contains(line, gitlab.ObsoleteDomain){ //包含gitlab过时domain的行
			fmt.Println("before replace line = ",line)
			newLine := strings.Replace(line, gitlab.ObsoleteDomain, gitlab.NewDomain, 1) //替换内容
			fmt.Println("after replace line = ", newLine)
			bytes := []byte(newLine)
			file.WriteAt(bytes, pos) //将新内容写到指定位置
		}
		pos += int64(len(line)) //每读取一行，更新读取到文件内不的位置信息
	}
}

