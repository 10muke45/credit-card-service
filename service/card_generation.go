package services

import (
	"fmt"
	"kreval/utils"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCardNumber(cardCompany string) (string, error) {
	switch cardCompany {
	case "Visa":
		return generateNumberWithPrefix("4", 16), nil
	case "MasterCard":
		prefixes := []string{"51", "52", "53", "54", "55"}
		return generateNumberWithPrefix(prefixes[rand.Intn(len(prefixes))], 16), nil
	case "American Express":
		prefixes := []string{"34", "37"}
		return generateNumberWithPrefix(prefixes[rand.Intn(len(prefixes))], 15), nil
	case "Discover":
		return generateNumberWithPrefix("6011", 16), nil
	default:
		return "", fmt.Errorf("Card company not supported")
	}
}

func generateNumberWithPrefix(prefix string, length int) string {
	rand.Seed(time.Now().UnixNano())

	number := prefix
	for i := 0; i < length-len(prefix)-1; i++ {
		number += strconv.Itoa(rand.Intn(10))
	}

	return number + strconv.Itoa(utils.GenerateLuhnChecksum(number))
}
