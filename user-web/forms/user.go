package forms

type PasswordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"` //手机号码格式有规范可寻，自定义validator
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
}

type RegisterUserForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password  string `form:"password" json:"password" binding:"required,min=6,max=20"`
	NickName string `form:"nick_name" json:"nick_name" binding:"required,min=6,max=20"`
	Code   string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
