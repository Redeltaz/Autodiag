package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

var validFulLArgs = [2]string{"target", "key"}
var validShortArgs = [2]string{"t", "k"}

type argument struct {
    name string
    value string
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

    for _, value := range(args) {
        var parsedKey string

        if string(value[0]) == "-" && string(value[1]) != "-" {
            parsedKey = regexPattern.FindStringSubmatch(value)[0]

            switch parsedKey {
            case "t":
                parsedKey = "target"
            case "k":
                parsedKey = "key"
            default:
                errorMsg := "Unknow argument -" + parsedKey
                sendError(&errorMsg)
            }
        } else if string(value[0]) + string(value[1]) == "--" {
            fmt.Println(len(value))
            os.Exit(0)
            if len(value) > 2 && string(value[1]) == "-" {
                errorMsg := "Argument format invalid "
                sendError(&errorMsg)
                fmt.Println("wtf")
            }

            parsedKey = regexPattern.FindStringSubmatch(value)[0]

            isParameterValid := false
            for _, v := range(validFulLArgs) {
                if v == parsedKey {
                    isParameterValid = true
                }
            }
            if !isParameterValid {
                errorMsg := "Unknow argument -" + parsedKey
                sendError(&errorMsg)
            }
        }

        //fmt.Println(parsedKey)
    }

    //fmt.Println(args)
}

func sendError(msgPtr *string) {
    log.Fatal(*msgPtr)
    //fmt.Println(*msgPtr)
    //os.Exit(1)
}
