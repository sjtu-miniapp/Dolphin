package service
//
//import (
//	"bytes"
//	"encoding/json"
//	"fmt"
//	"io"
//	"net/http"
//	"net/url"
//	"time"
//)
//type PayData struct {
//
//	data map[string]interface{}
//}
//
//func NewPayData() *PayData  {
//
//	return &PayData{
//		data:make(map[string]interface{}),
//	}
//}
//func (pd *PayData) FromJson(r io.Reader) error  {
//	buf := BufPool.Get().(*bytes.Buffer)
//	buf.Reset()
//	defer BufPool.Put(buf)
//	_, err := buf.ReadFrom(r)
//	if err != nil {
//		return err
//	}
//	var jsonTemplate interface{}
//	err = json.Unmarshal(buf.Bytes(), &jsonTemplate)
//	if err != nil {
//		return err
//	}
//	pd.data = jsonTemplate.(map[string]interface{})
//	return nil
//}
//type PayClient struct {
//
//	httpClient *http.Client
//}
//
//func NewPayClient(httpClient *http.Client) *PayClient  {
//
//	if httpClient == nil {
//
//		httpClient = http.DefaultClient
//		httpClient.Timeout = time.Second * 5
//	}
//
//	return &PayClient{
//		httpClient:httpClient,
//	}
//}
//// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/login/auth.code2Session.html
//func (pc *PayClient) Login(jscode string) (*PayData, error)  {
//
//	//js_code fetched from frontend
//	loginUrl := "https://api.weixin.qq.com/sns/jscode2session?appid=" + url.QueryEscape(payConfig.AppId()) +
//		"&secret=" + url.QueryEscape(payConfig.AppSecret()) +
//		"&js_code=" + url.QueryEscape(jscode) +
//		"&grant_type=authorization_code"
//
//
//	httpResp, err := pc.httpClient.Get(loginUrl)
//
//	if err != nil {
//		return  nil, err
//	}
//	defer httpResp.Body.Close()
//
//	if httpResp.StatusCode != http.StatusOK {
//		return nil, fmt.Errorf("http.Status: %s", httpResp.Status)
//	}
//	respData := NewPayData()
//	err = respData.FromJson(httpResp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	return respData, nil
//}