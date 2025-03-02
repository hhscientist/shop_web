package forms

type PassWordLoginForm struct {
	Mobile    string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	Captcha   string `form:"captcha" json:"captcha"`
	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`
	Answer    string `form:"answer" json:"answer" binding:"required"`
}

type RegisterForm struct {
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	PassWord string `form:"password" json:"password" binding:"required,min=3,max=20" `
	Code     string `form:"code" json:"code" binding:"required,min=6,max=6"`
}
