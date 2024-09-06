package connections

import (
	"fmt"
	"log"
	"product-ms/apps/views"
	"product-ms/libs/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SetMySQL struct {
	DBConfig configs.DatabaseConfig
}

func (m *SetMySQL) SetMySQL() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.DBConfig.Username,
		m.DBConfig.Password,
		m.DBConfig.Host,
		m.DBConfig.Port,
		m.DBConfig.Name,
	)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	})
	if err != nil {
		log.Fatalf("Failed to connect to MySQL! err: %v", err)
	}

	if err = database.AutoMigrate(&views.Product{}); err != nil {
		log.Fatalf("Failed to auto migrate user! err: %v", err)
	}

	return database
}
