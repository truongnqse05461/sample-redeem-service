package service

import (
	"errors"
	"sample-redeem-service/internal/model"
)

type RedeemService struct{}

func NewRedeemService() *RedeemService {
	return &RedeemService{}
}

func (s *RedeemService) Redeem(uid, code string) (*model.RedeemResponse, error) {
	// Hardcoded rewards
	rewards := make(map[string]interface{})

	switch code {
	case "WELCOMEGIFT2025":
		rewards["Diamonds"] = 200
		rewards["Gold"] = 10000
		rewards["Weapon"] = "Beginner's Weapon"
		return &model.RedeemResponse{
			Success: true,
			Message: "Successfully redeemed WELCOMEGIFT2025",
			Rewards: rewards,
		}, nil
	case "HELLO":
		rewards["Diamonds"] = 100
		return &model.RedeemResponse{
			Success: true,
			Message: "Successfully redeemed HELLO",
			Rewards: rewards,
		}, nil
	default:
		return nil, errors.New("invalid or expired code")
	}
}
