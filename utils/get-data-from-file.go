package utils

import (
	"os"
	"strings"
)

func GetDataFromFile(path string) []string {
  data, err := os.ReadFile(path)
  if err != nil {
    panic(err)
  }

  return strings.Split(string(data), "\n")
}
