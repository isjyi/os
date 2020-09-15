package request

type CaptchaStruct struct {
	Phone string `json:"phone" label:"手机号" binding:"required,len=11,phone"`
} //@name Captcha
