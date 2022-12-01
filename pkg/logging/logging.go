package logging

import (
    "log"
)

func SendError(msgPtr *string) {
    log.Fatal(*msgPtr)
}
