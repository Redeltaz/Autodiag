package main

import (
    "fmt"
    "log"
    "os"
    "regexp"
)

var validFullArgs = [2]string{"target", "key"}
var validShortArgs = map[string]string{"t": "target", "k": "key"}

type argument struct {
    name string
    value string
}

func main() {
    cliArguments := os.Args
    var arguments []argument

    parseArgs(&cliArguments, &arguments)

    fmt.Println(arguments)
}

func parseArgs(cliArgPtr *[]string, args *[]argument) {
    cliArgs := *cliArgPtr

    regexKeyPattern := regexp.MustCompile(`[-]+`)

    //remove useless first argument
    cliArgs = append(cliArgs[:0], cliArgs[0+1:]...)

    for index, value := range(cliArgs) {
        parsedKey := regexKeyPattern.FindStringSubmatch(value)

        //it's an argument key
        if len(parsedKey) > 0 {
            var argKey string
            isArgKeyValid := false

            //check how many dash there is before the argument key
            if len(parsedKey[0]) == 1 {
                argKey = regexKeyPattern.Split(value, -1)[1]

                //check and assign short argument key to it long version
                for k, v := range(validShortArgs) {
                    if k == argKey {
                        isArgKeyValid = true
                        argKey = v
                    }
                }
            } else if len(parsedKey[0]) == 2 {
                argKey = regexKeyPattern.Split(value, -1)[1]

                for _, v := range(validFullArgs) {
                    if v == argKey {
                        isArgKeyValid = true
                    }
                }
            }

            if !isArgKeyValid {
                errorMsg := "Unknow argument " + value
                sendError(&errorMsg)
            }

            argValue := cliArgs[index + 1]

            newArg := argument{name: argKey, value: argValue}
            *args = append(*args, newArg)
        }
    }
}

func sendError(msgPtr *string) {
    log.Fatal(*msgPtr)
}
