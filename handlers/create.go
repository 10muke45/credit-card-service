package handlers

import (
	services "kreval/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GenerationRequest struct {
	CardCompany string `json:"card_company" binding:"required"`
}

func GenerateCard(c *gin.Context) {
	var req GenerationRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	cardNumber, err := services.GenerateCardNumber(req.CardCompany)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Card number generated successfully",
		"card_number": cardNumber,
	})
}
