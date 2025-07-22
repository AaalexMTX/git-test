package app

import "fmt"

func Start() {
	InitMysql()
	InitRedis()
}

func InitMysql() {
	fmt.Println("InitMysql")
}

func InitRedis() {
	fmt.Println("InitRedis")
}
