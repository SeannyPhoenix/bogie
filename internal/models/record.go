package models

import (
	"time"

	"github.com/google/uuid"
)

type Record struct {
	Id        uuid.UUID  `json:"id" dynamodbav:"id"`
	Type      string     `json:"type" dynamodbav:"t"`
	Status    string     `json:"status" dynamodbav:"s"`
	CreatedAt time.Time  `json:"createdAt" dynamodbav:"ca"`
	UpdatedAt time.Time  `json:"updatedAt" dynamodbav:"ua"`
	User      *uuid.UUID `json:"user,omitempty" dynamodbav:"u,omitempty"`
}

func newRecord(t string, user *uuid.UUID) Record {
	now := time.Now()
	return Record{
		Id:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Status:    DocStatusActive,
		Type:      t,
		User:      user,
	}
}

const (
	DocTypeEvent = "event"
	DocTypeUnit  = "unit"
	DocTypeUser  = "user"
)

const (
	DocStatusActive   = "active"
	DocStatusInactive = "inactive"
)

var NameMap = map[string]uuid.UUID{
	DocTypeEvent:      uuid.MustParse("88c2333e-2bc2-4063-b865-719c24211d2c"),
	DocTypeUnit:       uuid.MustParse("915ddb34-93ba-4e2f-99f2-ea814bb2790d"),
	DocTypeUser:       uuid.MustParse("8d930d82-24e9-4fc1-824b-9e1253d4ee02"),
	DocStatusActive:   uuid.MustParse("5c96e882-112a-42e3-adf3-941f28ff9956"),
	DocStatusInactive: uuid.MustParse("aca32352-08c1-40ee-8a6e-e95d77e68724"),
}

var IDMap map[uuid.UUID]string

func init() {
	IDMap = make(map[uuid.UUID]string, len(NameMap))
	for k, v := range NameMap {
		IDMap[v] = k
	}
}
