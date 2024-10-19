package models

type CardDetails struct {
	CardNumber string `json:"card_number" binding:"required"`
}
