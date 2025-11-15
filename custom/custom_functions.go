// here we define some custom functions used in the parser.
// also, that's the place where you write your own functions to adjust data as you want.

package custom

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/jexlor/cs2scraper/scraper/config"
)

func ParsePrice(raw, stattRaw string) config.Price {
	raw = strings.TrimSpace(raw)
	stattrakRaw := strings.TrimSpace(stattRaw)

	var currency string
	if strings.Contains(raw, "$") || strings.Contains(stattrakRaw, "$") {
		currency = "USD"
	}

	price := config.Price{
		PriceString:         raw,
		PriceStattrakString: stattrakRaw,
		Currency:            currency,
		Min:                 config.PriceValue{Value: 0, StattrakValue: 0, Unit: currency},
		Max:                 config.PriceValue{Value: 0, StattrakValue: 0, Unit: currency},
		UpdatedAt:           time.Now().Format(time.RFC3339),
	}

	re := regexp.MustCompile(`[\d.,]+`)
	if raw != "" {
		parts := strings.Split(raw, "-")
		if len(parts) == 1 {
			v, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[0]), ",", ""), 64)
			price.Min.Value = v
			price.Max.Value = v
		} else if len(parts) >= 2 {
			v1, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[0]), ",", ""), 64)
			v2, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[1]), ",", ""), 64)
			price.Min.Value = v1
			price.Max.Value = v2
		}
	}

	if stattrakRaw != "" {
		parts := strings.Split(stattrakRaw, "-")
		if len(parts) == 1 {
			v, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[0]), ",", ""), 64)
			price.Min.StattrakValue = v
			price.Max.StattrakValue = v
		} else if len(parts) >= 2 {
			v1, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[0]), ",", ""), 64)
			v2, _ := strconv.ParseFloat(strings.ReplaceAll(re.FindString(parts[1]), ",", ""), 64)
			price.Min.StattrakValue = v1
			price.Max.StattrakValue = v2
		}
	}

	return price
}

func SpecialMark(weapon string) string {
	keywords := []string{"Knife", "Gloves", "Wraps"}
	for _, keyword := range keywords {
		if strings.Contains(weapon, keyword) {
			return "★ " + weapon
		}
	}
	return weapon
}
