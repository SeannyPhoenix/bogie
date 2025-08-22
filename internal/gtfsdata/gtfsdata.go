package gtfsdata

import (
	"context"
	"log/slog"

	"github.com/seannyphoenix/bogie/pkg/gtfs"
)

type BogieGtfsData struct {
	AgencyData   agencyData
	RouteData    routeData
	TripData     tripData
	StopData     stopData
	StopTimeData stopTimeData
}

func ParseSchedule(ctx context.Context, sch gtfs.GTFSSchedule, gzip bool) (BogieGtfsData, error) {
	slog.InfoContext(ctx, "Parsing GTFS schedule")

	data := BogieGtfsData{}

	agencies, err := parseAgencies(ctx, sch)
	if err != nil {
		return data, err
	}
	data.AgencyData = agencies

	routes, err := parseRoutes(ctx, sch, data)
	if err != nil {
		return data, err
	}
	data.RouteData = routes

	trips, err := parseTrips(ctx, sch, data)
	if err != nil {
		return data, err
	}
	data.TripData = trips

	stops, err := parseStops(ctx, sch)
	if err != nil {
		return data, err
	}
	data.StopData = stops

	stopTimes, err := parseStopTimes(ctx, sch, data)
	if err != nil {
		return data, err
	}
	data.StopTimeData = stopTimes

	err = writeAllCSVs(ctx, data, gzip)
	if err != nil {
		slog.ErrorContext(ctx, "Error writing GTFS data to CSV files")
		return data, err
	}

	return data, nil
}

func writeAllCSVs(ctx context.Context, data BogieGtfsData, gzip bool) error {
	slog.InfoContext(ctx, "Writing GTFS data to CSV files")

	mm := make(map[string]meta)

	m, err := writeCsv(ctx, writeCsvOptions[Agency]{
		records: sortedKeyedData(data.AgencyData.Agencies),
		path:    "gtfs_out/agency",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["agency"] = m

	m, err = writeCsv(ctx, writeCsvOptions[Route]{
		records: sortedKeyedData(data.RouteData.Routes),
		path:    "gtfs_out/routes",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["routes"] = m

	m, err = writeCsv(ctx, writeCsvOptions[Trip]{
		records: sortedKeyedData(data.TripData.Trips),
		path:    "gtfs_out/trips",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["trips"] = m

	m, err = writeCsv(ctx, writeCsvOptions[Stop]{
		records: sortedKeyedData(data.StopData.Stops),
		path:    "gtfs_out/stops",
		gzip:    gzip,
	})
	if err != nil {
		return err
	}
	mm["stops"] = m

	// m, err = writeCsv(ctx, writeCsvOptions[StopTime]{
	// 	records: sortedKeyedData(data.StopTimeData.StopTimes),
	// 	path:    "gtfs_out/stop_times",
	// })
	// if err != nil {
	// 	return err
	// }
	// mm["stop_times"] = m

	err = writeMeta("gtfs_out", mm)

	return err
}
