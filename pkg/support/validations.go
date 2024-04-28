package support

import (
	"fmt"
	"log"
)

func PanicOnError(err error, format string, args ...interface{}) {
	if err != nil {
		msg := fmt.Sprintf(format, args...)
		log.Fatalf("%s: %v", msg, err)
	}
}
