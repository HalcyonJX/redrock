package main

import (
	"redrock/work6/api"
	"redrock/work6/dao"
)

func main() {
	dao.InitDB()
	api.InitRouter()
}
