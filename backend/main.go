package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	if err := createDB(); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                                                                        // Permite todas as origens (qualquer IP)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                                  // Permite todos os métodos HTTP
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With", "X-Custom-Header"}, // Permite todos os tipos de cabeçalhos, incluindo 'multipart/form-data'
		ExposeHeaders:    []string{"Content-Length"},                                                                           // Permite expor cabeçalhos, se necessário
		AllowCredentials: true,                                                                                                 // Permite credenciais (se necessário)
		MaxAge:           12 * 3600,                                                                                            // O tempo que o navegador deve armazenar o CORS (em segundos)
	}))
	r.POST("/files", authMiddleware(), uploadFiles)
	r.GET("/files", authMiddleware(), listFiles)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
