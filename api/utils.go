package api

import (
	"log"
	"time"
)

// Useful functions for the application

// ExpireIn returns unix timestamp given a time duration
func ExpireIn(tm time.Duration) int64 {
	return EpochNow() + int64(tm.Seconds())
}

// EpochNow returns a unix timestamp
func EpochNow() int64 {
	return time.Now().UTC().Unix()
}

// CheckError checks the error if it is not nil logs the error
func CheckError(err error) {
	if err != nil {
		log.Println(err)
	}
}
