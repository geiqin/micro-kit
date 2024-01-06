package xreponse

//结果数据
type AjaxResultData struct {
	Code    int64       `json:"code"`              //错误代码: 成功：1 ，其它数字为失败
	Data    interface{} `json:"data"`              //成功数据
	Msg     string      `json:"msg,omitempty"`     //错误消息
	Message string      `json:"message,omitempty"` //错误消息[兼容旧版本]

}

//输出成功json数据
func AjaxSucceed(data interface{}) *AjaxResultData {
	ret := &AjaxResultData{
		Code: 1,
		Data: data,
	}
	return ret
}
