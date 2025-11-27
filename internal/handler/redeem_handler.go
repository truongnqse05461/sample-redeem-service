package handler

import (
	"net/http"
	"sample-redeem-service/internal/model"
	"sample-redeem-service/internal/service"

	"github.com/gin-gonic/gin"
)

type RedeemHandler struct {
	service *service.RedeemService
}

func NewRedeemHandler(s *service.RedeemService) *RedeemHandler {
	return &RedeemHandler{service: s}
}

func (h *RedeemHandler) Redeem(c *gin.Context) {
	var req model.RedeemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.Redeem(req.UID, req.Code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
