package globalData

//公共常量数据
type ConstService struct {
	// -------------------- 公共 --------------------
	CommonUserRegTypeList              []*ConstListInfo `json:"common_user_reg_type_list"`              // 用户注册类型列表
	CommonLoginTypeList                []*ConstListInfo `json:"common_login_type_list"`                 // 登录方式
	CommonGenderList                   []*ConstListInfo `json:"common_gender_list"`                     // 性别
	CommonCloseOpenList                []*ConstListInfo `json:"common_close_open_list"`                 // 关闭开启状态
	CommonIsEnableTips                 []*ConstListInfo `json:"common_is_enable_tips"`                  // 是否启用
	CommonIsEnableList                 []*ConstListInfo `json:"common_is_enable_list"`                  // 是否启用
	CommonIsShowList                   []*ConstListInfo `json:"common_is_show_list"`                    // 是否显示
	CommonStateList                    []*ConstListInfo `json:"common_state_list"`                      // 状态
	CommonExcelCharsetList             []*ConstListInfo `json:"common_excel_charset_list"`              // excel编码列表
	CommonExcelExportTypeList          []*ConstListInfo `json:"common_excel_export_type_list"`          // excel导出类型列表
	CommonMapTypeList                  []*ConstListInfo `json:"common_map_type_list"`                   // 地图类型列表
	CommonOrderPayStatus               []*ConstListInfo `json:"common_order_pay_status"`                // 支付支付状态
	CommonOrderStatus                  []*ConstListInfo `json:"common_order_status"`                    // 订单状态
	CommonPlatformType                 []*ConstListInfo `json:"common_platform_type"`                   // 所属平台
	CommonAppType                      []*ConstListInfo `json:"common_app_type"`                        // app平台
	CommonAppminiType                  []*ConstListInfo `json:"common_appmini_type"`                    // 小程序平台
	CommonDeductionInventoryRulesList  []*ConstListInfo `json:"common_deduction_inventory_rules_list"`  // 扣除库存规则
	CommonSalesCountIncRulesList       []*ConstListInfo `json:"common_sales_count_inc_rules_list"`      // 商品增加销量规则
	CommonIsReadList                   []*ConstListInfo `json:"common_is_read_list"`                    // 是否已读
	CommonMessageTypeList              []*ConstListInfo `json:"common_message_type_list"`               // 消息类型
	CommonIntegralLogTypeList          []*ConstListInfo `json:"common_integral_log_type_list"`          // 用户积分 - 操作类型
	CommonIsShelvesList                []*ConstListInfo `json:"common_is_shelves_list"`                 // 是否上架/下架
	CommonIsTextList                   []*ConstListInfo `json:"common_is_text_list"`                    // 是否
	CommonAppEventType                 []*ConstListInfo `json:"common_app_event_type"`                  //app事件类型
	CommonOrderAftersaleTypeList       []*ConstListInfo `json:"common_order_aftersale_type_list"`       //订单售后类型
	CommonOrderAftersaleStatusList     []*ConstListInfo `json:"common_order_aftersale_status_list"`     // 订单售后状态
	CommonOrderAftersaleRefundmentList []*ConstListInfo `json:"common_order_aftersale_refundment_list"` // 订单售后退款方式
	CommonSiteTypeList                 []*ConstListInfo `json:"common_site_type_list"`                  // 站点类型
	CommonOrderTypeList                []*ConstListInfo `json:"common_order_type_list"`                 // 订单类型
	CommonAdminStatusList              []*ConstListInfo `json:"common_admin_status_list"`               // 管理员状态
	CommonPayLogStatusList             []*ConstListInfo `json:"common_pay_log_status_list"`             // 支付日志状态
	CommonTimezoneList                 []*ConstListInfo `json:"common_timezone_list"`                   // 时区
	// -------------------- 正则 --------------------
	CommonRegexUsername        string `json:"common_regex_username"`         // 用户名
	CommonRegexPwd             string `json:"common_regex_pwd"`              // 用户名
	CommonRegexAlphaNumber     string `json:"common_regex_alpha_number"`     // 包含字母和数字、6~16个字符
	CommonRegexMobile          string `json:"common_regex_mobile"`           // 手机号码
	CommonRegexTel             string `json:"common_regex_tel"`              // 座机号码
	CommonRegexEmail           string `json:"common_regex_email"`            // 电子邮箱
	CommonRegexIdCard          string `json:"common_regex_id_card"`          // 身份证号码
	CommonRegexPrice           string `json:"common_regex_price"`            // 价格格式
	CommonRegexIp              string `json:"common_regex_ip"`               // ip
	CommonRegexUrl             string `json:"common_regex_url"`              // url
	CommonRegexSort            string `json:"common_regex_sort"`             // 顺序
	CommonRegexDate            string `json:"common_regex_date"`             // 日期
	CommonRegexScore           string `json:"common_regex_score"`            // 分数
	CommonRegexPageNumber      string `json:"common_regex_page_number"`      // 分页
	CommonRegexInterval        string `json:"common_regex_interval"`         // 时段格式 10:00-10:45
	CommonRegexColor           string `json:"common_regex_color"`            // 颜色
	CommonRegexIdCommaSplit    string `json:"common_regex_id_comma_split"`   // id逗号隔开
	CommonRegexUrlHtmlSuffix   string `json:"common_regex_url_html_suffix"`  // url伪静态后缀
	CommonRegexImageProportion string `json:"common_regex_image_proportion"` // 图片比例值
	CommonRegexVersion         string `json:"common_regex_version"`          // 版本号

}

//加载配置信息
func loadConst() *ConstService {
	return &ConstService{
		CommonUserRegTypeList: []*ConstListInfo{
			{Value: "username", Name: "账号"},
			{Value: "sms", Name: "短信"},
			{Value: "email", Name: "邮箱"},
		},
		CommonLoginTypeList: []*ConstListInfo{
			{Value: "username", Name: "帐号密码", Checked: true},
			{Value: "email", Name: "邮箱验证码"},
			{Value: "sms", Name: "手机验证码"},
		},
		CommonGenderList: []*ConstListInfo{
			{Value: "0", Name: "保密", Checked: true},
			{Value: "1", Name: "女"},
			{Value: "2", Name: "男"},
		},
		CommonCloseOpenList: []*ConstListInfo{
			{Value: "0", Name: "关闭"},
			{Value: "1", Name: "开启"},
		},
		CommonIsEnableTips: []*ConstListInfo{
			{Value: "0", Name: "未启用"},
			{Value: "1", Name: "已启用"},
		},
		CommonIsEnableList: []*ConstListInfo{
			{Value: "0", Name: "不启用"},
			{Value: "1", Name: "启用", Checked: true},
		},
		CommonIsShowList: []*ConstListInfo{
			{Value: "0", Name: "不显示"},
			{Value: "1", Name: "显示", Checked: true},
		},
		CommonStateList: []*ConstListInfo{
			{Value: "0", Name: "不可用"},
			{Value: "1", Name: "可用", Checked: true},
		},
		CommonExcelCharsetList: []*ConstListInfo{
			{Value: "utf-8", Name: "utf-8", Checked: true},
			{Value: "gbk", Name: "gbk"},
		},
		CommonExcelExportTypeList: []*ConstListInfo{
			{Value: "0", Name: "CSV", Checked: true},
			{Value: "1", Name: "Excel"},
		},
		CommonMapTypeList: []*ConstListInfo{
			{Value: "baidu", Name: "百度地图", Checked: true},
			{Value: "amap", Name: "高德地图"},
			{Value: "tencent", Name: "腾讯地图"},
			{Value: "tianditu", Name: "天地图"},
		},
		CommonOrderPayStatus: []*ConstListInfo{
			{Value: "0", Name: "待支付", Checked: true},
			{Value: "1", Name: "已支付"},
			{Value: "2", Name: "已退款"},
			{Value: "3", Name: "部分退款"},
		},
		CommonOrderStatus: []*ConstListInfo{
			{Value: "0", Name: "待确认", Checked: true},
			{Value: "1", Name: "待付款"},
			{Value: "2", Name: "待发货"},
			{Value: "3", Name: "待收货"},
			{Value: "4", Name: "已完成"},
			{Value: "5", Name: "已取消"},
			{Value: "6", Name: "已关闭"},
		},
		CommonPlatformType: []*ConstListInfo{
			{Value: "pc", Name: "PC网站"},
			{Value: "h5", Name: "H5手机网站"},
			{Value: "ios", Name: "苹果APP"},
			{Value: "android", Name: "安卓APP"},
			{Value: "weixin", Name: "微信小程序"},
			{Value: "alipay", Name: "支付宝小程序"},
			{Value: "baidu", Name: "百度小程序"},
			{Value: "toutiao", Name: "头条小程序"},
			{Value: "qq", Name: "QQ小程序"},
			{Value: "kuaishou", Name: "快手小程序"},
		},
		CommonAppType: []*ConstListInfo{
			{Value: "ios", Name: "苹果APP"},
			{Value: "android", Name: "安卓APP"},
		},
		CommonAppminiType: []*ConstListInfo{
			{Value: "weixin", Name: "微信小程序"},
			{Value: "alipay", Name: "支付宝小程序"},
			{Value: "baidu", Name: "百度小程序"},
			{Value: "toutiao", Name: "头条小程序"},
			{Value: "qq", Name: "QQ小程序"},
			{Value: "kuaishou", Name: "快手小程序"},
		},
		CommonDeductionInventoryRulesList: []*ConstListInfo{
			{Value: "0", Name: "订单确认成功"},
			{Value: "1", Name: "订单支付成功"},
			{Value: "2", Name: "订单发货"},
		},
		CommonSalesCountIncRulesList: []*ConstListInfo{
			{Value: "0", Name: "订单支付"},
			{Value: "1", Name: "订单收货"},
		},
		CommonIsReadList: []*ConstListInfo{
			{Value: "0", Name: "未读", Checked: true},
			{Value: "1", Name: "已读"},
		},
		CommonMessageTypeList: []*ConstListInfo{
			{Value: "0", Name: "默认", Checked: true},
		},
		CommonIntegralLogTypeList: []*ConstListInfo{
			{Value: "0", Name: "减少", Checked: true},
			{Value: "1", Name: "增加"},
		},
		CommonIsShelvesList: []*ConstListInfo{
			{Value: "0", Name: "下架"},
			{Value: "1", Name: "上架", Checked: true},
		},
		CommonIsTextList: []*ConstListInfo{
			{Value: "0", Name: "否", Checked: true},
			{Value: "1", Name: "是"},
		},
		CommonAppEventType: []*ConstListInfo{
			{Value: "0", Name: "WEB页面"},
			{Value: "1", Name: "内部页面(小程序/APP内部地址)"},
			{Value: "2", Name: "外部小程序(同一个主体下的小程序appid)"},
			{Value: "3", Name: "跳转原生地图查看指定位置"},
			{Value: "4", Name: "拨打电话"},
		},
		CommonOrderAftersaleTypeList: []*ConstListInfo{
			{Value: "0", Name: "仅退款", Desc: "未收到货(未签收),协商同意前提下", Icon: "am-icon-random", Class: "am-fl"},
			{Value: "1", Name: "退款退货", Desc: "已收到货,需要退换已收到的货物", Icon: "am-icon-retweet", Class: "am-fr"},
		},
		CommonOrderAftersaleStatusList: []*ConstListInfo{
			{Value: "0", Name: "待确认"},
			{Value: "1", Name: "待退货"},
			{Value: "2", Name: "待审核"},
			{Value: "3", Name: "退款中"},
			{Value: "4", Name: "已完成"},
			{Value: "5", Name: "已拒绝"},
			{Value: "6", Name: "已取消"},
		},
		CommonOrderAftersaleRefundmentList: []*ConstListInfo{
			{Value: "0", Name: "原路退回"},
			{Value: "1", Name: "退至钱包"},
			{Value: "2", Name: "手动处理"},
		},
		CommonSiteTypeList: []*ConstListInfo{
			{Value: "0", Name: "快递"},
			{Value: "1", Name: "展示"},
			{Value: "2", Name: "自提"},
			{Value: "3", Name: "虚拟售卖"},
			{Value: "4", Name: "快递+自提", Checked: true}, //, "is_ext": 1
		},
		CommonOrderTypeList: []*ConstListInfo{
			{Value: "0", Name: "快递"},
			{Value: "1", Name: "展示"},
			{Value: "2", Name: "自提"},
			{Value: "3", Name: "虚拟销售"},
		},
		CommonAdminStatusList: []*ConstListInfo{
			{Value: "0", Name: "暂停"},
			{Value: "1", Name: "正常", Checked: true},
			{Value: "2", Name: "已离职"},
		},
		CommonPayLogStatusList: []*ConstListInfo{
			{Value: "0", Name: "待支付", Checked: true},
			{Value: "1", Name: "已支付"},
			{Value: "2", Name: "已关闭"},
		},
		CommonTimezoneList: []*ConstListInfo{
			{Value: "Pacific/Pago_Pago", Name: "(标准时-11:00) 中途岛、萨摩亚群岛"},
			{Value: "Pacific/Rarotonga", Name: "(标准时-10:00) 夏威夷"},
			{Value: "Pacific/Gambier", Name: "(标准时-9:00) 阿拉斯加"},
			{Value: "America/Dawson", Name: "(标准时-8:00) 太平洋时间(美国和加拿大)"},
			{Value: "America/Creston", Name: "(标准时-7:00) 山地时间(美国和加拿大)"},
			{Value: "America/Belize", Name: "(标准时-6:00) 中部时间(美国和加拿大)、墨西哥城"},
			{Value: "America/Eirunepe", Name: "(标准时-5:00) 东部时间(美国和加拿大)、波哥大"},
			{Value: "America/Antigua", Name: "(标准时-4:00) 大西洋时间(加拿大)、加拉加斯"},
			{Value: "America/Argentina/Buenos_Aires", Name: "(标准时-3:00) 巴西、布宜诺斯艾利斯、乔治敦"},
			{Value: "America/Noronha", Name: "(标准时-2:00) 中大西洋"},
			{Value: "Atlantic/Cape_Verde", Name: "(标准时-1:00) 亚速尔群岛、佛得角群岛"},
			{Value: "Africa/Ouagadougou", Name: "(格林尼治标准时) 西欧时间、伦敦、卡萨布兰卡"},
			{Value: "Europe/Andorra", Name: "(标准时+1:00) 中欧时间、安哥拉、利比亚"},
			{Value: "Europe/Mariehamn", Name: "(标准时+2:00) 东欧时间、开罗，雅典"},
			{Value: "Asia/Bahrain", Name: "(标准时+3:00) 巴格达、科威特、莫斯科"},
			{Value: "Asia/Dubai", Name: "(标准时+4:00) 阿布扎比、马斯喀特、巴库"},
			{Value: "Asia/Kolkata", Name: "(标准时+5:00) 叶卡捷琳堡、伊斯兰堡、卡拉奇"},
			{Value: "Asia/Dhaka", Name: "(标准时+6:00) 阿拉木图、 达卡、新亚伯利亚"},
			{Value: "Indian/Christmas", Name: "(标准时+7:00) 曼谷、河内、雅加达"},
			{Value: "Asia/Shanghai", Name: "(标准时+8:00)北京、重庆、香港、新加坡"},
			{Value: "Australia/Darwin", Name: "(标准时+9:00) 东京、汉城、大阪、雅库茨克"},
			{Value: "Australia/Adelaide", Name: "(标准时+10:00) 悉尼、关岛"},
			{Value: "Australia/Currie", Name: "(标准时+11:00) 马加丹、索罗门群岛"},
			{Value: "Pacific/Fiji", Name: "(标准时+12:00) 奥克兰、惠灵顿、堪察加半岛"},
		},
		CommonRegexUsername:        "^[A-Za-z0-9_]{2,18}$",
		CommonRegexPwd:             "^.{6,18}$",
		CommonRegexAlphaNumber:     "^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{6,16}$",
		CommonRegexMobile:          "^1((3|4|5|6|7|8|9){1}\\d{1})\\d{8}$",
		CommonRegexTel:             "^\\d{3,4}-?\\d{8}$",
		CommonRegexEmail:           "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$",
		CommonRegexIdCard:          "^(\\d{15}$|^\\d{18}$|^\\d{17}(\\d|X|x))$",
		CommonRegexPrice:           "^([0-9]{1}\\d{0,7})(\\.\\d{1,2})?$",
		CommonRegexIp:              "^(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])\\.(\\d{1,2}|1\\d\\d|2[0-4]\\d|25[0-5])$",
		CommonRegexUrl:             "^http[s]?:\\/\\/[A-Za-z0-9-]+\\.[A-Za-z0-9]+[\\/=\\?%\\-&_~`@[\\]\\':+!]*([^<>\\\"\\\"])*$",
		CommonRegexSort:            "^[0-9]{1,3}$",
		CommonRegexDate:            "^\\d{4}-\\d{2}-\\d{2}$",
		CommonRegexScore:           "^[0-9]{1,3}$",
		CommonRegexPageNumber:      "^[1-9]{1}[0-9]{0,2}$",
		CommonRegexInterval:        "^\\d{2}\\:\\d{2}\\-\\d{2}\\:\\d{2}$",
		CommonRegexColor:           "^(#([a-fA-F0-9]{6}|[a-fA-F0-9]{3}))?$",
		CommonRegexIdCommaSplit:    "^\\d(\\d|,?)*\\d$",
		CommonRegexUrlHtmlSuffix:   "^[a-z]{0,8}$",
		CommonRegexImageProportion: "^([1-9]{1}[0-9]?|[1-9]{1}[0-9]?\\.{1}[0-9]{1,2}|100|0)?$",
		CommonRegexVersion:         "^[0-9]{1,6}\\.[0-9]{1,6}\\.[0-9]{1,6}$",
	}
}

//常量列表对象
type ConstListInfo struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	Icon    string `json:"icon"`
	Class   string `json:"class"`
	Desc    string `json:"desc"`
	Checked bool   `json:"checked"`
}

//获取公共常量数据
func MyConst() *ConstService {
	return loadConst()
}

//获取常量值的名称
func GetMyConstName(dataList []*ConstListInfo, value string) string {
	if val := GetMyConstItem(dataList, value); val != nil {
		return val.Name
	}
	return ""
}

//获取常量值的对象
func GetMyConstItem(dataList []*ConstListInfo, value string) *ConstListInfo {
	for _, v := range dataList {
		if v.Value == value {
			return v
		}
	}
	return nil
}
