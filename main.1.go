package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"regexp"

// 	"github.com/ivahaev/go-logger"
// )

// var (
// 	version = "0.0.0"
// )

// func main() {
// 	file, err := ioutil.ReadFile("main.1.go")
// 	if err != nil {
// 		fmt.Printf("Can't read file: %v", err.Error())
// 		return
// 	}
// 	rgxVersion := regexp.MustCompile(`\s+version\s+=\s+"\d+\.\d+\.\d+"`)
// 	logger.Debug(rgxVersion.Find(file))

// }
