package main

import (
	"log"
	"sample-redeem-service/internal/handler"
	"sample-redeem-service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Services
	otpService := service.NewOTPService()
	redeemService := service.NewRedeemService()

	// Initialize Handlers
	otpHandler := handler.NewOTPHandler(otpService)
	redeemHandler := handler.NewRedeemHandler(redeemService)

	// Setup Router
	r := gin.Default()

	// Routes
	api := r.Group("/api")
	{
		otp := api.Group("/otp")
		{
			otp.POST("/send", otpHandler.SendOTP)
			otp.POST("/verify", otpHandler.VerifyOTP)
		}
		api.POST("/redeem", redeemHandler.Redeem)
	}

	// Start Server
	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
