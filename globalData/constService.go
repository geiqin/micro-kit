package globalData

//公共常量数据
type ConstService struct {
	// -------------------- 公共 --------------------
	CommonUnitType                     []*ConstListInfo `json:"common_uint_type" name:"计量单位类型"`                         // 计量单位类型
	CommonStoreStatusList              []*ConstListInfo `json:"common_store_status_list" name:"店铺状态"`                   // 店铺状态列表
	CommonArticleStatusList            []*ConstListInfo `json:"common_article_status_list" name:"文章状态"`                 // 文章状态列表
	CommonNavigationRouteType          []*ConstListInfo `json:"common_navigation_route_type" name:"网站导航路由类型"`           //网站导航路由类型列表
	CommonUserRegType                  []*ConstListInfo `json:"common_user_reg_type" name:"用户注册类型"`                     // 用户注册类型列表
	CommonLoginType                    []*ConstListInfo `json:"common_login_type" name:"登录方式"`                          // 登录方式
	CommonGenderList                   []*ConstListInfo `json:"common_gender_list" name:"性别"`                           // 性别
	CommonCloseOpenList                []*ConstListInfo `json:"common_close_open_list" name:"关闭开启状态"`                   // 关闭开启状态
	CommonIsEnableTips                 []*ConstListInfo `json:"common_is_enable_tips" name:"是否启用提示"`                    // 是否启用提示
	CommonIsEnableList                 []*ConstListInfo `json:"common_is_enable_list" name:"是否启用"`                      // 是否启用
	CommonIsShowList                   []*ConstListInfo `json:"common_is_show_list" name:"是否显示"`                        // 是否显示
	CommonStateList                    []*ConstListInfo `json:"common_state_list" name:"状态"`                            // 状态
	CommonExcelCharsetList             []*ConstListInfo `json:"common_excel_charset_list" name:"Excel编码列表"`             // excel编码列表
	CommonExcelExportTypeList          []*ConstListInfo `json:"common_excel_export_type_list" name:"Excel导出类型列表"`       // excel导出类型列表
	CommonMapTypeList                  []*ConstListInfo `json:"common_map_type_list" name:"地图类型列表"`                     // 地图类型列表
	CommonOrderPayStatus               []*ConstListInfo `json:"common_order_pay_status" name:"订单支付状态"`                  // 订单支付状态
	CommonOrderStatus                  []*ConstListInfo `json:"common_order_status" name:"订单状态"`                        // 订单状态
	CommonClientType                   []*ConstListInfo `json:"common_client_type" name:"终端类型"`                         // 终端类型
	CommonAppType                      []*ConstListInfo `json:"common_app_type" name:"App平台"`                           // app平台
	CommonAppminiType                  []*ConstListInfo `json:"common_appmini_type" name:"小程序平台"`                       // 小程序平台
	CommonPaymentType                  []*ConstListInfo `json:"common_payment_type" name:"支付类型"`                        // 支付类型
	CommonDeductionInventoryRulesList  []*ConstListInfo `json:"common_deduction_inventory_rules_list" name:"扣除库存规则"`    // 扣除库存规则
	CommonSalesCountIncRulesList       []*ConstListInfo `json:"common_sales_count_inc_rules_list" name:"商品增加销量规则"`      // 商品增加销量规则
	CommonIsReadList                   []*ConstListInfo `json:"common_is_read_list" name:"是否已读"`                        // 是否已读
	CommonMessageTypeList              []*ConstListInfo `json:"common_message_type_list" name:"消息类型"`                   // 消息类型
	CommonIntegralLogTypeList          []*ConstListInfo `json:"common_integral_log_type_list" name:"用户积分操作类型"`          // 用户积分 - 操作类型
	CommonIsShelvesList                []*ConstListInfo `json:"common_is_shelves_list" name:"是否上架"`                     // 是否上架/下架
	CommonIsTextList                   []*ConstListInfo `json:"common_is_text_list" name:"是否"`                          // 是否
	CommonIncDecList                   []*ConstListInfo `json:"common_inc_dec_list" name:"增减"`                          // 增减
	CommonAppEventType                 []*ConstListInfo `json:"common_app_event_type" name:"App事件类型"`                   //app事件类型
	CommonOrderAftersaleTypeList       []*ConstListInfo `json:"common_order_aftersale_type_list" name:"订单售后类型"`         //订单售后类型
	CommonOrderAftersaleStatusList     []*ConstListInfo `json:"common_order_aftersale_status_list" name:"订单售后状态"`       // 订单售后状态
	CommonOrderAftersaleRefundmentList []*ConstListInfo `json:"common_order_aftersale_refundment_list" name:"订单售后退款方式"` // 订单售后退款方式
	CommonOrderTypeList                []*ConstListInfo `json:"common_order_type_list" name:"订单类型"`                     // 订单类型
	CommonLogisticsTypeList            []*ConstListInfo `json:"common_logistics_type_list" name:"物流类型"`                 // 物流类型
	CommonDeliveryModeList             []*ConstListInfo `json:"common_delivery_mode_list" name:"发货模式"`                  // 发货模式
	CommonSiteTypeList                 []*ConstListInfo `json:"common_site_type_list" name:"站点类型"`                      // 站点类型
	CommonAdminStatusList              []*ConstListInfo `json:"common_admin_status_list" name:"管理员状态"`                  // 管理员状态
	CommonMemberStatusList             []*ConstListInfo `json:"common_member_status_list" name:"会员状态"`                  // 会员状态
	CommonMemberTypeList               []*ConstListInfo `json:"common_member_type_list" name:"会员类型"`                    // 会员类型
	CommonLevelPayRuleList             []*ConstListInfo `json:"common_level_pay_rule_list" name:"会员等级支付规则"`             // 会员等级支付规则列表
	CommonPayLogStatusList             []*ConstListInfo `json:"common_pay_log_status_list" name:"支付日志状态"`               // 支付日志状态
	CommonCreateSourceList             []*ConstListInfo `json:"common_create_source_list" name:"创建来源"`                  // 创建来源
	CommonTimezoneList                 []*ConstListInfo `json:"common_timezone_list" name:"时区"`                         // 时区
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

//加载公共配置信息
func LoadCommonConst() *ConstService {
	return &ConstService{
		CommonUnitType: []*ConstListInfo{
			{Value: "1", Text: "计数", Desc: "个、件、次、台、包、箱"},
			{Value: "2", Text: "计重", Desc: "克、千克、吨、公斤"},
			{Value: "3", Text: "计时", Desc: "分钟、小时、天、周、月、季度、年"},
			//{Value: "4", Text: "长度", Desc: "毫米、厘米、分米、米、千米、公里"},
			//{Value: "5", Text: "面积", Desc: "平方毫米、平方厘米、平方分米、平方米、平方千米"},
			//{Value: "6", Text: "体积", Desc: "立方毫米、立方厘米、立方分米、立方米、毫升、升"},
		},
		CommonStoreStatusList: []*ConstListInfo{
			{Value: "0", Text: "已禁用"},
			{Value: "1", Text: "运行中"},
			{Value: "2", Text: "已打烊"},
			{Value: "3", Text: "维护中"},
			{Value: "4", Text: "已过期"},
		},
		CommonArticleStatusList: []*ConstListInfo{
			{Value: "0", Text: "未发布"},
			{Value: "1", Text: "已发布"},
			{Value: "2", Text: "草稿中"},
		},
		CommonNavigationRouteType: []*ConstListInfo{
			{Value: "route", Text: "路由"},
			{Value: "route_parent", Text: "嵌套路由"},
			{Value: "route_child", Text: "子路由"},
			{Value: "link", Text: "外部链接"},
		},
		CommonUserRegType: []*ConstListInfo{
			{Value: "username", Text: "账号"},
			{Value: "sms", Text: "短信"},
			{Value: "email", Text: "邮箱"},
		},
		CommonLoginType: []*ConstListInfo{
			{Value: "username", Text: "帐号密码", Checked: true},
			{Value: "email", Text: "邮箱验证码"},
			{Value: "sms", Text: "手机验证码"},
		},
		CommonGenderList: []*ConstListInfo{
			{Value: "0", Text: "保密", Checked: true},
			{Value: "1", Text: "男"},
			{Value: "2", Text: "女"},
		},
		CommonCloseOpenList: []*ConstListInfo{
			{Value: "0", Text: "关闭"},
			{Value: "1", Text: "开启"},
		},
		CommonIsEnableTips: []*ConstListInfo{
			{Value: "0", Text: "未启用"},
			{Value: "1", Text: "已启用"},
		},
		CommonIsEnableList: []*ConstListInfo{
			{Value: "0", Text: "不启用"},
			{Value: "1", Text: "启用", Checked: true},
		},
		CommonIsShowList: []*ConstListInfo{
			{Value: "0", Text: "不显示"},
			{Value: "1", Text: "显示", Checked: true},
		},
		CommonStateList: []*ConstListInfo{
			{Value: "0", Text: "禁用"},
			{Value: "1", Text: "启用", Checked: true},
		},
		CommonExcelCharsetList: []*ConstListInfo{
			{Value: "utf-8", Text: "utf-8", Checked: true},
			{Value: "gbk", Text: "gbk"},
		},
		CommonExcelExportTypeList: []*ConstListInfo{
			{Value: "0", Text: "CSV", Checked: true},
			{Value: "1", Text: "Excel"},
		},
		CommonMapTypeList: []*ConstListInfo{
			{Value: "baidu", Text: "百度地图", Checked: true},
			{Value: "amap", Text: "高德地图"},
			{Value: "tencent", Text: "腾讯地图"},
			{Value: "tianditu", Text: "天地图"},
		},
		CommonOrderPayStatus: []*ConstListInfo{
			{Value: "0", Text: "待支付", Checked: true},
			{Value: "1", Text: "已支付"},
			{Value: "2", Text: "已退款"},
			{Value: "3", Text: "部分退款"},
		},
		CommonOrderStatus: []*ConstListInfo{
			{Value: "0", Text: "待确认", Checked: true},
			{Value: "1", Text: "待付款"},
			{Value: "2", Text: "待发货"},
			{Value: "3", Text: "待收货"},
			{Value: "4", Text: "已完成"},
			{Value: "5", Text: "已取消"},
			{Value: "6", Text: "已关闭"},
		},
		CommonClientType: []*ConstListInfo{
			{Value: "pc", Text: "PC网站"},
			{Value: "h5", Text: "H5手机网站"},
			{Value: "ios", Text: "苹果APP"},
			{Value: "android", Text: "安卓APP"},
			{Value: "weixin", Text: "微信小程序"},
			{Value: "alipay", Text: "支付宝小程序"},
			{Value: "baidu", Text: "百度小程序"},
			{Value: "toutiao", Text: "头条小程序"},
			{Value: "qq", Text: "QQ小程序"},
			{Value: "kuaishou", Text: "快手小程序"},
			{Value: "cashier", Text: "收银台"},
		},
		CommonAppType: []*ConstListInfo{
			{Value: "ios", Text: "苹果APP"},
			{Value: "android", Text: "安卓APP"},
			{Value: "harmony", Text: "鸿蒙APP"},
		},
		CommonAppminiType: []*ConstListInfo{
			{Value: "weixin", Text: "微信小程序"},
			{Value: "alipay", Text: "支付宝小程序"},
			{Value: "baidu", Text: "百度小程序"},
			{Value: "toutiao", Text: "头条小程序"},
			{Value: "qq", Text: "QQ小程序"},
			{Value: "kuaishou", Text: "快手小程序"},
		},
		CommonPaymentType: []*ConstListInfo{
			{Value: "weixin", Text: "微信支付", Flag: "platform", Icon: ""},
			{Value: "alipay", Text: "支付宝支付", Flag: "platform", Icon: ""},
			{Value: "digitalcny", Text: "数字人民币", Flag: "platform", Icon: ""},
			{Value: "wallet", Text: "钱包支付", Icon: ""},
			{Value: "cod", Text: "货到付款", Icon: ""},
			{Value: "offline", Text: "线下支付", Icon: "", Desc: "需要后台确认已付款"},
			{Value: "hanging", Text: "挂账支付", Icon: ""},
		},
		CommonDeductionInventoryRulesList: []*ConstListInfo{
			{Value: "0", Text: "订单确认成功"},
			{Value: "1", Text: "订单支付成功"},
			{Value: "2", Text: "订单发货"},
		},
		CommonSalesCountIncRulesList: []*ConstListInfo{
			{Value: "0", Text: "订单支付"},
			{Value: "1", Text: "订单收货"},
		},
		CommonIsReadList: []*ConstListInfo{
			{Value: "0", Text: "未读", Checked: true},
			{Value: "1", Text: "已读"},
		},
		CommonMessageTypeList: []*ConstListInfo{
			{Value: "0", Text: "默认", Checked: true},
		},
		CommonIntegralLogTypeList: []*ConstListInfo{
			{Value: "0", Text: "减少", Checked: true},
			{Value: "1", Text: "增加"},
		},
		CommonIsShelvesList: []*ConstListInfo{
			{Value: "0", Text: "下架"},
			{Value: "1", Text: "上架", Checked: true},
		},
		CommonIsTextList: []*ConstListInfo{
			{Value: "0", Text: "否", Checked: true},
			{Value: "1", Text: "是"},
		},
		CommonIncDecList: []*ConstListInfo{
			{Value: "0", Text: "减少"},
			{Value: "1", Text: "增加"},
		},
		CommonAppEventType: []*ConstListInfo{
			{Value: "0", Text: "WEB页面"},
			{Value: "1", Text: "内部页面(小程序/APP内部地址)"},
			{Value: "2", Text: "外部小程序(同一个主体下的小程序appid)"},
			{Value: "3", Text: "跳转原生地图查看指定位置"},
			{Value: "4", Text: "拨打电话"},
		},
		CommonOrderAftersaleTypeList: []*ConstListInfo{
			{Value: "0", Text: "仅退款", Desc: "未收到货(未签收),协商同意前提下", Icon: "am-icon-random", Class: "am-fl"},
			{Value: "1", Text: "退款退货", Desc: "已收到货,需要退换已收到的货物", Icon: "am-icon-retweet", Class: "am-fr"},
		},
		CommonOrderAftersaleStatusList: []*ConstListInfo{
			{Value: "0", Text: "待确认"},
			{Value: "1", Text: "待退货"},
			{Value: "2", Text: "待审核"},
			{Value: "3", Text: "退款中"},
			{Value: "4", Text: "已完成"},
			{Value: "5", Text: "已拒绝"},
			{Value: "6", Text: "已取消"},
		},
		CommonOrderAftersaleRefundmentList: []*ConstListInfo{
			{Value: "0", Text: "原路退回"},
			{Value: "1", Text: "退至钱包"},
			{Value: "2", Text: "手动处理"},
		},
		CommonLogisticsTypeList: []*ConstListInfo{
			{Value: "1", Text: "物流快递"},
			{Value: "2", Text: "同城配送"},
			{Value: "3", Text: "虚拟发货"},
			{Value: "4", Text: "用户自提"},
		},
		CommonDeliveryModeList: []*ConstListInfo{
			{Value: "1", Text: "统一发货"},
			{Value: "2", Text: "分拆发货"},
		},
		CommonSiteTypeList: []*ConstListInfo{
			{Value: "0", Text: "快递"},
			{Value: "1", Text: "展示"},
			{Value: "2", Text: "自提"},
			{Value: "3", Text: "虚拟售卖"},
			{Value: "4", Text: "快递+自提", Checked: true}, //, "is_ext": 1
		},
		CommonOrderTypeList: []*ConstListInfo{
			{Value: "0", Text: "快递"},
			{Value: "1", Text: "展示"},
			{Value: "2", Text: "自提"},
			{Value: "3", Text: "虚拟销售"},
		},
		CommonAdminStatusList: []*ConstListInfo{
			{Value: "0", Text: "暂停"},
			{Value: "1", Text: "正常", Checked: true},
			{Value: "2", Text: "已离职"},
		},
		CommonMemberStatusList: []*ConstListInfo{
			{Value: "0", Text: "禁用"},
			{Value: "1", Text: "正常", Checked: true},
			{Value: "2", Text: "锁定"},
			{Value: "9", Text: "已注销"},
		},
		CommonMemberTypeList: []*ConstListInfo{
			{Value: "free", Text: "免费会员"},
			{Value: "pay", Text: "付费会员"},
		},
		CommonLevelPayRuleList: []*ConstListInfo{
			{Value: "1", Text: "1个月"},
			{Value: "3", Text: "3个月"},
			{Value: "6", Text: "6个月"},
			{Value: "12", Text: "1年"},
			{Value: "36", Text: "3年"},
			{Value: "60", Text: "5年"},
			{Value: "120", Text: "10年"},
		},
		CommonPayLogStatusList: []*ConstListInfo{
			{Value: "0", Text: "待支付", Checked: true},
			{Value: "1", Text: "已支付"},
			{Value: "2", Text: "已关闭"},
		},
		CommonCreateSourceList: []*ConstListInfo{
			{Value: "0", Text: "商家创建"},
			{Value: "1", Text: "系统预设"},
		},
		CommonTimezoneList: []*ConstListInfo{
			{Value: "Pacific/Pago_Pago", Text: "(标准时-11:00) 中途岛、萨摩亚群岛"},
			{Value: "Pacific/Rarotonga", Text: "(标准时-10:00) 夏威夷"},
			{Value: "Pacific/Gambier", Text: "(标准时-9:00) 阿拉斯加"},
			{Value: "America/Dawson", Text: "(标准时-8:00) 太平洋时间(美国和加拿大)"},
			{Value: "America/Creston", Text: "(标准时-7:00) 山地时间(美国和加拿大)"},
			{Value: "America/Belize", Text: "(标准时-6:00) 中部时间(美国和加拿大)、墨西哥城"},
			{Value: "America/Eirunepe", Text: "(标准时-5:00) 东部时间(美国和加拿大)、波哥大"},
			{Value: "America/Antigua", Text: "(标准时-4:00) 大西洋时间(加拿大)、加拉加斯"},
			{Value: "America/Argentina/Buenos_Aires", Text: "(标准时-3:00) 巴西、布宜诺斯艾利斯、乔治敦"},
			{Value: "America/Noronha", Text: "(标准时-2:00) 中大西洋"},
			{Value: "Atlantic/Cape_Verde", Text: "(标准时-1:00) 亚速尔群岛、佛得角群岛"},
			{Value: "Africa/Ouagadougou", Text: "(格林尼治标准时) 西欧时间、伦敦、卡萨布兰卡"},
			{Value: "Europe/Andorra", Text: "(标准时+1:00) 中欧时间、安哥拉、利比亚"},
			{Value: "Europe/Mariehamn", Text: "(标准时+2:00) 东欧时间、开罗，雅典"},
			{Value: "Asia/Bahrain", Text: "(标准时+3:00) 巴格达、科威特、莫斯科"},
			{Value: "Asia/Dubai", Text: "(标准时+4:00) 阿布扎比、马斯喀特、巴库"},
			{Value: "Asia/Kolkata", Text: "(标准时+5:00) 叶卡捷琳堡、伊斯兰堡、卡拉奇"},
			{Value: "Asia/Dhaka", Text: "(标准时+6:00) 阿拉木图、 达卡、新亚伯利亚"},
			{Value: "Indian/Christmas", Text: "(标准时+7:00) 曼谷、河内、雅加达"},
			{Value: "Asia/Shanghai", Text: "(标准时+8:00)北京、重庆、香港、新加坡"},
			{Value: "Australia/Darwin", Text: "(标准时+9:00) 东京、汉城、大阪、雅库茨克"},
			{Value: "Australia/Adelaide", Text: "(标准时+10:00) 悉尼、关岛"},
			{Value: "Australia/Currie", Text: "(标准时+11:00) 马加丹、索罗门群岛"},
			{Value: "Pacific/Fiji", Text: "(标准时+12:00) 奥克兰、惠灵顿、堪察加半岛"},
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
	Text    string `json:"text"`
	Value   string `json:"value"`
	Icon    string `json:"icon"`
	Class   string `json:"class"`
	Desc    string `json:"desc"`
	Flag    string `json:"flag"`
	Checked bool   `json:"checked"`
}

//获取公共常量数据
func MyConst() *ConstService {
	return LoadCommonConst()
}

//获取常量值的名称
func GetMyConstName(dataList []*ConstListInfo, value string) string {
	if val := GetMyConstItem(dataList, value); val != nil {
		return val.Text
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
