package ddbmum

import (
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func marshalString(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: v.String()}, nil
}

func marshalInt(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberN{Value: strconv.FormatInt(v.Int(), 10)}, nil
}

func marshalUint(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberN{Value: strconv.FormatUint(v.Uint(), 10)}, nil
}

func marshalFloatFactory(size int) marshaler {
	return func(v reflect.Value) (types.AttributeValue, error) {
		return &types.AttributeValueMemberN{Value: strconv.FormatFloat(v.Float(), 'f', -1, size)}, nil
	}
}

func marshalBool(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberBOOL{Value: v.Bool()}, nil
}

func marshalPointer(v reflect.Value) (types.AttributeValue, error) {
	if v.IsNil() {
		return nil, nil
	}

	return marshalValue(v.Elem())
}

func marshalValue(v reflect.Value) (types.AttributeValue, error) {
	if marshaler, ok := defaultMarshalers[v.Kind()]; ok {
		return marshaler(v)
	}

	return nil, &UnsupportedTypeError{v.Kind()}
}

type marshaler func(reflect.Value) (types.AttributeValue, error)

var defaultMarshalers = map[reflect.Kind]marshaler{
	reflect.String: marshalString,
}

func init() {
	defaultMarshalers = map[reflect.Kind]marshaler{
		reflect.String:  marshalString,
		reflect.Int:     marshalInt,
		reflect.Int8:    marshalInt,
		reflect.Int16:   marshalInt,
		reflect.Int32:   marshalInt,
		reflect.Int64:   marshalInt,
		reflect.Uint:    marshalUint,
		reflect.Uint8:   marshalUint,
		reflect.Uint16:  marshalUint,
		reflect.Uint32:  marshalUint,
		reflect.Uint64:  marshalUint,
		reflect.Float32: marshalFloatFactory(32),
		reflect.Float64: marshalFloatFactory(64),
		reflect.Bool:    marshalBool,
		reflect.Pointer: marshalPointer,
	}
}
