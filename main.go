// @title OS Example API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name OS API
// @contact.email zhangbiao19931203@gmail.com

// @host localhost:8080
// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
package main

import (
	"github.com/isjyi/os/cmd"
)

func main() {
	cmd.Execute()
}
