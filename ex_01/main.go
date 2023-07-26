package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hoangtk0100/dc-go-23/ex_01/util"
)

func main() {
	args := os.Args[1:]
	if err := util.ValidateArgs(args); err != nil {
		log.Fatalln(err)
	}

	length := len(args)
	firstName := args[0]
	lastName := args[1]
	countryCode := args[length-1]
	middleName := ""
	if length > util.MinArgsLength {
		middleName = strings.Join(args[2:length-1], " ")
	}

	formatter := util.NewNameFormatter(countryCode)
	output := formatter.Format(firstName, lastName, middleName)

	fmt.Println("Output:", output)
}
