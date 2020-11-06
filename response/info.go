package response

type InfoResponse struct {
	//头像
	Aatar string `json:"avatar"`
	//手机号
	UserName string `json:"userName"`
	//id
	UserID uint64 `json:"userId"`
	//昵称
	Name string `json:"name"`
	//角色名称
	Roles []string `json:"roles"`
	//角色权限
	Permissions []string `json:"permissions"`
} //@name InfoResponse
