package main

import (
	"github.com/gin-gonic/gin"
	"go-gin/tool"
)

func main() {
	router := gin.Default()
	router.POST("/worktool/callback/ops", tool.WorkTool)
	router.Run("0.0.0.0:18888")
}
