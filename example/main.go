package main

import (
	zhipuai "zhipuai-go"
	"zhipuai-go/consts"
	"zhipuai-go/httpclient"

	"fmt"
)

func main() {
	consts.ApiKey = "your_api_key"

	var msg string

	fmt.Scanln(&msg)

	Msg := httpclient.Message{
		Role:    "user",
		Content: msg,
	}

	/*
		//CogView 模型
		zhipuai.Glmctrl(consts.CogView,Msg)
		for line := range zhipuai.CogRespch {
			fmt.Println(line.Data[0].Url)
		}
	*/
	// Glm-4
	zhipuai.Glmctrl(consts.GLM4, Msg)

	for line := range zhipuai.GlmRespch {
		fmt.Printf(line.Choices[0].Delta.Content)
	}
	fmt.Println()

}
