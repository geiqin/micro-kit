package ajax

// 结果数据
type ResultData struct {
	Code int64       `json:"code"`          //错误代码: 成功：0 ，其它数字为失败
	Data interface{} `json:"data"`          //成功数据
	Msg  string      `json:"msg,omitempty"` //消息

}

// 分页数据
type PageData struct {
	Pager interface{} `json:"pager,omitempty"`
	Items interface{} `json:"items"`
}

// 输出错误json数据
func Failed(message string, errCode ...int64) *ResultData {
	var code int64 = 400
	if errCode != nil {
		c := errCode[0]
		if c > 0 {
			code = c
		}
	}
	ret := &ResultData{
		Code: code,
	}
	return ret
}

// 输出成功json数据
func Succeed(data interface{}) *ResultData {
	ret := &ResultData{
		Code: 0,
		Data: data,
	}
	return ret
}
