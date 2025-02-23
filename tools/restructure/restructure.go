package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/internal/models"
	"github.com/seannyphoenix/bogie/pkg/csvmum"
	"github.com/seannyphoenix/bogie/pkg/util"
)

type orig struct {
	User            string
	Agency          string
	Route           string
	Run             string
	VehicleID       string
	Orientation     string
	VehicleCount    *int
	VehiclePosition *int
	Date            string
	DepartureStop   string
	DepartureTime   string
	ArrivalStop     string
	ArrivalTime     string
	Notes0          string
	Notes1          string
	Notes2          string
	Notes3          string
	Notes4          string
}

var userSeanny = uuid.MustParse("40eab4cc-fd49-4116-91ae-e7534b236121")
var userAndy = uuid.MustParse("d7d61060-9c40-4017-9bec-c2677d38f63c")

func main() {
	oo, err := getOrig()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ee, _ := split(oo)

	util.PrintAsFormattedJSON(ee)
	// util.PrintAsFormattedJSON(dd)
}

func getOrig() ([]orig, error) {
	var oo []orig

	f, err := os.Open("orig.csv")
	if err != nil {
		return nil, err
	}

	cm, err := csvmum.NewUnmarshaler[orig](f)
	if err != nil {
		return nil, err
	}

	for {
		var o orig
		err = cm.Unmarshal(&o)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		oo = append(oo, o)
	}
	return oo, nil
}

func split(oo []orig) ([]models.Event, []models.Document) {
	var ee = make([]models.Event, len(oo))
	var dd = make([]models.Document, len(oo))

	for i, o := range oo {
		user := getUser(o)
		if user == uuid.Nil {
			log.Printf("User not found for %s in recoerd %d", o.User, i)
			continue
		}

		leg := models.NewDocument("leg", &user)
		dd = append(dd, leg)

		var hasDep bool

		if o.DepartureStop != "" || o.DepartureTime != "" {
			e := models.NewEventDocuemnt(&user)
			e.EventType = "departure"
			if o.DepartureTime != "" {
				time, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", o.Date, o.DepartureTime))
				if err != nil {
					log.Printf("Error parsing time for %s in record %d", o.DepartureTime, i)
					continue
				}
				e.Timestamp = time
				e.Exact = true
			}
			if e.Location != "" {
				e.Location = o.DepartureStop
			}
			e.Agency = o.Agency
			e.Route = o.Route
			e.Vehicle = getVehicle(o)
			e.Tags = append(e.Tags, leg.Id)
			e.Notes = getNotes(o)
			ee = append(ee, e)
		}

		if o.ArrivalStop != "" || o.ArrivalTime != "" {
			e := models.NewEventDocuemnt(&user)
			e.EventType = "arrival"
			if o.ArrivalTime != "" {
				time, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", o.Date, o.ArrivalTime))
				if err != nil {
					log.Printf("Error parsing time for %s in record %d", o.ArrivalTime, i)
					continue
				}
				e.Timestamp = time
				e.Exact = true
			}
			if e.Location != "" {
				e.Location = o.ArrivalStop
			}
			e.Agency = o.Agency
			e.Route = o.Route
			if !hasDep {
				e.Vehicle = getVehicle(o)
			}
			e.Tags = append(e.Tags, leg.Id)
			e.Notes = getNotes(o)
			ee = append(ee, e)
		}
	}

	return ee, dd
}

func getUser(o orig) uuid.UUID {
	o.User = strings.ToLower(o.User)
	if o.User == "seanny" {
		return userSeanny
	} else if o.User == "andy" {
		return userAndy
	}
	return uuid.Nil
}

func getNotes(o orig) []string {
	notes := []string{o.Notes0, o.Notes1, o.Notes2, o.Notes3, o.Notes4}
	var nn []string
	for _, n := range notes {
		if n != "" {
			nn = append(nn, n)
		}
	}
	return nn
}

func getVehicle(o orig) *models.Vehicle {
	if o.VehicleID == "" &&
		o.VehiclePosition == nil &&
		o.VehicleCount == nil &&
		o.Orientation == "" {
		return nil
	}

	v := models.Vehicle{
		ID: o.VehicleID,
	}

	if o.VehiclePosition != nil && o.VehicleID != "" {
		v.Sequence[o.VehicleID] = *o.VehiclePosition
	} else {
		v.Sequence[o.VehicleID] = 0
	}

	if o.VehicleCount != nil {
		v.Length = *o.VehicleCount
	}

	return &v
}
