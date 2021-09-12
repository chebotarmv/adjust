package events

import (
	"adjust/helpers"
	"errors"
)

func GetEvents(path string) ([][]string, error) {
	events, err := helpers.ReadCsv(path)
	if err != nil {
		return nil, errors.New("can't get info from events.csv file")
	}
	return events, nil

}
