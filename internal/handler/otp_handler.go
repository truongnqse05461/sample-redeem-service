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
		c.JSON(http.StatusBadRequest, model.SendOTPResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	otp, err := h.service.SendOTP(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.SendOTPResponse{
			Success: false,
			Message: "Failed to generate OTP",
		})
		return
	}

	c.JSON(http.StatusOK, model.SendOTPResponse{
		Success: true,
		OTP:     otp,
		Message: "OTP generated successfully",
	})
}

func (h *OTPHandler) VerifyOTP(c *gin.Context) {
	var req model.VerifyOTPRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.VerifyOTPResponse{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	valid, err := h.service.VerifyOTP(req.Email, req.OTP)
	if err != nil || !valid {
		msg := "Invalid OTP"
		if err != nil {
			msg = err.Error()
		}
		c.JSON(http.StatusOK, model.VerifyOTPResponse{ // User requested specific body, usually 200 OK with success: false is preferred for logic errors in this style
			Success: false,
			Message: msg,
		})
		return
	}

	c.JSON(http.StatusOK, model.VerifyOTPResponse{
		Success: true,
		Message: "Identity verification successful",
	})
}
