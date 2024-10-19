package handlers

import (
	"kreval/models"
	services "kreval/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateCard(c *gin.Context) {
	var card models.CardDetails

	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	isValid, cardCompany, err := services.ValidateCard(card)
	if !isValid {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Card is valid",
		"card_company": cardCompany,
	})
}
