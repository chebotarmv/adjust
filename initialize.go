package main

import (
	"adjust/events"
	"errors"
	"fmt"
	"os"
)

func checkFiles(path string) error {
	_, err := os.Stat(path + "actors.csv")
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("file actors.csv required")
	}
	_, err = os.Stat(path + "events.csv")
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("file events.csv required")
	}
	_, err = os.Stat(path + "repos.csv")
	if errors.Is(err, os.ErrNotExist) {
		return errors.New("file repos.csv required")
	}
	return nil
}

func getEvents(path string) ([][]string, error) {
	ev, err := events.GetEvents(fmt.Sprintf("%s%s", path, "events.csv"))
	if err != nil {
		return nil, err
	}
	return ev, nil
}
