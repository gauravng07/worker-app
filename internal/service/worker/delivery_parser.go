package worker

import (
	"regexp"
)

const deliveryRegex = `\b((0?[1-9]|1[012])([:.][0-5][0-9])?(\s?[AP]M)|([01]?[0-9]|2[0-3])([:.][0-5][0-9]))\b`

func parseDelivery(delivery string) (string, string) {
	padString := func(s string) string {
		if len(s) == 3 {
			return "0" + s
		}
		return s
	}
	regex := regexp.MustCompile(deliveryRegex)
	matches := regex.FindAllString(delivery, -1)
	return padString(matches[0]), padString(matches[1])
}

func postCodeDeliveryTime(f1 string, f2 string, t1 string, t2 string) bool {
	p1, _  := parseTime(f1)
	p2, _ := parseTime(f2)

	d1, _ := parseTime(t1)
	d2, _ := parseTime(t2)

	k1, k2 := p1.Sub(*d1).Hours(), p2.Sub(*d2).Hours()
	if k1 >= 0 && k2 <= 0 {
		return true
	}
	return false
}
