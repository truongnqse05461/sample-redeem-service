package model

type RedeemRequest struct {
	UID  string `json:"uid" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type RedeemResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Rewards map[string]interface{} `json:"rewards,omitempty"`
}
