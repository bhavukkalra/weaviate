//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

package sorter

import (
	"reflect"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/semi-technologies/weaviate/entities/models"
	"github.com/semi-technologies/weaviate/entities/schema"
	"github.com/semi-technologies/weaviate/entities/storobj"
	"github.com/stretchr/testify/require"
)

func Test_lsmPropertyExtractor_getProperty(t *testing.T) {
	className := "MyFavoriteClass"
	sch := schema.Schema{
		Objects: &models.Schema{
			Classes: []*models.Class{
				{
					Class: className,
					Properties: []*models.Property{
						{
							Name:     "stringProp",
							DataType: []string{string(schema.DataTypeString)},
						},
						{
							Name:     "textProp",
							DataType: []string{string(schema.DataTypeText)},
						},
						{
							Name:     "stringPropArray",
							DataType: []string{string(schema.DataTypeStringArray)},
						},
						{
							Name:     "textPropArray",
							DataType: []string{string(schema.DataTypeTextArray)},
						},
						{
							Name:     "intProp",
							DataType: []string{string(schema.DataTypeInt)},
						},
						{
							Name:     "numberProp",
							DataType: []string{string(schema.DataTypeNumber)},
						},
						{
							Name:     "intPropArray",
							DataType: []string{string(schema.DataTypeIntArray)},
						},
						{
							Name:     "numberPropArray",
							DataType: []string{string(schema.DataTypeNumberArray)},
						},
						{
							Name:     "boolProp",
							DataType: []string{string(schema.DataTypeBoolean)},
						},
						{
							Name:     "boolPropArray",
							DataType: []string{string(schema.DataTypeBooleanArray)},
						},
						{
							Name:     "dateProp",
							DataType: []string{string(schema.DataTypeDate)},
						},
						{
							Name:     "datePropArray",
							DataType: []string{string(schema.DataTypeDateArray)},
						},
						{
							Name:     "phoneProp",
							DataType: []string{string(schema.DataTypePhoneNumber)},
						},
						{
							Name:     "geoProp",
							DataType: []string{string(schema.DataTypeGeoCoordinates)},
						},
						{
							Name:     "emptyStringProp",
							DataType: []string{string(schema.DataTypeString)},
						},
						{
							Name:     "crefProp",
							DataType: []string{string(schema.DataTypeCRef)},
						},
					},
				},
			},
		},
	}

	obj := storobj.FromObject(
		&models.Object{
			Class:              className,
			CreationTimeUnix:   123456,
			LastUpdateTimeUnix: 56789,
			ID:                 strfmt.UUID("73f2eb5f-5abf-447a-81ca-74b1dd168247"),
			Properties: map[string]interface{}{
				"stringProp":      "string",
				"textProp":        "text",
				"stringPropArray": []string{"string", "string"},
				"textPropArray":   []string{"text", "text"},
				"intProp":         100,
				"numberProp":      float64(17),
				"intPropArray":    []int{10, 20, 30},
				"numberPropArray": []float64{1, 2, 3},
				"boolProp":        true,
				"boolPropArray":   []bool{true, false, true},
				"dateProp":        "1980-01-01T00:00:00+02:00",
				"datePropArray":   []string{"1980-01-01T00:00:00+02:00"},
				"phoneProp": &models.PhoneNumber{
					CountryCode:            49,
					DefaultCountry:         "DE",
					Input:                  "0171 1000000",
					Valid:                  true,
					InternationalFormatted: "+49 171 1000000",
					National:               1000000,
					NationalFormatted:      "0171 1000000",
				},
				"geoProp": &models.GeoCoordinates{
					Longitude: pointerFloat32(1),
					Latitude:  pointerFloat32(2),
				},
			},
		},
		[]float32{1, 2, 0.7},
	)
	asBinary, err := obj.MarshalBinary()
	require.Nil(t, err)

	tests := []struct {
		name     string
		property string
		want     interface{}
	}{
		{
			name:     "stringProp",
			property: "stringProp",
			want:     "string",
		},
		{
			name:     "textProp",
			property: "textProp",
			want:     "text",
		},
		{
			name:     "stringPropArray",
			property: "stringPropArray",
			want:     []string{"string", "string"},
		},
		{
			name:     "textPropArray",
			property: "textPropArray",
			want:     []string{"text", "text"},
		},
		{
			name:     "intProp",
			property: "intProp",
			want:     float64(100),
		},
		{
			name:     "numberProp",
			property: "numberProp",
			want:     float64(17),
		},
		{
			name:     "intPropArray",
			property: "intPropArray",
			want:     []float64{10, 20, 30},
		},
		{
			name:     "numberPropArray",
			property: "numberPropArray",
			want:     []float64{1, 2, 3},
		},
		{
			name:     "boolProp",
			property: "boolProp",
			want:     true,
		},
		{
			name:     "boolPropArray",
			property: "boolPropArray",
			want:     []bool{true, false, true},
		},
		{
			name:     "dateProp",
			property: "dateProp",
			want:     "1980-01-01T00:00:00+02:00",
		},
		{
			name:     "datePropArray",
			property: "datePropArray",
			want:     []string{"1980-01-01T00:00:00+02:00"},
		},
		{
			name:     "phoneProp",
			property: "phoneProp",
			want: &models.PhoneNumber{
				CountryCode:            49,
				DefaultCountry:         "DE",
				Input:                  "0171 1000000",
				Valid:                  true,
				InternationalFormatted: "+49 171 1000000",
				National:               1000000,
				NationalFormatted:      "0171 1000000",
			},
		},
		{
			name:     "geoProp",
			property: "geoProp",
			want: &models.GeoCoordinates{
				Longitude: pointerFloat32(1),
				Latitude:  pointerFloat32(2),
			},
		},
		{
			name:     "emptyStringProp",
			property: "emptyStringProp",
			want:     nil,
		},
		{
			name:     "crefProp",
			property: "crefProp",
			want:     nil,
		},
		{
			name:     "nonExistentProp",
			property: "nonExistentProp",
			want:     nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := newPropertyExtractor(schema.ClassName(className), newClassHelper(sch), tt.property)
			if got := e.getProperty(asBinary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lsmPropertyExtractor.getProperty() = %v, want %v", got, tt.want)
			}
		})
	}
}
