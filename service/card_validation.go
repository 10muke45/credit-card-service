package services

import (
	"errors"
	"kreval/models"
	"kreval/utils"
)

func ValidateCard(card models.CardDetails) (bool, string, error) {
	if !utils.LuhnCheck(card.CardNumber) {
		return false, "", errors.New("Invalid card number")
	}

	cardCompany := getCardCompany(card.CardNumber)
	if cardCompany == "Unknown" {
		return true, "Unknown card issuer", nil
	}

	return true, cardCompany, nil
}

func getCardCompany(cardNumber string) string {
	if len(cardNumber) == 16 && cardNumber[0] == '4' {
		return "Visa"
	} else if len(cardNumber) == 16 && (cardNumber[:2] == "51" || cardNumber[:2] == "52" || cardNumber[:2] == "53" || cardNumber[:2] == "54" || cardNumber[:2] == "55") {
		return "MasterCard"
	} else if len(cardNumber) == 15 && (cardNumber[:2] == "34" || cardNumber[:2] == "37") {
		return "American Express"
	} else if len(cardNumber) == 16 && (cardNumber[:4] == "6011" || (cardNumber[:3] >= "644" && cardNumber[:3] <= "649") || cardNumber[:2] == "65") {
		return "Discover"
	}

	return "Unknown"
}
