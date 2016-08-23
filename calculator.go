package gokatas

import (
	"strings"
	"strconv"
	"errors"
	"regexp"
)

type StringCalculator struct {

}

func (s *StringCalculator) Add(numbers string) (int, error) {

	if len(numbers) == 0 {
		return 0, nil
	}

	var nums []string
	if strings.Contains(numbers, `//`) {
		re := regexp.MustCompile(`[^\/\/\\n]+`)
		delimiter := re.FindString(numbers)
		re2 := regexp.MustCompile(`[^\/\/\\n]+\d`)
		nb := re2.FindString(numbers)
		nums = strings.Split(nb, delimiter)
	} else if strings.Contains(numbers, `,`) {
		nums = strings.Split(numbers, `,`)
	} else if strings.Contains(numbers, "\n") {
		nums = strings.Split(numbers, "\n")
	} else {
		i, err := strconv.Atoi(numbers)
		if err != nil {
			return 0, err
		}

		if i < 0 {
			return 0, errors.New("Cannot add negative number")
		}

		return i, nil
	}
	result := 0

	for _, n := range nums {
		x, err := strconv.Atoi(n)
		if err != nil {
			return 0, err
		}

		if x < 0 {
			return 0, errors.New("Negatives not allowed")
		}

		if x < 1000 {
			result += x
		}

	}

	return result, nil
}


