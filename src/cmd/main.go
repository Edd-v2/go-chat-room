package main

import (
	"go-chat-room/src/chat"
	"go-chat-room/src/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting GoLang Chat Room Server .. ")

	config.LoadConfiguration()

	code := init_server()

	if code != 0 {
		log.Fatal("Failed to start server!")
	}
	log.Println("Starting GoLang Chat Room Server .. ")

}

func init_server() (code int) {

	router := gin.Default()

	router.GET(config.AppConfig.BasePath+"/join/:username", chat.JoinChat)
	router.POST(config.AppConfig.BasePath+"/send", chat.SendMessage)
	router.GET(config.AppConfig.BasePath+"/messages", chat.GetMessages)

	err := router.Run(":" + config.AppConfig.AppPort)
	if err != nil {
		return -1
	}

	return 0
}
