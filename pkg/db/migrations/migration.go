package migrations

import "github.com/isjyi/os/pkg/db"

func init() {

}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
