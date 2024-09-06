package main

import (
	"fmt"
	"log"
	"product-ms/apps/controllers/routers"
	"product-ms/libs/configs"
	"product-ms/libs/connections"
)

func main() {
	configs.InitConfig()

	db := connections.SetMySQL{DBConfig: configs.DB}
	mysql := db.SetMySQL()
	mysqlConn, _ := mysql.DB()
	defer mysqlConn.Close()

	r := routers.SetupRouter(mysql)

	if err := r.Listen(fmt.Sprintf("%s:%d", configs.App.Host, configs.App.Port)); err != nil {
		log.Fatalln("Error start product management system, err: " + err.Error())
	}
}
