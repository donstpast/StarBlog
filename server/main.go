package main

import (
	"starblog/model"
	"starblog/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
