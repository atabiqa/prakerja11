package configs

import "atabiqa/prakerja11/models/user"

func initMigrate() {
	DB.AutoMigrate(&user.User{})
}
