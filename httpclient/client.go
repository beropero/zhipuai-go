package httpclient

import (
	"zhipuai-go/consts"
	"zhipuai-go/utils"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// 创建http请求
func CreatedHttpRequest(jsonData []byte,url string) (req *http.Request, err error) {
	
	// 获取token
	token, err := utils.GetTocken()
	if err != nil {
		return  
	}

	req, err = http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	// 请求头添加鉴权信息
	req.Header.Add("Authorization", "Bearer " + token)

	req.Header.Set("Content-Type", "application/json")
	
	return
}

// 流式获取响应数据
func GetResponseStream(msg Message, ch chan Glmresp) (err error){

	Msg.Messages = append(Msg.Messages, msg)
	jsonData, err := json.Marshal(Msg)

	if err != nil {
		return err
	}

	
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 获取请求URL
	url := consts.Glm_Api_Url

	// 创建一个请求
	req, _ := CreatedHttpRequest(jsonData, url)


	
	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()


	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	// 以流式方式读取响应体
    scanner := bufio.NewScanner(resp.Body)
    for scanner.Scan() {
        // 将回应发送到通道
		var respData Glmresp 
		line := scanner.Text()
		if line == "" || strings.Contains(line,"[DONE]"){
			continue
		}
		line = strings.TrimPrefix(line, "data:")
		err := json.Unmarshal([]byte(line), &respData)
		if err != nil {
            return  err
        }
		
        ch <- respData
    }
	close(ch)

	return nil
}

func GetResponseImage(msg Message, ch chan Cogresp) (err error) {
	defer close(ch)

	data := &Cogreq{
		consts.CogView,
		msg.Content,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}


	client := &http.Client{}
	url := consts.Cog_Api_Url
	req, err :=  CreatedHttpRequest([]byte(jsonData),url)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respData := Cogresp{}
	body, err  := io.ReadAll(resp.Body)
	json.Unmarshal(body, &respData)
	ch <- respData
	return
}
