package main

import (
	"time"

	"github.com/seannyphoenix/binarytime/pkg/binarytime"
)

type EventV2CSV struct {
	Begin            string `csv:"begin"`
	Cost             string `csv:"cost"`
	EventType        string `csv:"event_type"`
	Timestamp        string `csv:"timestamp"`
	Date             string `csv:"date"`
	Time             string `csv:"time"`
	OutOfOrder       string `csv:"out_of_order"`
	Location         string `csv:"location"`
	Platform         string `csv:"platform"`
	StopID           string `csv:"stop_id"`
	Agency           string `csv:"agency"`
	Route            string `csv:"route"`
	Run              string `csv:"run"`
	VehicleID        string `csv:"vehicle_id"`
	Count            string `csv:"count"`
	Unit1ID          string `csv:"unit_1_id"`
	Unit1Orientation string `csv:"unit_1_orientation"`
	Unit2ID          string `csv:"unit_2_id"`
	Unit2Orientation string `csv:"unit_2_orientation"`
	Unit3ID          string `csv:"unit_3_id"`
	Unit3Orientation string `csv:"unit_3_orientation"`
	Unit4ID          string `csv:"unit_4_id"`
	Unit4Orientation string `csv:"unit_4_orientation"`
	Unit5ID          string `csv:"unit_5_id"`
	Unit5Orientation string `csv:"unit_5_orientation"`
	Unit6ID          string `csv:"unit_6_id"`
	Unit6Orientation string `csv:"unit_6_orientation"`
	Unit7ID          string `csv:"unit_7_id"`
	Unit7Orientation string `csv:"unit_7_orientation"`
	Unit8ID          string `csv:"unit_8_id"`
	Unit8Orientation string `csv:"unit_8_orientation"`
	Unit9ID          string `csv:"unit_9_id"`
	Unit9Orientation string `csv:"unit_9_orientation"`
	Operator         string `csv:"operator"`
	Note0            string `csv:"note_0"`
	Note1            string `csv:"note_1"`
	Note2            string `csv:"note_2"`
	Note3            string `csv:"note_3"`
	Note4            string `csv:"note_4"`

	binarytime binarytime.Date
}

func (e *EventV2CSV) fillDate(last EventV2CSV) {
	if e.Date == "" && e.Time != "" && last.Date != "" {
		currentTime, _ := time.Parse("15:04", e.Time)
		lastTime, _ := time.Parse("15:04", last.Time)

		if !currentTime.IsZero() && !lastTime.IsZero() && !currentTime.Before(lastTime) {
			e.Date = last.Date
		}
	}
}

func (e *EventV2CSV) setTimestamp() {
	dateField, _ := time.Parse("01/02/2006", e.Date+"/2025")
	timeField, _ := time.Parse("15:04", e.Time)
	d := time.Duration(timeField.Hour()*int(time.Hour) + timeField.Minute()*int(time.Minute))
	t := dateField.Add(d)
	e.binarytime = binarytime.DateFromTime(t)
	ts, err := e.binarytime.Fixed128().StringWithPrecision(7, 11)
	if err != nil {
		panic(err)
	}
	e.Timestamp = ts
}

func (e *EventV2CSV) checkOrder(last EventV2CSV) {
	if !last.binarytime.IsZero() && !e.binarytime.IsZero() {
		cbi := e.binarytime.BigInt()
		lbi := last.binarytime.BigInt()
		if (&cbi).Cmp(&lbi) < 0 {
			e.OutOfOrder = "true"
		}
	}
}
