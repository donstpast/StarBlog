package errmsg

import "errors"

// 错误信息处理
// jwt中Token错误
var (
	TokenExpired     = errors.New("token已过期,请重新登录")
	TokenNotValidYet = errors.New("token无效,请重新登录")
	TokenMalformed   = errors.New("token不正确,请重新登录")
	TokenInvalid     = errors.New("这不是一个token,请重新登录")
)

// 定义一些错误常量（最常见的时候状态码）
const (
	SUCCESS = 200
	ERROR   = 500

	// 之后定义一些错误代码，我们规定如下:

	//code在(1000,2000)为用户模块错误:
	ERROR_USER_NOT_EXIST                = 1001 //用户不存在
	ERROR_USERNAME_USED                 = 1002 //用户名已存在
	ERROR_PASSWORD_WRONG                = 1003 //密码错误
	ERROR_USERNAME_OR_PASSWORD_IS_EMPTY = 1004 //用户名或者密码为空
	ERROR_USER_NOT_ADMIN                = 1005 //用户不是管理员
	ERROR_TOKEN_NOT_EXIST               = 1006 //token不存在
	ERROR_TOKEN_OVERTIME                = 1007 //token过期
	ERROR_TOKEN_WRONG                   = 1008 //token错误
	ERROR_TOKEN_FORMAT_WRONG            = 1009 //token格式错误
	ERROE_ROLE_WRONG                    = 1010 //role错误
	//code在(2000,3000)为文章模块错误
	ERROR_ARTICLE_NOT_EXIST = 2001 //文章不存在

	//code在(3000,4000)为分类模块错误
	ERROR_CATEGORY_NOT_EXIST = 3001 //分类不存在
	ERROR_CATEGORY_USED      = 3002 //分类已存在
	ERROR_CATEGORY_IS_EMPTY  = 3003 //分类名称不能为空

	//code在(4000,5000)为标签模块错误

	//code在(5000,6000)为评论模块错误
	ERROR_COMMENT_NOT_EXIST         = 5001 //评论不存在
	ERROR_COMMENT_IS_EMPTY          = 5002 //评论为空
	ERROR_COMMENT_EMAIL_IS_EMPTY    = 5003 //邮箱不能为空
	ERROR_COMMENT_NICKNAME_IS_EMPTY = 5004 //昵称不能为空

	//code在(6000，7000)为友链模块错误
	ERROR_FRIENDS_NOT_EXIST      = 6001 //友链不存在
	ERROR_FRIENDS_LINKS_IS_EMPTY = 6002 //友链为空

)

// 定义一个map
// int处为上方状态码常量，string处为需要抛出的错误信息
var codeMsg = map[int]string{
	//通用状态码
	SUCCESS: "OK",
	ERROR:   "FAIL",
	//用户模块错误码
	ERROR_USER_NOT_EXIST:                "用户不存在",
	ERROR_USERNAME_USED:                 "用户名已存在",
	ERROR_PASSWORD_WRONG:                "密码错误",
	ERROR_USERNAME_OR_PASSWORD_IS_EMPTY: "用户名或密码为空",
	ERROR_USER_NOT_ADMIN:                "用户没有管理员权限",
	ERROR_TOKEN_NOT_EXIST:               "TOKEN不存在",
	ERROR_TOKEN_OVERTIME:                "TOKEN已过期",
	ERROR_TOKEN_WRONG:                   "TOKEN错误",
	ERROR_TOKEN_FORMAT_WRONG:            "TOKEN格式错误",
	ERROE_ROLE_WRONG:                    "权限错误",
	//文章模块错误码
	ERROR_ARTICLE_NOT_EXIST: "文章不存在",
	//分类模块错误码
	ERROR_CATEGORY_NOT_EXIST: "分类不存在",
	ERROR_CATEGORY_USED:      "分类已存在",
	ERROR_CATEGORY_IS_EMPTY:  "分类名称不能为空",
	//标签模块错误码

	//评论模块错误码
	ERROR_COMMENT_NOT_EXIST:         "评论不存在",
	ERROR_COMMENT_IS_EMPTY:          "评论不能为空",
	ERROR_COMMENT_EMAIL_IS_EMPTY:    "邮箱不能为空",
	ERROR_COMMENT_NICKNAME_IS_EMPTY: "昵称不能为空",

	//友链模块错误码
	ERROR_FRIENDS_NOT_EXIST:      "友链不存在",
	ERROR_FRIENDS_LINKS_IS_EMPTY: "友链为空",
}

// GetErrMsg 用来输出错误信息的函数,传入一个int型，返回一个string类型
func GetErrMsg(code int) string {
	return codeMsg[code]
}
