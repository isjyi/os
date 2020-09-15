package request

// user register struct
type RegisterStruct struct {
	Phone    string `json:"phone" label:"手机号" binding:"required,len=11,phone"`   // 手机号
	Password string `json:"passwrod" label:"密码" binding:"required,min=8,max=20"` // 密码
	NickName string `json:"nickname" binding:"omitempty,min=6,max=12"`           // 用户昵称
	Code     string `json:"code" binding:"required,len=4"`                       //验证码
} //@name Register

// user login struct
type LoginStruct struct {
	Phone    string `json:"phone" label:"手机号" binding:"required,len=11"`          // 手机号
	Password string `json:"passwrod" label:"密码" binding:"omitempty,min=8,max=20"` //密码
	Code     int    `json:"code" label:"验证码" binding:"required_without=Password"` //验证码
} //@name Login
