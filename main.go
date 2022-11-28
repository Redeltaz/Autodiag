package main

import (
    "fmt"
    "os"
    "regexp"
)

var validFulLArgs = [2]string{"target", "key"}
var validShortArgs = [2]string{"t", "k"}

type argument struct {
    name string
    value string
    type string
}

func main() {
    arguments := os.Args

    parseArgs(&arguments)
}

func parseArgs(argPtr *[]string) {
    args := *argPtr
    //parsedArgs = [...]argument{}
    regexPattern := regexp.MustCompile(`[^-]+`)

    //remove useless first argument
    args = append(args[:0], args[0+1:]...)

    for index, value := range(args) {
        if string(value[0]) == "-" && string(value[1]) != "-" {
            parsedValue := regexPattern.FindStringSubmatch(value)[0]
            
            newArg := argument{name: parsedValue, value: args[index + 1], type: "short"}
            fmt.Println(newArg)
        } else if string(value[0]) + string(value[1]) == "--" {
            fmt.Println(value, "full")
        }
        //fmt.Println(index)
    }

    //fmt.Println(args)
}
