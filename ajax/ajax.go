package ajax

import (
	"github.com/geiqin/gotools/helper"
	"github.com/micro/go-micro/v2/errors"
)

// 主要针对Service层结果转换为Json数据
func ResponseData(who interface{}, err error) string {
	ret := &ResultData{}
	if err != nil {
		e := errors.Parse(err.Error())
		ret.Code = int64(e.Code)
		ret.Msg = e.Detail
	} else {
		ret.Code = 0
		ret.Data = who
	}
	return helper.JsonEncode(ret)
}

// 返回文本信息(默认为成功信息)
func ResponseMsg(msg string, errCode ...int64) string {
	ret := &ResultData{}
	ret.Code = 0
	ret.Msg = msg

	if errCode != nil {
		ret.Code = errCode[0]
	}
	return helper.JsonEncode(ret)
}
