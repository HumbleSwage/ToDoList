package cache

import (
	"ToDoList/pkg/utils"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"strconv"
)

// RedisClient
var (
	RedisClient *redis.Client
	RedisAddr   string
	RedisPw     string
	RedisDbName string
)

func init() {
	file, err := ini.Load("conf/conf.ini")
	if err != nil {
		utils.LogrusObj.Info("config file err, check your config's path")
		panic(err)
	}
	LoadRedis(file)
	Redis()
}

func LoadRedis(file *ini.File) {
	section := file.Section("redis")
	RedisAddr = section.Key("RedisAddr").String()
	RedisPw = section.Key("RedisPw").String()
	RedisDbName = section.Key("RedisDbName").String()
}
func Redis() {
	db, _ := strconv.ParseUint(RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisAddr,
		Password: RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		utils.LogrusObj.Info(err)
		panic(err)
	}
	RedisClient = client
}
