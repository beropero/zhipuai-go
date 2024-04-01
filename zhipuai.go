package zhipuai

import (
	"zhipuai-go/consts"
	"zhipuai-go/httpclient"
)

var (
	GlmRespch chan httpclient.Glmresp
	CogRespch chan httpclient.Cogresp
)


func Glmctrl(modelName string, messages httpclient.Message) {
	GlmRespch = make(chan httpclient.Glmresp, 10)
	httpclient.Msg.Model = modelName
	if modelName == consts.GLM4 {
		go httpclient.GetResponseStream(messages, GlmRespch)
	} 

	if modelName == consts.CogView {
		CogRespch = make(chan httpclient.Cogresp, 10)
		go httpclient.GetResponseImage(messages, CogRespch)
	}
}