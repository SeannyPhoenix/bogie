package gtfs

import (
	"archive/zip"
	"fmt"
	"slices"
	"strings"
	"time"
)

type GTFSSchedule struct {
	// Required files
	Agencies      map[string]Agency
	Stops         map[string]Stop
	Routes        map[string]Route
	Calendar      map[string]Calendar
	CalendarDates map[string]CalendarDate
	Trips         map[string]Trip
	StopTimes     map[string]StopTime
	Levels        map[string]Level

	unusedFiles []string
	errors      errorList
	warnings    errorList
}

func (s GTFSSchedule) Errors() errorList {
	return s.errors
}

type gtfsSpec[R record] struct {
	set func(*GTFSSchedule, map[string]R)
}

type fileParser interface {
	parseFile(*zip.File, *GTFSSchedule, *errorList)
}

func (spec gtfsSpec[R]) parseFile(f *zip.File, schedule *GTFSSchedule, errors *errorList) {
	r, err := f.Open()
	if err != nil {
		_ = errors.add(fmt.Errorf("error opening file: %w", err))
		return
	}
	defer r.Close()

	records := make(map[string]R)

	parse(r, records, errors)

	spec.set(schedule, records)
}

var gtfsSpecs = map[string]fileParser{
	"agency.txt":         gtfsSpec[Agency]{set: func(s *GTFSSchedule, r map[string]Agency) { s.Agencies = r }},
	"stops.txt":          gtfsSpec[Stop]{set: func(s *GTFSSchedule, r map[string]Stop) { s.Stops = r }},
	"routes.txt":         gtfsSpec[Route]{set: func(s *GTFSSchedule, r map[string]Route) { s.Routes = r }},
	"calendar.txt":       gtfsSpec[Calendar]{set: func(s *GTFSSchedule, r map[string]Calendar) { s.Calendar = r }},
	"calendar_dates.txt": gtfsSpec[CalendarDate]{set: func(s *GTFSSchedule, r map[string]CalendarDate) { s.CalendarDates = r }},
	"trips.txt":          gtfsSpec[Trip]{set: func(s *GTFSSchedule, r map[string]Trip) { s.Trips = r }},
	"stop_times.txt":     gtfsSpec[StopTime]{set: func(s *GTFSSchedule, r map[string]StopTime) { s.StopTimes = r }},
	"levels.txt":         gtfsSpec[Level]{set: func(s *GTFSSchedule, r map[string]Level) { s.Levels = r }},
}

func OpenScheduleFromZipFile(fn string) (GTFSSchedule, error) {
	r, err := zip.OpenReader(fn)
	if err != nil {
		return GTFSSchedule{}, err
	}
	defer r.Close()

	s := parseSchedule(r)

	return s, nil
}

func parseSchedule(r *zip.ReadCloser) GTFSSchedule {
	var s GTFSSchedule

	for _, f := range r.File {
		spec := gtfsSpecs[f.Name]
		if spec == nil {
			s.unusedFiles = append(s.unusedFiles, f.Name)
			_ = s.warnings.add(fmt.Errorf("unused file: %s", f.Name))
			continue
		}
		spec.parseFile(f, &s, &s.errors)
	}

	return s
}

type StopTimeInfo struct {
	Time      string
	Route     string
	Headsign  string
	ServiceID string
}

type StopInfo struct {
	Stop  Stop
	Times []StopTimeInfo
}

func (si StopInfo) String() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Stop: %s\n", si.Stop.Name)
	for _, info := range si.Times {
		fmt.Fprintf(&sb, "  %s Route %s Headsign: %s ServiceID: %s\n", info.Time, info.Route, info.Headsign, info.ServiceID)
	}
	return sb.String()
}

func GetStopInfo(id string, schedule GTFSSchedule) StopInfo {
	var si StopInfo

	stop, ok := schedule.Stops[id]
	if !ok {
		return si
	}
	si.Stop = stop

	for _, st := range schedule.StopTimes {
		if st.StopID == id {
			r := schedule.Routes[schedule.Trips[st.TripID].RouteID]
			h := st.StopHeadsign

			trip := schedule.Trips[st.TripID]

			if h == "" {
				h = trip.Headsign
			}
			si.Times = append(si.Times, StopTimeInfo{
				Time:      st.DepartureTime.Format(time.TimeOnly),
				Route:     r.ShortName,
				Headsign:  h,
				ServiceID: trip.ServiceID,
			})
		}
	}

	slices.SortFunc(si.Times, func(a, b StopTimeInfo) int {
		t := strings.Compare(a.Time, b.Time)
		if t != 0 {
			return t
		}
		return strings.Compare(a.ServiceID, b.ServiceID)
	})

	return si
}
