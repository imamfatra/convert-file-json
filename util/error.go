package util

import "log"

func IfErrorPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
