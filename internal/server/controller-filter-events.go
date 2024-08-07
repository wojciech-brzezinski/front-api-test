package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Filter struct {
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

func filterEvents() gin.HandlerFunc {
	var events []Event

	err := json.Unmarshal([]byte(jsonEvents), &events)

	if nil != err {
		fmt.Println(err)
	}

	for _, event := range events {
		fmt.Println(event.TimeStart)
	}

	return func(context *gin.Context) {
		filter := &Filter{}

		if err := context.BindJSON(filter); nil != err {
			return
		}

		items := make([]Event, len(events))
		index := 0

		for _, event := range events {
			if "" != filter.LocationType && event.LocationType != filter.LocationType {
				continue
			}

			if !filter.TimeStart.IsZero() && event.TimeStart.Before(filter.TimeStart) {
				continue
			}

			if !filter.TimeEnd.IsZero() && event.TimeEnd.After(filter.TimeEnd) {
				continue
			}

			items[index] = event
			index++
		}

		context.JSON(http.StatusOK, items[0:index])
	}
}
