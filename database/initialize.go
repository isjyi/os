package database

func Setup(driver string) {
	if driver == "mysql" {
		var db = new(Mysql)
		db.Setup()
	}
}
