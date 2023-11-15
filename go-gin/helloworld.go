package main
/**
全部放开注释替换main方法，是呀gin helloworld example
 */



//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"go-gin/tool"
//	"net/http"
//)
//
//func main() {
//
//	fmt.Println("======分割线--------")
//	router := gin.Default()
//
//	router.LoadHTMLGlob("./*.html")
//
//	router.GET("/go/login", func(c *gin.Context) {
//		c.HTML(http.StatusOK, "index.html", nil)
//	})
//
//	router.GET("/go/troy", func(c *gin.Context) {
//		fmt.Println("欢迎点击查看troy，马上跳转到troy-info")
//		c.Redirect(301, "http://127.0.0.1:8888/go/troy-info")
//	})
//
//	type msg struct {
//		Name    string `json:"name"`
//		Message string `json:"message"`
//		Age     int    `json:"age"`
//	}
//
//	router.GET("/go/troy-info", func(c *gin.Context) {
//		data := msg{
//			Name:    "troy",
//			Message: "平平无奇小书童",
//			Age:     18,
//		}
//		c.JSON(200, data)
//	})
//
//	router.Run(":8888")
//}
