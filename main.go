package main

import (
	"merchant-service/controllers"
	"github.com/MuhammadSuryono1997/module-go/base/http"
	"os"
)

func main()  {
	server := http.CreateHttpServer()

	server.POST("register", controllers.Register)
	server.POST("resend-otp", controllers.ResendOtp)
	port := os.Getenv("PORT")
	if port == "" {
		server.Run()
	}
	server.Run(":" + port)
}