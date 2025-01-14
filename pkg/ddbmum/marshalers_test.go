package ddbmum

import (
	"errors"
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
		"string":      {"string", &types.AttributeValueMemberS{Value: "string"}, nil},
		"int":         {27, &types.AttributeValueMemberN{Value: "27"}, nil},
		"int8":        {int8(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int16":       {int16(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int32":       {int32(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"int64":       {int64(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint":        {uint(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint8":       {uint8(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint16":      {uint16(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint32":      {uint32(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"uint64":      {uint64(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"float32":     {float32(3.14), &types.AttributeValueMemberN{Value: "3.14"}, nil},
		"float64":     {3.14, &types.AttributeValueMemberN{Value: "3.14"}, nil},
		"bool":        {true, &types.AttributeValueMemberBOOL{Value: true}, nil},
		"pointer":     {util.Ptr(27), &types.AttributeValueMemberN{Value: "27"}, nil},
		"pointerNil":  {util.NilPtr[int](), nil, nil},
		"unsupported": {struct{}{}, nil, errors.New("unsupported type: struct")},
	}

	for name, tc := range tt {
		tc := tc

		t.Run(name, func(t *testing.T) {
			t.Parallel()
			assert := assert.New(t)

			v, err := marshalValue(reflect.ValueOf(tc.field))

			if tc.err != nil {
				assert.EqualError(err, tc.err.Error())
			} else {
				assert.NoError(err)
			}

			assert.Equal(tc.expected, v)
		})
	}
}
