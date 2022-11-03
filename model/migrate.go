package model

func migrate() {
	//自动迁移
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{}, &Task{})
	if err != nil {
		return
	}
	//DB.Model(&Task{}).AddForeignKey("uid", "User(id)", "CASCADE", "CASCADE")
}
