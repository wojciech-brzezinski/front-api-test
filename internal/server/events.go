package server

import (
	"encoding/json"
	"time"
)

type EventFilter struct {
	LocationType string    `json:"locationType"`
	TimeStart    time.Time `json:"timeStart"`
	TimeEnd      time.Time `json:"timeEnd"`
}

type EventParticipant struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

type Event struct {
	Name         string             `json:"name"`
	TimeStart    time.Time          `json:"timeStart"`
	TimeEnd      time.Time          `json:"timeEnd"`
	Location     string             `json:"location"`
	LocationType string             `json:"locationType"`
	Participants []EventParticipant `json:"participants"`
}

var _EVENTS []Event

func events(filter *EventFilter) []Event {
	items := make([]Event, len(_EVENTS))
	index := 0

	for _, event := range _EVENTS {
		if nil != filter {
			if "" != filter.LocationType && event.LocationType != filter.LocationType {
				continue
			}

			if !filter.TimeStart.IsZero() && event.TimeStart.Before(filter.TimeStart) {
				continue
			}

			if !filter.TimeEnd.IsZero() && event.TimeEnd.After(filter.TimeEnd) {
				continue
			}
		}

		items[index] = event
		index++
	}

	return items
}

func init() {
	_ = json.Unmarshal([]byte(jsonEvents), &_EVENTS)
}
