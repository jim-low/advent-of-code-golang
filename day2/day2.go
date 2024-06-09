package day2

import (
	utils "advent-of-code/utils"
	"fmt"
	// "fmt"
	"regexp"
	"strings"
)

var numbersText = map[string] int {
  "one": 1,
  "1": 1,

  "two": 2,
  "2": 2,

  "three": 3,
  "3": 3,

  "four": 4,
  "4": 4,

  "five": 5,
  "5": 5,

  "six": 6,
  "6": 6,

  "seven": 7,
  "7": 7,

  "eight": 8,
  "8": 8,

  "nine": 9,
  "9": 9,
}

func extractNumber(token string) int64 {
  if len(token) == 0 {
    return 0
  }

  r := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9)")
  matches := r.FindAllString(strings.ToLower(token), -1)

  fmt.Printf("%s\n", token)
  fmt.Printf("%v\n", matches)

  if len(matches) == 0 {
    return 0
  }

  if len(matches) == 1 {
    return int64(numbersText[matches[0]])
  }

  var firstDigit = numbersText[matches[0]]
  var secondDigit = numbersText[matches[len(matches) - 1]]
  var answer = (firstDigit * 10) + secondDigit
  return int64(answer)
}

func Trebuchet2() int64 {
  data := utils.GetDataFromFile("./day2/input.txt")

  var sum int64 = 0
  for i := 0; i < len(data); i++ {
    extractedNumber := extractNumber(data[i])
    println(extractedNumber)
    sum += extractedNumber
  }

  return sum
}
