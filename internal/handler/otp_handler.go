package handler

import (
	"net/http"
	"sample-redeem-service/internal/model"
	"sample-redeem-service/internal/service"

	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
	service *service.OTPService
}

func NewOTPHandler(s *service.OTPService) *OTPHandler {
	return &OTPHandler{service: s}
}

func (h *OTPHandler) SendOTP(c *gin.Context) {
	var req model.SendOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.SendOTP(req.Phone); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req model.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid, err := h.service.VerifyOTP(req.Phone, req.OTP)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if valid {
		c.JSON(http.StatusOK, gin.H{"message": "Identity verification successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
	}
}
