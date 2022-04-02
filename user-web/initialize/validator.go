package initialize

//func InitTrans(locale string) (err error)  {
//	//修改gin框架中的validator引擎属性，实现定制
//	if v,ok := binding.Validator.Engine().(*validator.Validate);ok {
//		//注册一个获取json的tag的自定义方法
//		v.RegisterTagNameFunc(func(field reflect.StructField) string {
//			name := strings.Split(field.Tag.Get("json"),",")[0]
//			if name== "-" {
//				return ""
//			}
//			return name
//		})
//		fmt.Println("-------------------------")
//		fmt.Println("-------------------------")
//		zhT := zh.New() //中文翻译器
//		enT := en.New() //英文翻译器
//		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
//		uni := ut.New(enT, zhT, enT)
//		global.Trans, ok = uni.GetTranslator(locale)
//		if !ok {
//			return fmt.Errorf("uni.GetTranslator(%s)", locale)
//		}
//
//		switch locale{
//		case "en":
//			err := en_translations.RegisterDefaultTranslations(v, global.Trans)
//			if err!=nil {
//				zap.S().Errorf("英文翻译失败")
//			}
//		case "zh":
//			err := zh_translations.RegisterDefaultTranslations(v, global.Trans)
//			if err!=nil {
//				zap.S().Errorf("中文翻译失败")
//			}
//		default:
//			err := en_translations.RegisterDefaultTranslations(v, global.Trans)
//			if err!=nil {
//				zap.S().Errorf("英文翻译失败")
//			}
//		}
//		return
//	}
//	return
//}

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"mxshop-api/user-web/global"
)

//var (
//	uni      *ut.UniversalTranslator
//	validate *validator.Validate
//	trans ut.Translator
//)

func InitTrans() (err error) {
	//注册翻译器
	Zh := zh.New()
	global.Uni = ut.New(Zh, Zh)

	global.Trans, _ = global.Uni.GetTranslator("zh")

	//获取gin的校验器
	validate := binding.Validator.Engine().(*validator.Validate)
	//注册翻译器
	if err:=zh_translations.RegisterDefaultTranslations(validate, global.Trans);err!=nil{
		return err
	}
	return nil
}


