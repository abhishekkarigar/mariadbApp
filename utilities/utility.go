package utilities

import (
	_ "github.com/Shopify/sarama"
)

const Topic = "test"

//var GormDatabse = func() *gorm.DB {
//	conf := config.ReadEnv()
//	gormConfig, err := gormdb.NewGormConfig(conf)
//	if err != nil {
//		return nil
//	}
//	return gormConfig
//}
