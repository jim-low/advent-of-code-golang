package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Steps:
// 1. read input.txt to get data
// 2. for each input line, combine the first and last digit to form a number
// 3. sum all lines' number to get the answer

func getData() []string {
  data, err := os.ReadFile("./day1/input.txt")
  if err != nil {
    panic(err)
  }

  return strings.Split(string(data), "\n")
}

func extractNumber(token string) int64 {
  if len(token) == 0 {
    return 0
  }

  var firstDigit byte
  var secondDigit byte

  for i := 0; i < len(token); i++ {
    var digit = token[i]

    if unicode.IsDigit(rune(digit)) {
      firstDigit = digit
      break;
    }
  }

  for i := len(token) - 1; i >= 0; i-- {
    var digit = token[i]

    if unicode.IsDigit(rune(digit)) {
      secondDigit = digit
      break;
    }
  }

  var resultingNumber = fmt.Sprintf("%c%c", firstDigit, secondDigit)
  var result, err = strconv.ParseInt(resultingNumber, 10, 64)
  if err != nil {
    panic(err)
  }

  return result
}

func Trebuchet() int64 {
  var data = getData()

  var sum int64 = 0
  for i := 0; i < len(data); i++ {
    sum += extractNumber(data[i])
  }

  return sum
}
