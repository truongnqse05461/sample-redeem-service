package model

type SendOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type SendOTPResponse struct {
	Success bool   `json:"success"`
	OTP     string `json:"otp"`
	Message string `json:"message"`
}

type VerifyOTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

type VerifyOTPResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
