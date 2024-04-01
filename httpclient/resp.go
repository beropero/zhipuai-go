package httpclient


// glm-4 响应体

type Glmresp struct {
    ID      string `json:"id"`
    Created int64  `json:"created"`
    Model   string `json:"model"`
    Choices []struct {
        Index int `json:"index"`
        Delta struct {
            Role    string `json:"role"`
            Content string `json:"content"`
        } `json:"delta"`
    } `json:"choices"`
}

// cogview-3 响应体

type Cogresp struct {
    Created string `json:"created"`
    Data []struct{
        Url string `json:"url"`
    }`json:"data"`
}