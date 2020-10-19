// @title OS Example API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name OS API
// @contact.email zhangbiao19931203@gmail.com

// @host localhost:8000
// @BasePath /api
// @query.collection.format multi

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @Param Authorization header string false "Bearer 用户令牌"
package main

import "github.com/isjyi/os/cmd"

func main() {
	cmd.Execute()
}
