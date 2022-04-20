package main

import (
	"diy_project/config"
	"diy_project/models"
	"diy_project/server"
	"diy_project/service"
)

func Initialize() {
	service.EstablishGormConnection()
	models.CreateModels()
}

func main() {
	config.LoadConfiguration("config.json")
	service.CheckDatabase()
	Initialize()
	server.StartServer()

}
