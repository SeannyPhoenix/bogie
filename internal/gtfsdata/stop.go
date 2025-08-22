package gtfsdata

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
)

type Stop struct {
	Id uuid.UUID `json:"id" csv:"stop_id"` // partition key

	Code               string   `json:"code,omitempty" csv:"stop_code"`
	Name               string   `json:"name,omitempty" csv:"stop_name"`
	TtsName            string   `json:"ttsName,omitempty" csv:"tts_stop_name"`
	Desc               string   `json:"desc,omitempty" csv:"stop_desc"`
	Lat                *float64 `json:"lat,omitempty" csv:"stop_lat"`
	Lon                *float64 `json:"lon,omitempty" csv:"stop_lon"`
	ZoneId             string   `json:"zoneId,omitempty" csv:"zone_id"`
	Url                string   `json:"url,omitempty" csv:"stop_url"`
	LocationType       *int     `json:"locationType,omitempty" csv:"location_type"`
	Timezone           string   `json:"timezone,omitempty" csv:"stop_timezone"`
	WheelchairBoarding *int     `json:"wheelchairBoarding,omitempty" csv:"wheelchair_boarding"`
	PlatformCode       string   `json:"platformCode,omitempty" csv:"platform_code"`
	// ParentStation      uuid.UUID `json:"parentStation" csv:"parent_station"`
	// LevelId            uuid.UUID `json:"levelId,omitempty" csv:"level_id"`
}

func (s Stop) partitionKey() []byte {
	return s.Id[:]
}

func (s Stop) sortKey() []byte {
	return nil
}

type stopData struct {
	Stops       map[uuid.UUID]Stop   `json:"stops"`
	GtfsToBogie map[string]uuid.UUID `json:"-"`
}

func parseStops(ctx context.Context, sch gtfs.GTFSSchedule) (stopData, error) {
	slog.InfoContext(ctx, "Parsing stops")

	stopData := stopData{
		Stops:       make(map[uuid.UUID]Stop),
		GtfsToBogie: make(map[string]uuid.UUID),
	}

	for _, stop := range sch.Stops {
		stopId := uuid.New()
		stopData.GtfsToBogie[stop.ID] = stopId

		stopData.Stops[stopId] = Stop{
			Id:                 stopId,
			Code:               stop.Code,
			Name:               stop.Name,
			TtsName:            stop.TTSName,
			Desc:               stop.Desc,
			Lat:                stop.Latitude,
			Lon:                stop.Longitude,
			ZoneId:             stop.ZoneID,
			Url:                stop.URL,
			LocationType:       stop.LocationType,
			Timezone:           stop.Timezone,
			WheelchairBoarding: stop.WheelchairBoarding,
			PlatformCode:       stop.PlatformCode,
			// ParentStation:
			// LevelId:
		}
	}

	return stopData, nil
}
