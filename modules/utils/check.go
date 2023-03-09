package utils

import (
	"errors"
	"log"
	"os"
)

func Err(err error) bool {
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func FileExists(name string) (bool, error) {
	_, err := os.Stat(name)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
