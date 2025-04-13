package gtfsdata

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
	slogctx "github.com/veqryn/slog-context"
)

type StopTime struct {
	RouteId uuid.UUID `json:"routeId" csv:"route_id"` // partition key
	Id      string    `json:"id" csv:"stop_time_id"`  // sort key

	TripId                   uuid.UUID `json:"tripId" csv:"trip_id"`
	StopSequence             int       `json:"stopSequence" csv:"stop_sequence"`
	ArrivalTime              gtfs.Time `json:"arrivalTime" csv:"arrival_time"`
	DepartureTime            gtfs.Time `json:"departureTime" csv:"departure_time"`
	StopId                   uuid.UUID `json:"stopId" csv:"stop_id"`
	StopHeadsign             string    `json:"stopHeadsign,omitempty" csv:"stop_headsign"`
	StartPickupDropOffWindow gtfs.Time `json:"startPickupDropOffWindow,omitempty" csv:"start_pickup_drop_off_window"`
	EndPickupDropOffWindow   gtfs.Time `json:"endPickupDropOffWindow,omitempty" csv:"end_pickup_drop_off_window"`
	PickupType               *int      `json:"pickupType,omitempty" csv:"pickup_type"`
	DropOffType              *int      `json:"dropOffType,omitempty" csv:"drop_off_type"`
	ContinuousPickup         *int      `json:"continuousPickup,omitempty" csv:"continuous_pickup"`
	ContinuousDropOff        *int      `json:"continuousDropOff,omitempty" csv:"continuous_drop_off"`
	ShapeDistTraveled        *float64  `json:"shapeDistTraveled,omitempty" csv:"shape_dist_traveled"`
	Timepoint                *int      `json:"timepoint,omitempty" csv:"timepoint"`

	// LocationGroupId          uuid.UUID `json:"locationGroupId,omitempty" csv:"location_group_id"`
	// LocationId               uuid.UUID `json:"locationId,omitempty" csv:"location_id"`
	// PickupBookingRuleId      uuid.UUID `json:"pickupBookingRuleId" csv:"pickup_booking_rule_id"`
	// DropOffBookingRuleId     uuid.UUID `json:"dropOffBookingRuleId" csv:"drop_off_booking_rule_id"`
}

func (st StopTime) partitionKey() []byte {
	return st.RouteId[:]
}

func (st StopTime) sortKey() []byte {
	return []byte(st.Id)
}

type StopTimeData struct {
	StopTimes   map[string]StopTime `json:"stopTimes"`
	GtfsToBogie map[string]string   `json:"-"`
}

func parseStopTimes(ctx context.Context, sch gtfs.GTFSSchedule, data BogieGtfsData) (StopTimeData, error) {
	log := slogctx.FromCtx(ctx)
	log.Info("Parsing data")

	stopTimes := StopTimeData{
		StopTimes:   make(map[string]StopTime),
		GtfsToBogie: make(map[string]string),
	}

	for _, gtfsStopTime := range sch.StopTimes {
		tripId := data.TripData.GtfsToBogie[gtfsStopTime.TripID]
		routeId := data.TripData.TripRoutes[tripId]
		stopTimeId := getStopTimeId(tripId, gtfsStopTime.StopSequence)

		stopTimes.GtfsToBogie[getStopTimeId(gtfsStopTime.TripID, gtfsStopTime.StopSequence)] = stopTimeId

		stopId := data.StopData.GtfsToBogie[gtfsStopTime.StopID]

		stopTimes.StopTimes[stopTimeId] = StopTime{
			Id:                       stopTimeId,
			TripId:                   tripId,
			RouteId:                  routeId,
			StopSequence:             gtfsStopTime.StopSequence,
			ArrivalTime:              gtfsStopTime.ArrivalTime,
			DepartureTime:            gtfsStopTime.DepartureTime,
			StopId:                   stopId,
			StopHeadsign:             gtfsStopTime.StopHeadsign,
			StartPickupDropOffWindow: gtfsStopTime.StartPickupDropOffWindow,
			EndPickupDropOffWindow:   gtfsStopTime.EndPickupDropOffWindow,
			PickupType:               gtfsStopTime.PickupType,
			DropOffType:              gtfsStopTime.DropOffType,
			ContinuousPickup:         gtfsStopTime.ContinuousPickup,
			ContinuousDropOff:        gtfsStopTime.ContinuousDropOff,
			ShapeDistTraveled:        gtfsStopTime.ShapeDistTraveled,
			Timepoint:                gtfsStopTime.Timepoint,

			// LocationGroupId:
			// LocationId:
			// PickupBookingRuleId:
			// DropOffBookingRuleId:
		}
	}

	return stopTimes, nil
}

func getStopTimeId[I uuid.UUID | string](tripId I, stopSequence int) string {
	return fmt.Sprintf("%s:%04d", tripId, stopSequence)
}
