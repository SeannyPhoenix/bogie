package gtfsdata

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
)

type Trip struct {
	RouteId uuid.UUID `json:"routeId" csv:"route_id"` // partition key
	Id      uuid.UUID `json:"id" csv:"trip_id"`       // sort key

	ServiceId            string `json:"serviceId" csv:"service_id"`
	Headsign             string `json:"headSign,omitempty" csv:"trip_headsign"`
	ShortName            string `json:"shortName,omitempty" csv:"trip_short_name"`
	DirectionId          int    `json:"directionId,omitempty" csv:"direction_id"`
	BlockId              string `json:"blockId,omitempty" csv:"block_id"`
	ShapeId              string `json:"shapeId,omitempty" csv:"shape_id"`
	WheelchairAccessible int    `json:"wheelchairAccessible,omitempty" csv:"wheelchair_accessible"`
	BikesAllowed         int    `json:"bikesAllowed,omitempty" csv:"bikes_allowed"`
}

func (t Trip) partitionKey() []byte {
	return t.RouteId[:]
}

func (t Trip) sortKey() []byte {
	return t.Id[:]
}

type tripData struct {
	Trips       map[uuid.UUID]Trip      `json:"trips"`
	GtfsToBogie map[string]uuid.UUID    `json:"-"`
	TripRoutes  map[uuid.UUID]uuid.UUID `json:"-"`
}

func parseTrips(ctx context.Context, sch gtfs.GTFSSchedule, data BogieGtfsData) (tripData, error) {
	slog.InfoContext(ctx, "Parsing trips")

	trips := tripData{
		Trips:       make(map[uuid.UUID]Trip),
		GtfsToBogie: make(map[string]uuid.UUID),
		TripRoutes:  make(map[uuid.UUID]uuid.UUID),
	}

	for _, gtfsTrip := range sch.Trips {
		id := uuid.New()

		routeId := data.RouteData.GtfsToBogie[gtfsTrip.RouteID]

		trips.GtfsToBogie[gtfsTrip.ID] = id
		trips.TripRoutes[id] = routeId

		trips.Trips[id] = Trip{
			Id:                   id,
			RouteId:              routeId,
			ServiceId:            gtfsTrip.ServiceID,
			Headsign:             gtfsTrip.Headsign,
			ShortName:            gtfsTrip.ShortName,
			DirectionId:          gtfsTrip.DirectionID,
			BlockId:              gtfsTrip.BlockID,
			ShapeId:              gtfsTrip.ShapeID,
			WheelchairAccessible: gtfsTrip.WheelchairAccessible,
			BikesAllowed:         gtfsTrip.BikesAllowed,
		}
	}

	return trips, nil
}
