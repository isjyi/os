package tools

type (
	Mode string
)

const (
	ModeDev  Mode = "dev"
	ModeTest Mode = "test"
	ModeProd      = "prod"
	Mysql         = "mysql"
	Sqlite        = "sqllite3"
)
