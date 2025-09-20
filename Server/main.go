package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"smartcontentsummarizer/Server/controller"
)

func main() {
	r := gin.Default()

	r.POST("/summarize", controller.ReqTextToAi)
	err := godotenv.Load("C:/Users/Win10_Game_OS/IdeaProjects/Smart-Content-Summarizer/Server/.env")
	if err != nil {
		log.Println("Ошибка загрузки .env файла:", err)
	}
	r.Run(":3030") // используй свободный порт
}
