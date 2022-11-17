package main

import (
	"redrock/work5/api"
	"redrock/work5/dao"
)

func main() {
	dao.ReadUser()
	api.InitRouter()
}
