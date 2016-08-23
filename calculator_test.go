package gokatas

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestStringCalculator_Add_EmptyString_Return_Zero(t *testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("")
	assert.Equal(t, 0, result)
}

func TestStringCalculator_Add_SingleNumber_Return_ExactNumber(t * testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("2")
	assert.Equal(t, 2, result)
}

func TestStringCalculator_AddTwoNumbers_Comma_Delimited(t *testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("1,2")
	assert.Equal(t, 3, result)
	result, _ = s.Add("4,3")
	assert.Equal(t, 7, result)
}

func TestStringCalculator_AddTwoNumbers_Newline_Delimited(t *testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("1\n2")
	assert.Equal(t, 3, result)
	result, _ = s.Add("4\n3")
	assert.Equal(t, 7, result)
}

func TestStringCalculator_AddThreeNumbers_Newline_Delimited(t *testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("1\n2\n5")
	assert.Equal(t, 8, result)
	result, _ = s.Add("4\n3\n10")
	assert.Equal(t, 17, result)
}

func TestStringCalculator_AddNegativeNumbers_Return_Errors(t *testing.T) {
	s := new(StringCalculator)
	_, err := s.Add("-1\n2\n5")
	assert.Error(t, err, "Negatives not allowed")
}

func TestStringCalculator_AddNumber1000And3_Return3(t *testing.T) {
	s := new(StringCalculator)
	result, _ := s.Add("1000\n3")
	assert.Equal(t, 3, result)
}

func TestStringCalculator_AddWithDynamicDelimiter(t *testing.T) {
	s := new(StringCalculator)
	result, err := s.Add(`//***\n1***2***5`)

	if err != nil {
		fmt.Printf(err.Error())
	}

	assert.Equal(t, 8, result)
}
