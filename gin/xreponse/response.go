package xreponse

import (
	"github.com/geiqin/gotools/ajax"
	"github.com/geiqin/micro-kit/protobuf/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 失败数据处理
func Error(c *gin.Context, err error, msg ...string) {
	var res common.Error

	if err != nil {
		res.Message = err.Error()
	}
	if msg != nil {
		res.Message = msg[0]
	}
	//res.RequestId = tools.GenerateMsgIDFromContext(c)
	c.AbortWithStatusJSON(http.StatusOK, res)
}

// 失败数据处理
func Failed(c *gin.Context, errMsg string, errCode ...int32) {
	var res common.Error
	var code int32 = 400
	if errCode != nil {
		code = errCode[0]
	}
	res.Message = errMsg
	res.Code = code
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func SucceedByEntity(ctx *gin.Context, entity interface{}) {
	type AjaxEntityData struct {
		Entity interface{} `json:"entity,omitempty"`
	}
	Succeed(ctx, ajax.Succeed(&AjaxEntityData{Entity: entity}))
}

func SucceedByItems(ctx *gin.Context, items interface{}, pager ...interface{}) {
	type AjaxItemData struct {
		Items interface{} `json:"items,omitempty"`
	}
	Succeed(ctx, ajax.Succeed(&AjaxItemData{Items: items}))
}

// 通常成功数据处理
func Succeed(c *gin.Context, data interface{}, msg ...string) {
	c.AbortWithStatusJSON(http.StatusOK, data)
}

// 分页数据处理
func PageOK(c *gin.Context, items interface{}, count int, pageIndex int, pageSize int) {
	var res ResultData
	res.Items = items
	res.Pager = &common.Pager{
		Paged:     0,
		Total:     0,
		PageCount: 0,
		PageSize:  0,
		PrevPage:  0,
		LastPage:  0,
	}
	Succeed(c, res)
}

// 兼容函数
func Custom(c *gin.Context, data gin.H) {
	c.AbortWithStatusJSON(http.StatusOK, data)
}
