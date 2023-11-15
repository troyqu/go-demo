package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type WtReq struct {
	Spoken			string `json:"spoken"` //不带@me文本消息
	RawSpoken 		string `json:"rawSpoken"` //原始问题文本
	ReceivedName 	string `json:"receivedName"` //发送消息人
	GroupName 		string `json:"groupName"`
	GroupRemark 	string `json:"groupRemark"`
	RoomType 		string `json:"roomType"`
	AtMe 			string `json:"atMe"`
	TextType 		string `json:"textType"`
}

type WtResp struct {
	Code 			int32 `json:"code"`
	Message 		string `json:"message"`
	Data            WtRespData `json:"data"`
}

type WtRespData struct {
	Type			int32 			`json:"type"`
	Info  			WtRespDataInfo  `json:"info"`
}

type WtRespDataInfo struct {
	Text 			string `json:"text"`
}

func WorkTool(c *gin.Context)  {
	fmt.Printf("进入回调信息, time= %v \n", time.Now().String())
	var request WtReq
	if err := c.ShouldBind(&request); err != nil{
		fmt.Printf("请求格式不合法，请求参数= %+v \n", request)
		fmt.Printf("错误信息：\n", err)
		return
	}
	fmt.Printf("请求内容：request= %+v \n", request)

	var returnMsg string

	fmt.Println("request.Spoken=", 		request.Spoken)
	fmt.Println("request.RawSpoken=", 		request.RawSpoken)
	fmt.Println("request.ReceivedName=", 	request.ReceivedName)
	fmt.Println("request.GroupName=", 		request.GroupName)
	fmt.Println("request.GroupRemark=", 	request.GroupRemark)
	fmt.Println("request.AtMe=", 			request.AtMe)

	instruction := request.Spoken

	if strings.Contains(instruction, "###工单操作###"){
		returnMsg = "我是自定义工单回复消息,帮助创建工单"
	}else if strings.Contains(instruction, "###告警操作###"){
		returnMsg = "我是自定义告警回复消息，帮助处理告警"
	}else if strings.Contains(instruction, "###资源操作###") {
		returnMsg = "我是自定义资源回复消息，帮助操作资源"
	}else{
		returnMsg = "我是自定义通用回复消息，啥也不干"
	}

	wtRespDataInfo := WtRespDataInfo{
		Text: returnMsg,
	}

	wtRespData := WtRespData{
		Type: 0,
		Info: wtRespDataInfo,
	}

	wtResp := WtResp{
		Code 	: 0,
		Message : "SUCCESS",
		Data 	: wtRespData,
	}

	c.JSON(200, wtResp)
}

func handlerNBSP(msg string) string{
	newMsg := strings.Replace(msg," ", "", 1)
	return newMsg;
}