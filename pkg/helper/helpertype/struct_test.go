package helpertype_test

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"

	"ddd/pkg/helper/helpertest/mock"
	"ddd/pkg/helper/helpertype"
)

func TestStructTool_FilterZeroValueField(t *testing.T) {
	// helperlog.DevelopSetting()

	type QueryCondition struct {
		CreatedTime   helpertype.Time `db:"created_time"`
		UserName      string          `db:"user_name"`
		Orders        []string        `db:"order"`
		Age           *int            `db:"age"`
		NullableValue interface{}     `db:"money"`
	}

	tests := []struct {
		name        string
		rawStruct   *QueryCondition
		expectedMap map[helpertype.FieldName]interface{}
	}{
		{
			rawStruct: &QueryCondition{
				UserName: "caesar",
				Orders:   []string{"book", "tea"},
			},
			expectedMap: map[helpertype.FieldName]interface{}{
				"user_name": "caesar",
				"order":     []string{"book", "tea"},
			},
		},
		{
			rawStruct: &QueryCondition{
				CreatedTime: mock.Time("2020-08-23"),
				UserName:    "caesar",
			},
			expectedMap: map[helpertype.FieldName]interface{}{
				"created_time": mock.Time("2020-08-23"),
				"user_name":    "caesar",
			},
		},
		{
			name:        "All Zero value fields # Struct Not Nil",
			rawStruct:   &QueryCondition{},
			expectedMap: map[string]interface{}{},
		},
		{
			name:        "All Zero value fields # Struct Is Nil",
			rawStruct:   nil,
			expectedMap: map[string]interface{}{},
		},
		{
			name: "field have value but is js null",
			rawStruct: &QueryCondition{
				NullableValue: (*float64)(nil), // note!! reflect.ValueOf(field).IsZero=false
			},
			expectedMap: map[string]interface{}{
				"money": (*float64)(nil),
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			log.Debug().Msgf("\n%v", spew.Sdump(tt.rawStruct))

			actualMap := helpertype.StructTool{}.FilterZeroValueField(tt.rawStruct, "db")
			assert.Equal(t, tt.expectedMap, actualMap)
		})
	}
}
