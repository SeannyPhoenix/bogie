package ddbmum

import (
	"errors"
	"reflect"
	"testing"
	"unsafe"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMarshalValue(t *testing.T) {
	t.Parallel()

	tt := map[string]struct {
		field    any
		expected types.AttributeValue
		err      error
	}{
		"bool":       {true, &types.AttributeValueMemberBOOL{Value: true}, nil},
		"int":        {27, &types.AttributeValueMemberN{Value: "27"}, nil},
		"int8":       {int8(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int16":      {int16(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int32":      {int32(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int64":      {int64(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint":       {uint(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint8":      {uint8(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint16":     {uint16(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint32":     {uint32(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint64":     {uint64(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"float32":    {float32(3.14), &types.AttributeValueMemberN{Value: "3.14"}, nil},
		"float64":    {3.14, &types.AttributeValueMemberN{Value: "3.14"}, nil},
		"complex64":  {complex64(3.14 + 2.71i), nil, errors.New("unsupported type: complex64")},
		"complex128": {3.14 + 2.71i, nil, errors.New("unsupported type: complex128")},
		"array": {[3]string{"a", "b", "c"}, &types.AttributeValueMemberL{Value: []types.AttributeValue{
			&types.AttributeValueMemberS{Value: "a"},
			&types.AttributeValueMemberS{Value: "b"},
			&types.AttributeValueMemberS{Value: "c"},
		}}, nil},
		"chan": {make(chan int), nil, errors.New("unsupported type: chan")},
		"func": {func() {}, nil, errors.New("unsupported type: func")},
		// "interface"
		"map": {map[string]int{"a": 1, "b": 2}, &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"a": &types.AttributeValueMemberN{Value: "1"},
			"b": &types.AttributeValueMemberN{Value: "2"},
		}}, nil},
		"mapIntKey": {map[int]string{1: "a", 2: "b"}, &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"1": &types.AttributeValueMemberS{Value: "a"},
			"2": &types.AttributeValueMemberS{Value: "b"},
		}}, nil},
		"pointer":    {util.Ptr(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"pointerNil": {util.NilPtr[int](), nil, nil},
		"slice": {[]int{1, 2, 3}, &types.AttributeValueMemberL{Value: []types.AttributeValue{
			&types.AttributeValueMemberN{Value: "1"},
			&types.AttributeValueMemberN{Value: "2"},
			&types.AttributeValueMemberN{Value: "3"},
		}}, nil},
		"string": {"string", &types.AttributeValueMemberS{Value: "string"}, nil},
		"struct": {struct {
			A int
			B bool
		}{27, true}, &types.AttributeValueMemberM{Value: map[string]types.AttributeValue{
			"A": &types.AttributeValueMemberN{Value: "27"},
			"B": &types.AttributeValueMemberBOOL{Value: true},
		}}, nil},
		"unsafePointer":   {unsafe.Pointer(util.Ptr(27)), nil, errors.New("unsupported type: unsafe.Pointer")},
		"binaryMarshaler": {uuid.MustParse("12db069a-c8dc-4a5a-af0c-22282ed89f3c"), &types.AttributeValueMemberB{Value: []byte{0x12, 0xdb, 0x06, 0x9a, 0xc8, 0xdc, 0x4a, 0x5a, 0xaf, 0x0c, 0x22, 0x28, 0x2e, 0xd8, 0x9f, 0x3c}}, nil},
	}

	for name, tc := range tt {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)

			v, err := marshalValue(reflect.ValueOf(tc.field))

			if tc.err != nil {
				var e *UnsupportedTypeError
				assert.ErrorAs(err, &e)
				assert.EqualError(err, tc.err.Error())
			} else {
				assert.NoError(err)
			}

			assert.Equal(tc.expected, v)
		})
	}
}
