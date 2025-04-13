package gtfsdata

import (
	"context"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
	slogctx "github.com/veqryn/slog-context"
)

type Agency struct {
	Id uuid.UUID `json:"id" csv:"agency_id"` // partition key

	Name        string `json:"name" csv:"agency_name"`
	URL         string `json:"url" csv:"agency_url"`
	Timezone    string `json:"timezone" csv:"agency_timezone"`
	Lang        string `json:"lang,omitempty" csv:"agency_lang"`
	Phone       string `json:"phone,omitempty" csv:"agency_phone"`
	FareURL     string `json:"fareUrl,omitempty" csv:"agency_fare_url"`
	AgencyEmail string `json:"email,omitempty" csv:"agency_email"`
}

func (a Agency) partitionKey() []byte {
	return a.Id[:]
}

func (a Agency) sortKey() []byte {
	return nil
}

type agencyData struct {
	Agencies    map[uuid.UUID]Agency `json:"agencies"`
	GtfsToBogie map[string]uuid.UUID `json:"-"`
}

func parseAgencies(ctx context.Context, sch gtfs.GTFSSchedule) (agencyData, error) {
	log := slogctx.FromCtx(ctx)
	log.Info("Parsing data")

	agencies := agencyData{
		Agencies:    make(map[uuid.UUID]Agency),
		GtfsToBogie: make(map[string]uuid.UUID),
	}

	for _, agency := range sch.Agencies {
		id := uuid.New()
		agencies.GtfsToBogie[agency.ID] = id

		agencies.Agencies[id] = Agency{
			Id:          id,
			Name:        agency.Name,
			URL:         agency.URL,
			Timezone:    agency.Timezone,
			Lang:        agency.Lang,
			Phone:       agency.Phone,
			FareURL:     agency.FareURL,
			AgencyEmail: agency.AgencyEmail,
		}
	}

	return agencies, nil
}
