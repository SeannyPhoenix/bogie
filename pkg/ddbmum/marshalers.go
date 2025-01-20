package ddbmum

import (
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func marshalBool(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberBOOL{Value: v.Bool()}, nil
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

func marshalList(v reflect.Value) (types.AttributeValue, error) {
	l := v.Len()
	values := make([]types.AttributeValue, l)

	for i := 0; i < l; i++ {
		value, err := marshalValue(v.Index(i))
		if err != nil {
			return nil, err
		}

		values[i] = value
	}

	return &types.AttributeValueMemberL{Value: values}, nil
}

func marshalMap(v reflect.Value) (types.AttributeValue, error) {
	keys := v.MapKeys()
	values := make(map[string]types.AttributeValue, len(keys))

	for _, key := range keys {
		value, err := marshalValue(v.MapIndex(key))
		if err != nil {
			return nil, err
		}

		values[key.String()] = value
	}

	return &types.AttributeValueMemberM{Value: values}, nil
}

func marshalPointer(v reflect.Value) (types.AttributeValue, error) {
	if v.IsNil() {
		return nil, nil
	}

	return marshalValue(v.Elem())
}

func marshalString(v reflect.Value) (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: v.String()}, nil
}

func marshalStruct(v reflect.Value) (types.AttributeValue, error) {
	fields := v.NumField()
	values := make(map[string]types.AttributeValue, fields)

	for i := 0; i < fields; i++ {
		field := v.Type().Field(i)
		value, err := marshalValue(v.Field(i))
		if err != nil {
			return nil, err
		}

		values[field.Name] = value
	}

	return &types.AttributeValueMemberM{Value: values}, nil
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
		reflect.Bool:    marshalBool,
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
		reflect.Array:   marshalList,
		reflect.Map:     marshalMap,
		reflect.Pointer: marshalPointer,
		reflect.Slice:   marshalList,
		reflect.String:  marshalString,
		reflect.Struct:  marshalStruct,
	}
}
