package main

import (
	"flag"
	"fmt"
	"github.com/hoangtk0100/dc-go-23/ex_02/constant"
	"github.com/hoangtk0100/dc-go-23/ex_02/sorting"
	validator "github.com/hoangtk0100/dc-go-23/ex_02/val"
	"log"
	"strings"
)

func main() {
	intFlag := flag.Bool("int", false, "Sort integers")
	floatFlag := flag.Bool("float", false, "Sort floats")
	stringFlag := flag.Bool("string", false, "Sort strings")
	mixFlag := flag.Bool("mix", false, "Sort mixed inputs (integers, floats, strings)")

	flag.Parse()

	inputType, err := constant.GetInputType(intFlag, floatFlag, stringFlag, mixFlag)
	if err != nil {
		log.Fatal(err)
	}

	values, err := validator.ValidateInputs(inputType, flag.Args())
	if err != nil {
		log.Fatal(err)
	}

	sorting.ASC(sorting.InputSlice(values))

	fmt.Println("Output:", strings.Trim(fmt.Sprintf("%v", values), "[]"))
}
