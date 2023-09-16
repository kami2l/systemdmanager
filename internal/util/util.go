package util

import "regexp"

func TrimNonAlphaRubbish(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z\\s-.]+")
	return reg.ReplaceAllString(input, "")
}
