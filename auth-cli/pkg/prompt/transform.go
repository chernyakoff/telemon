package prompt

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/dongri/phonenumber"
)

type Transform struct {
}

func (t *Transform) Phone(input string) string {
	result := phonenumber.Parse(input, "RU")
	if strings.Index(result, "+7") == 0 {
		re := regexp.MustCompile(`^\+7$`)
		result = re.ReplaceAllString(result, "8")
	}
	return result
}

func (t *Transform) AppId(input string) int {
	result, _ := strconv.Atoi(input)
	return result
}
