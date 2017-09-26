package utils

 type QYError struct {
	 Code int "json: code"
	 //Message string "json: message"
 }

var QYErrorMap = map[int]string{
	//system
	10001: "签名错误",
	10002: "请求已经过期",
	10003: "签名参数没有",
	10004: "没有权限访问",
    //user
	20000: "缺少参数",
	20100: "用户名或密码不对",
	20101: "没有这个手机号码的用户",
	20102: "用户已经注册，请直接登录",
	20103: "验证码不对",
	20104: "已经具备家长身份",
	20105: "该手机号码的用户不存在",
	20106: "手机号码格式不对",
    //seller
	20200: "已经申请过,不能重复申请",
	20201: "尚未申请商户",
	20202: "名称不能为空",
	20203: "商家描述不能为空",
	20204: "地址不能为空",
	20205: "真实姓名不能为空",
	20206: "身份证号不能为空",
    //qa
	20401: "小孩不存在",
	20402: "该问题已经关闭",
	//user_token
	20501: "找不到关联信息，请重新登录",
	20502: "图片上传失败",
	20503: "尚未绑定",
	20504: "用户不存在",
	//teacher
	20601: "找不到老师",
	20602: "无效的用户",
	20603: "用户名或密码不对",
	20604: "该手机号码的用户不存在",
	20605: "用户已经激活，请重新登录",
	20606: "验证码不对",
	//clase
	20701: "找不到班级",
	20702: "您没有权限",
	//product
	20801: "没有该产品，请刷新重试",
	20802: "产品已经下架",
	20803: "产品已经过了有效期",
	//order
	20901: "订单生成失败，请刷新重试",
	20902: "您的时间还很充裕，订单生成失败",
	20903: "订单结果验证失败",
	20904: "订单结果验证方法不对",
	//version
	21001: "您的版本太低，请更新到最新版",
	//message
	21101: "通知不存在",
	21102: "您没有权限删除通知",
	//album
	21201: "相册不存在",
	21202: "您没有权限删除相册",
	21203: "您没有权限评论相册",
	21204: "照片不存在",
	21205: "发布失败",
	//kid
	21301: "小孩已存在",
	21302: "小孩不存在",
	21303: "找不到这个孩子",
	//discussion
	21400: "评论类型不存在",
	//unconfirm_kids
	21500: "没有这个孩子",
}

func(qyEroor *QYError)GetCode() string{
	return QYErrorMap[qyEroor.Code]
}


func(qyEroor QYError)Error() string{
	return QYErrorMap[qyEroor.Code]
}