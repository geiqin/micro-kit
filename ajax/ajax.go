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
		ret.Message = e.Detail
	} else {
		ret.Code = 0
		ret.Data = who
	}
	return helper.JsonEncode(ret)
}
