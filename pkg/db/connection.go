package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	config "framework-auto-aprov-go/pkg/config"
	domain "framework-auto-aprov-go/pkg/domain"

	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	// sqlInfo := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	// db, dbErr := gorm.Open(sqlserver.Open(sqlInfo), &gorm.Config{
	// 	SkipDefaultTransaction: true,
	// })

	db.AutoMigrate(&domain.User{})

	db.AutoMigrate(&domain.Rol{})
	db.AutoMigrate(&domain.Rol_Usuario{})

	db.AutoMigrate(&domain.Auto_Aprov_Omnitok{})

	//db.Model(&domain.Market{}).Association("cx_estado")

	if err := db.SetupJoinTable(&domain.User{}, "Roles", &domain.Rol_Usuario{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table")
	}

	//	db.Model(&domain.User{}).Association("Roles")

	return db, dbErr
}
