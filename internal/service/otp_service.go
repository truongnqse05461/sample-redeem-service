package service

import (
	"errors"
	"log"
)

type OTPService struct{}

func NewOTPService() *OTPService {
	return &OTPService{}
}

func (s *OTPService) SendOTP(phone string) error {
	// Mock implementation: Just log the OTP request
	log.Printf("Sending OTP to %s", phone)
	return nil
}

func (s *OTPService) VerifyOTP(phone, otp string) (bool, error) {
	// Mock implementation: Check if OTP is "168168"
	if otp == "168168" {
		return true, nil
	}
	return false, errors.New("invalid OTP")
}
