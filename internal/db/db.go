package db

import (
	"errors"
	"strconv"
	"time"

	dynamodb "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

// Document Fields
const (
	Agency        = "a"
	ArrivalStop   = "as"
	ArrivalTime   = "at"
	CreatedAt     = "ca"
	DepartureStop = "ds"
	DepartureTime = "dt"
	ID            = "id"
	Notes         = "n"
	Route         = "r"
	Status        = "s"
	Type          = "t"
	Trip          = "tr"
	UnitID        = "u"
	UpdatedAt     = "ua"
	UnitCount     = "uc"
	UnitPosition  = "up"
	UserID        = "uid"
)

// Deserializers
func GetUUID(data dynamodb.AttributeValue) uuid.UUID {
	if data == nil {
		return uuid.Nil
	}

	if id, ok := data.(*dynamodb.AttributeValueMemberB); ok {
		value, err := uuid.FromBytes(id.Value)
		if err != nil {
			return uuid.Nil
		}
		return value
	}
	return uuid.Nil
}

func GetString(data dynamodb.AttributeValue) string {
	if data == nil {
		return ""
	}

	if s, ok := data.(*dynamodb.AttributeValueMemberS); ok {
		return s.Value
	}
	return ""
}

func GetStringSlice(data dynamodb.AttributeValue) []string {
	if data == nil {
		return nil
	}

	if ss, ok := data.(*dynamodb.AttributeValueMemberSS); ok {
		return ss.Value
	}
	return nil
}

func GetTime(data dynamodb.AttributeValue) time.Time {
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

func GetIntPtr(data dynamodb.AttributeValue) *int {
	if data == nil {
		return nil
	}

	if n, ok := data.(*dynamodb.AttributeValueMemberN); ok {
		if i, err := strconv.Atoi(n.Value); err == nil {
			return &i
		}
	}
	return nil
}

// Errors
var (
	ErrBadDocID     = errors.New("bad document ID")
	ErrBadDocType   = errors.New("bad document type")
	ErrBadDocStatus = errors.New("bad document status")
	ErrBadCreatedAt = errors.New("bad created at time")
	ErrBadUpdatedAt = errors.New("bad updated at time")
	ErrBadUserID    = errors.New("bad user ID")
	ErrBadCarrier   = errors.New("bad carrier")
	ErrBadUnitID    = errors.New("bad unit ID")
)

// DynamoDB Constants
const (
	DynamoDBBatchWriteLimit = 25
)

type DBDocument map[string]dynamodb.AttributeValue
