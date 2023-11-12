package globalData

type LangService struct {
	FormTimeEndTitle string `json:"form_time_end_title"` //结束时间
}

//获取公共常量数据
func MyLang() *LangService {
	return &LangService{
		FormTimeEndTitle: "结束时间",
	}
}
