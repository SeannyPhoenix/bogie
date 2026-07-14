package gtfs

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func Overview(c map[string]GTFSSchedule) string {
	var o strings.Builder

	for sid, s := range c {
		fmt.Fprintf(&o, "Schedule %s\n", sid[0:4])
		fmt.Fprintf(&o, "  %d agencies\n", len(s.Agencies))
		fmt.Fprintf(&o, "  %d stops\n", len(s.Stops))
		fmt.Fprintf(&o, "  %d routes\n", len(s.Routes))
		fmt.Fprintf(&o, "  %d calendar entries\n", len(s.Calendar))
		fmt.Fprintf(&o, "  %d calendar dates\n", len(s.CalendarDates))
		fmt.Fprintf(&o, "  %d trips\n", len(s.Trips))
		fmt.Fprintf(&o, "  %d stop times\n", len(s.StopTimes))
		fmt.Fprintf(&o, "  %d levels\n", len(s.Levels))
		fmt.Fprintf(&o, "  %d errors\n", len(s.errors))
		fmt.Fprint(&o, "\n")
	}

	return o.String()
}

func CreateGTFSCollection(zipFiles []string) (map[string]GTFSSchedule, error) {
	sc := make(map[string]GTFSSchedule)

	for _, path := range zipFiles {
		s, err := OpenScheduleFromZipFile(path)
		if err != nil {
			return sc, err
		}

		sc[uuid.NewString()] = s
	}

	return sc, nil
}
