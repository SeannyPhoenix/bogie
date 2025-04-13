package gtfsdata

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
	slogctx "github.com/veqryn/slog-context"
)

type Route struct {
	AgencyId uuid.UUID `json:"agencyId" csv:"agency_id"` // partition key
	Id       uuid.UUID `json:"id" csv:"route_id"`        // sort key

	ShortName         string `json:"shortName" csv:"route_short_name"`
	LongName          string `json:"longName" csv:"route_long_name"`
	Type              string `json:"type" csv:"route_type"`
	Desc              string `json:"desc,omitempty" csv:"route_desc"`
	URL               string `json:"url,omitempty" csv:"route_url"`
	Color             string `json:"color,omitempty" csv:"route_color"`
	TextColor         string `json:"textColor,omitempty" csv:"route_text_color"`
	SortOrder         string `json:"sortOrder,omitempty" csv:"route_sort_order"`
	ContinuousPickup  string `json:"continuousPickup,omitempty" csv:"route_continuous_pickup"`
	ContinuousDropOff string `json:"continuousDropOff,omitempty" csv:"route_continuous_drop_off"`
	NetworkID         string `json:"networkId,omitempty" csv:"route_network_id"`
}

func (r Route) partitionKey() []byte {
	return r.AgencyId[:]
}

func (r Route) sortKey() []byte {
	return r.Id[:]
}

type RouteData struct {
	Routes      map[uuid.UUID]Route  `json:"routes"`
	GtfsToBogie map[string]uuid.UUID `json:"-"`
}

func parseRoutes(ctx context.Context, sch gtfs.GTFSSchedule, data BogieGtfsData) (RouteData, error) {
	log := slogctx.FromCtx(ctx)
	log.Info("Parsing data")

	routes := RouteData{
		Routes:      make(map[uuid.UUID]Route),
		GtfsToBogie: make(map[string]uuid.UUID),
	}

	for _, gtfsRoute := range sch.Routes {
		id := uuid.New()
		var agencyId uuid.UUID

		if aid, ok := data.AgencyData.GtfsToBogie[gtfsRoute.AgencyID]; !ok {
			if len(data.AgencyData.GtfsToBogie) == 1 {
				for _, id := range data.AgencyData.GtfsToBogie {
					agencyId = id
					break
				}
			} else {
				return routes, fmt.Errorf("agency ID %s not found", gtfsRoute.AgencyID)
			}
		} else {
			agencyId = aid
		}

		routes.GtfsToBogie[gtfsRoute.ID] = id

		routes.Routes[id] = Route{
			Id:                id,
			AgencyId:          agencyId,
			ShortName:         gtfsRoute.ShortName,
			LongName:          gtfsRoute.LongName,
			Type:              gtfsRoute.Type,
			Desc:              gtfsRoute.Desc,
			URL:               gtfsRoute.URL,
			Color:             gtfsRoute.Color,
			TextColor:         gtfsRoute.TextColor,
			SortOrder:         gtfsRoute.SortOrder,
			ContinuousPickup:  gtfsRoute.ContinuousPickup,
			ContinuousDropOff: gtfsRoute.ContinuousDropOff,
			NetworkID:         gtfsRoute.NetworkID,
		}
	}

	return routes, nil
}
