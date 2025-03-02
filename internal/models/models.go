package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/gtfs"
)

type BaseDocument struct {
	Id        uuid.UUID `json:"id"`
	Type      uuid.UUID `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	User  *uuid.UUID  `json:"user,omitempty"`
	Tags  []uuid.UUID `json:"tags,omitempty"`
	Notes []string    `json:"notes,omitempty"`
}

func InitDoc(doc BaseDocument) BaseDocument {
	var now = time.Now()

	if doc.Id == uuid.Nil {
		doc.Id = uuid.New()
	}
	if doc.CreatedAt.IsZero() {
		doc.CreatedAt = now
	}
	if doc.UpdatedAt.IsZero() {
		doc.UpdatedAt = now
	}
	if doc.Status == "" {
		doc.Status = DocStatusActive
	}

	return doc
}

type User struct {
	BaseDocument

	Username string `json:"username"`
}

type Agency struct {
	BaseDocument

	GTFS gtfs.Agency `json:"gtfs"`
}

type Route struct {
	BaseDocument

	Agency uuid.UUID  `json:"agency"`
	GTFS   gtfs.Route `json:"gtfs"`
}

type Stop struct {
	BaseDocument

	Agency uuid.UUID `json:"agency"`
	GTFS   gtfs.Stop `json:"gtfs"`
}

type Event struct {
	BaseDocument

	EventType   string    `json:"eventType"`
	Timestamp   time.Time `json:"timestamp"`
	Granularity int       `json:"granularity"`

	Location *uuid.UUID `json:"location,omitempty"`
	Route    *uuid.UUID `json:"route,omitempty"`
	Vehicle  *uuid.UUID `json:"vehicle,omitempty"`
	Run      string     `json:"run,omitempty"`
}

type Vehicle struct {
	BaseDocument

	Unit     VehicleUnit   `json:"unit,omitempty"`
	Sequence []VehicleUnit `json:"sequence,omitempty"`
}

type VehicleUnit struct {
	Id          uuid.UUID `json:"id"`
	Orientation string    `json:"orientation,omitempty"`
}

type Unit struct {
	BaseDocument

	UnitID string    `json:"unitID"`
	Agency uuid.UUID `json:"agency"`
}

const (
	GranularityNone        = 0
	GranularitySecond      = 1
	GranularityMinute      = 2
	GranularityFiveMinutes = 3
)

const (
	DocTypeUser     = "user"
	DocTypeAgency   = "agency"
	DocTypeRoute    = "route"
	DocTypeStop     = "stop"
	DocTypeEvent    = "event"
	DocTypeLocation = "location"
	DocTypeVehicle  = "vehicle"
	DocTypeUnit     = "unit"
	DocTypeJourney  = "journey"
	DocTypeLeg      = "leg"

	DocStatusActive   = "active"
	DocStatusInactive = "inactive"

	EventTypeArrival   = "arrival"
	EventTypeDeparture = "departure"

	LocationTypeStop = "stop"
)

var NameMap = map[string]uuid.UUID{
	DocTypeUser:     uuid.MustParse("8d930d82-24e9-4fc1-824b-9e1253d4ee02"),
	DocTypeAgency:   uuid.MustParse("c5ebd7b0-5f83-4136-ac3a-7f1a46b7c084"),
	DocTypeRoute:    uuid.MustParse("2eccfc6f-9a31-4f0c-b3e6-0960cf83a3af"),
	DocTypeStop:     uuid.MustParse("15aefc34-9fa5-4a05-b616-3e5183108d33"),
	DocTypeEvent:    uuid.MustParse("88c2333e-2bc2-4063-b865-719c24211d2c"),
	DocTypeLocation: uuid.MustParse("958ade3b-0e67-4426-98b8-9902e40b8bd8"),
	DocTypeVehicle:  uuid.MustParse("128f17e9-8cde-458e-acd6-1a39926e0283"),
	DocTypeUnit:     uuid.MustParse("915ddb34-93ba-4e2f-99f2-ea814bb2790d"),
	DocTypeJourney:  uuid.MustParse("e68eac59-63a6-472d-b661-384e8453586b"),
	DocTypeLeg:      uuid.MustParse("3e3d7b12-b881-43fb-82db-33afc4d61d68"),

	DocStatusActive:   uuid.MustParse("5c96e882-112a-42e3-adf3-941f28ff9956"),
	DocStatusInactive: uuid.MustParse("aca32352-08c1-40ee-8a6e-e95d77e68724"),

	EventTypeArrival:   uuid.MustParse("5fe59547-84ea-4673-8665-dca2967c818e"),
	EventTypeDeparture: uuid.MustParse("a6de1146-8013-49aa-987a-31d3f935de4c"),
}

var IDMap map[uuid.UUID]string

func init() {
	IDMap = make(map[uuid.UUID]string, len(NameMap))
	for k, v := range NameMap {
		IDMap[v] = k
	}
}
