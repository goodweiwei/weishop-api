package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"weishop-api/user-web/forms"
)

func SendSms(ctx *gin.Context)  {
	SendSmsForm := forms.SendSmsForm{}
	if err := ctx.ShouldBind(&SendSmsForm);err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg":HandleValidatorError(err),
		})
		return
	}

	// 使用阿里云短信发送
	//client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSmsInfo.ApiKey, global.ServerConfig.AliSmsInfo.ApiSecrect)
	//if err != nil {
	//	panic(err)
	//}
	//smsCode := GenerateSmsCode(6)
	//request := requests.NewCommonRequest()
	//request.Method = "POST"
	//request.Scheme = "https" // https | http
	//request.Domain = "dysmsapi.aliyuncs.com"
	//request.Version = "2017-05-25"
	//request.ApiName = "SendSms"
	//request.QueryParams["RegionId"] = "cn-beijing"
	//request.QueryParams["PhoneNumbers"] = sendSmsForm.Mobile                           //手机号
	//request.QueryParams["SignName"] = "慕学在线"                                       //阿里云验证过的项目名 自己设置
	//request.QueryParams["TemplateCode"] = "SMS_181850725"                          //阿里云的短信模板号 自己设置
	//request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	//response, err := client.ProcessCommonRequest(request)
	//fmt.Print(client.DoAction(request, response))
	//if err != nil {
	//	fmt.Print(err.Error())
	//}
}
