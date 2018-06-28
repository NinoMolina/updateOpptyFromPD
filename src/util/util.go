package util

import (
	"log"
	"strings"
)

func CheckErr(err error, msg string) {
	if err != nil {
		if IsEmpty(msg) {
			log.Println("ERROR - " + err.Error())
		} else {
			log.Println("ERROR - " + msg)
		}
		panic(err)
	}
}

func IsError(err error) bool {
	if err != nil {
		log.Println("ERROR - " + err.Error())
	}
	return (err != nil)
}

func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

