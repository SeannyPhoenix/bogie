package gtfsdata

import (
	"context"

	"github.com/seannyphoenix/bogie/pkg/gtfs"
	slogctx "github.com/veqryn/slog-context"
)

type BogieGtfsData struct {
	AgencyData   agencyData
	RouteData    RouteData
	TripData     TripData
	StopData     StopData
	StopTimeData StopTimeData
}

func ParseSchedule(sch gtfs.GTFSSchedule, gzip bool) (BogieGtfsData, error) {
	ctx := newSlogCtx(newSlogCtxOptions{})

	log := slogctx.FromCtx(ctx)
	log.Info("Parsing GTFS schedule")

	data := BogieGtfsData{}

	agencies, err := parseAgencies(slogctx.With(ctx, "data", "agencies"), sch)
	if err != nil {
		return data, err
	}
	data.AgencyData = agencies

	routes, err := parseRoutes(slogctx.With(ctx, "data", "routes"), sch, data)
	if err != nil {
		return data, err
	}
	data.RouteData = routes

	trips, err := parseTrips(slogctx.With(ctx, "data", "trips"), sch, data)
	if err != nil {
		return data, err
	}
	data.TripData = trips

	stops, err := parseStops(slogctx.With(ctx, "data", "stops"), sch)
	if err != nil {
		return data, err
	}
	data.StopData = stops

	// stopTimes, err := parseStopTimes(slogctx.With(ctx, "data", "stopTimes"), sch, data)
	// if err != nil {
	// 	return data, err
	// }
	// data.StopTimeData = stopTimes

	err = writeAllCSVs(ctx, data, gzip)
	if err != nil {
		log.Error("Error writing GTFS data to CSV files")
		return data, err
	}

	return data, nil
}

func writeAllCSVs(ctx context.Context, data BogieGtfsData, gzip bool) error {
	log := slogctx.FromCtx(ctx)
	log.Info("Writing GTFS data to CSV files")

	mm := make(map[string]meta)

	m, err := writeCsv(slogctx.With(ctx, "data", "agencies"), writeCsvOptions[Agency]{
		records: sortedKeyedData(data.AgencyData.Agencies),
		path:    "gtfs_out/agency",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["agency"] = m

	m, err = writeCsv(slogctx.With(ctx, "data", "routes"), writeCsvOptions[Route]{
		records: sortedKeyedData(data.RouteData.Routes),
		path:    "gtfs_out/routes",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["routes"] = m

	m, err = writeCsv(slogctx.With(ctx, "data", "trips"), writeCsvOptions[Trip]{
		records: sortedKeyedData(data.TripData.Trips),
		path:    "gtfs_out/trips",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["trips"] = m

	m, err = writeCsv(slogctx.With(ctx, "data", "stops"), writeCsvOptions[Stop]{
		records: sortedKeyedData(data.StopData.Stops),
		path:    "gtfs_out/stops",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["stops"] = m

	// m, err = writeCsv(slogctx.With(ctx, "data", "stopTimes"), writeCsvOptions[StopTime]{
	// 	records: sortedKeyedData(data.StopTimeData.StopTimes),
	// 	path:    "gtfs_out/stop_times",
	// })
	// if err != nil {
	// 	return err
	// }
	// mm["stop_times"] = m

	err = writeMeta("gtfs_out", mm)

	return nil
}
