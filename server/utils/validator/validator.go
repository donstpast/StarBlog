package validator

import (
	"github.com/go-playground/locales/zh_Hans_CN"
	unitrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"reflect"
	"starblog/utils/errmsg"
)

// Validate 函数用于对数据进行验证
// 参数：
// - data: 待验证的数据，类型为任意类型 any
// 返回值：
// - string: 错误信息，如果验证通过则返回空字符串
// - int: 验证结果，成功为 errmsg.SUCCESS，失败为 errmsg.ERROR
func Validate(data any) (string, int) {
	// 实例化一个 validator 对象
	validate := validator.New()
	// 使用 zh_Hans_CN 翻译器创建一个 unitrans 对象
	uni := unitrans.New(zh_Hans_CN.New())
	// 获取翻译器对象
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	// 注册默认翻译和翻译器到 validator 中
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal(err)
	}
	// 注册 tag 名称的翻译方法
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("label")
		return label
	})
	// 对数据进行验证
	err = validate.Struct(data)
	if err != nil {
		// 遍历错误信息，返回第一个错误的翻译信息和错误码
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	// 如果验证通过，返回空字符串和成功的错误码
	return "", errmsg.SUCCESS
}
