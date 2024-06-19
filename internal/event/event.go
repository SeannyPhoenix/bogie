package event

import (
	"errors"
	"strconv"
	"time"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Event struct {
	Id            uuid.UUID  `json:"id"`
	Type          string     `json:"type"`
	Carrier       string     `json:"carrier,omitempty"`
	Line          string     `json:"line,omitempty"`
	UnitID        string     `json:"unitID,omitempty"`
	UnitCount		 int        `json:"unitCount,omitempty"`
	UnitPosition	 int        `json:"unitPosition,omitempty"`
	DepartureStop string     `json:"departureStop,omitempty"`
	ArrivalStop   string     `json:"arrivalStop,omitempty"`
	DepartureTime *time.Time `json:"departureTime,omitempty"`
	ArrivalTime   *time.Time `json:"arrivalTime,omitempty"`
	Notes         []string   `json:"notes,omitempty"`
}


func GetExampleEvent(id uuid.UUID) Event {
	if id == uuid.Nil {
		id = uuid.New()
	}

	extime := time.Now().Truncate(time.Second)

	return Event{
		Id:            id,
		Type:          documentType.Event,
		Carrier:       "BART",
		Line:          "Red",
		DepartureStop: "Richmond",
		ArrivalStop:   "Millbrae",
		DepartureTime: &extime,
		ArrivalTime:   &extime,
	}
}

func GetExampleEventArray(count int) []Event {
	evs := make([]Event, count)
	for i := 0; i < count; i++ {
		evs[i] = GetExampleEvent(uuid.Nil)
	}
	return evs
}

var ErrBadEventID = errors.New("bad event ID")
var ErrBadDocumentType = errors.New("bad event type")

func (e Event) MarshalDynamoDB() (map[string]dynamodb.AttributeValue, error) {
	if e.Id == uuid.Nil {
		return nil, ErrBadEventID
	}

	data := map[string]dynamodb.AttributeValue{
		"id": &dynamodb.AttributeValueMemberB{Value: e.Id[:]},
	}

	if id, ok := documentType.NameMap[e.Type]; ok {
		data["t"] = &dynamodb.AttributeValueMemberB{Value: id[:]}
	} else {
		return nil, ErrBadDocumentType
	}

	if e.Carrier != "" {
		data["c"] = &dynamodb.AttributeValueMemberS{Value: e.Carrier}
	}

	if e.Line != "" {
		data["l"] = &dynamodb.AttributeValueMemberS{Value: e.Line}
	}

	if e.UnitID != "" {
		data["u"] = &dynamodb.AttributeValueMemberS{Value: e.UnitID}
	}

	if e.DepartureStop != "" {
		data["ds"] = &dynamodb.AttributeValueMemberS{Value: e.DepartureStop}
	}

	if e.ArrivalStop != "" {
		data["as"] = &dynamodb.AttributeValueMemberS{Value: e.ArrivalStop}
	}

	if e.DepartureTime != nil && !e.DepartureTime.IsZero() {
		data["dt"] = &dynamodb.AttributeValueMemberN{Value: strconv.FormatInt(e.DepartureTime.Unix(), 10)}
	}

	if e.ArrivalTime != nil && !e.ArrivalTime.IsZero() {
		data["at"] = &dynamodb.AttributeValueMemberN{Value: strconv.FormatInt(e.ArrivalTime.Unix(), 10)}
	}

	if len(e.Notes) > 0 {
		data["n"] = &dynamodb.AttributeValueMemberSS{Value: e.Notes}
	}

	return data, nil
}

func (e *Event) UnmarshalDynamoDB(data map[string]dynamodb.AttributeValue) error {
	if id := getUUID(data["id"]); id != uuid.Nil {
		e.Id = id
	} else {
		return ErrBadEventID
	}

	e.Type = documentType.IDMap[getUUID(data["t"])]
	e.Carrier = getString(data["c"])
	e.Line = getString(data["l"])
	e.UnitID = getString(data["u"])
	e.DepartureStop = getString(data["ds"])
	e.ArrivalStop = getString(data["as"])
	e.Notes = getStringSlice(data["n"])

	if t := getTime(data["dt"]); !t.IsZero() {
		e.DepartureTime = &t
	}

	if t := getTime(data["at"]); !t.IsZero() {
		e.ArrivalTime = &t
	}

	return nil
}

func getUUID(data dynamodb.AttributeValue) uuid.UUID {
	if data == nil {
		return uuid.Nil
	}

	if id, ok := data.(*dynamodb.AttributeValueMemberB); ok {
		return uuid.UUID(id.Value)
	}
	return uuid.Nil
}

func getString(data dynamodb.AttributeValue) string {
	if data == nil {
		return ""
	}

	if s, ok := data.(*dynamodb.AttributeValueMemberS); ok {
		return s.Value
	}
	return ""
}

func getStringSlice(data dynamodb.AttributeValue) []string {
	if data == nil {
		return nil
	}

	if ss, ok := data.(*dynamodb.AttributeValueMemberSS); ok {
		return ss.Value
	}
	return nil
}

func getTime(data dynamodb.AttributeValue) time.Time {
	if data == nil {
		return time.Time{}
	}

	if n, ok := data.(*dynamodb.AttributeValueMemberN); ok {
		if i, err := strconv.ParseInt(n.Value, 10, 64); err == nil {
			return time.Unix(i, 0)
		}
	}
	return time.Time{}
}
