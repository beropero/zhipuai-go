package httpclient

var (
	Msg = &Glmreq{
		Messages: []Message{},
		Stream: true,
	} 
)


// glm-4 请求体
type Glmreq struct {
	Model    string `json:"model"`
	Messages []Message `json:"messages"`
	Stream bool `json:"stream"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
} 

// cogview-3 请求体
type Cogreq struct {
	Model string `json:"model"`
	Prompt string `json:"prompt"`
}