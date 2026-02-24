package models_test

import (
	"fmt"
	"testing"

	"github.com/seannyphoenix/bogie/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewGeoCoordinatesError(t *testing.T) {
	t.Parallel()

	tt := []struct {
		lat       float64
		lon       float64
		expectErr error
	}{
		{lat: -91, lon: 0, expectErr: models.InvalidLatitude},
		{lat: 91, lon: 0, expectErr: models.InvalidLatitude},
		{lat: 0, lon: -181, expectErr: models.InvalidLongitude},
		{lat: 0, lon: 181, expectErr: models.InvalidLongitude},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprint([]float64{tc.lat, tc.lon}), func(t *testing.T) {
			assert := assert.New(t)

			c, err := models.NewGeoCoordinates(tc.lat, tc.lon)
			assert.ErrorIs(err, tc.expectErr)
			assert.False(c.IsValid())
		})
	}
}

func TestGeoCoordinates_LatLon(t *testing.T) {
	t.Parallel()

	tt := []struct {
		lat float64
		lon float64
	}{
		{lat: 0, lon: 0},
		{lat: 45.1234, lon: -93.1234},
		{lat: -89.9999, lon: 179.9999},
		{lat: 90, lon: -180},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprint([]float64{tc.lat, tc.lon}), func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			coords, err := models.NewGeoCoordinates(tc.lat, tc.lon)
			require.NoError(err)
			assert.True(coords.IsValid())

			lat, lon := coords.LatLon()
			assert.Equal(tc.lat, lat)
			assert.Equal(tc.lon, lon)
		})
	}
}

func TestGeoCoordinates_MarshalJSON(t *testing.T) {
	t.Parallel()

	tt := []struct {
		coord models.GeoCoordinates
		json  string
	}{
		{
			coord: models.MustNewGeoCoordinates(0, 0),
			json:  `{"lat":0,"lon":0}`,
		},
		{
			coord: models.MustNewGeoCoordinates(45.1234, -93.1234),
			json:  `{"lat":45.1234,"lon":-93.1234}`,
		},
		{
			coord: models.GeoCoordinates{},
			json:  `null`,
		},
	}

	for _, tc := range tt {
		t.Run(tc.json, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			data, err := tc.coord.MarshalJSON()
			require.NoError(err)
			assert.JSONEq(tc.json, string(data))
		})
	}
}

func TestGeoCoordinates_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tt := []struct {
		json  string
		coord models.GeoCoordinates
	}{
		{
			json:  `{"lat":0,"lon":0}`,
			coord: models.MustNewGeoCoordinates(0, 0),
		},
		{
			json:  `{"lat":45.1234,"lon":-93.1234}`,
			coord: models.MustNewGeoCoordinates(45.1234, -93.1234),
		},
		{
			json:  `null`,
			coord: models.GeoCoordinates{},
		},
	}

	for _, tc := range tt {
		t.Run(tc.json, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			var coord models.GeoCoordinates
			err := coord.UnmarshalJSON([]byte(tc.json))
			require.NoError(err)
			assert.Equal(tc.coord, coord)
		})
	}
}

func TestGeoCoordinates_MarshalUnmarshalRoundTrip(t *testing.T) {
	t.Parallel()

	tt := []models.GeoCoordinates{
		models.MustNewGeoCoordinates(0, 0),
		models.MustNewGeoCoordinates(45.1234, -93.1234),
		models.MustNewGeoCoordinates(-89.9999, 179.9999),
		models.MustNewGeoCoordinates(90, -180),
	}

	for _, tc := range tt {
		t.Run(fmt.Sprint([]float64{tc.Lat(), tc.Lon()}), func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			data, err := tc.MarshalJSON()
			require.NoError(err)

			var coord models.GeoCoordinates
			err = coord.UnmarshalJSON(data)
			require.NoError(err)
			assert.Equal(tc, coord)
		})
	}
}

func TestGeoCoordinates_UnmarshalMarshalRoundTrip(t *testing.T) {
	t.Parallel()

	tt := []string{
		`{"lat":0,"lon":0}`,
		`{"lat":45.1234,"lon":-93.1234}`,
		`{"lat":-89.9999,"lon":179.9999}`,
		`{"lat":90,"lon":-180}`,
	}

	for _, tc := range tt {
		t.Run(tc, func(t *testing.T) {
			assert := assert.New(t)
			require := require.New(t)

			var coord models.GeoCoordinates
			err := coord.UnmarshalJSON([]byte(tc))
			require.NoError(err)

			data, err := coord.MarshalJSON()
			require.NoError(err)
			assert.JSONEq(tc, string(data))
		})
	}
}

func TestGeoCoordinates_UnmarshalJSONError(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name string
		json string
	}{
		// Out of range values
		{name: "latitude too low", json: `{"lat":-91,"lon":0}`},
		{name: "latitude too high", json: `{"lat":91,"lon":0}`},
		{name: "longitude too low", json: `{"lat":0,"lon":-181}`},
		{name: "longitude too high", json: `{"lat":0,"lon":181}`},

		// Missing fields
		{name: "missing longitude", json: `{"lat":45}`},
		{name: "missing latitude", json: `{"lon":90}`},
		{name: "empty object", json: `{}`},

		// Wrong types
		{name: "lat as string", json: `{"lat":"45.0","lon":-93.0}`},
		{name: "lon as string", json: `{"lat":45.0,"lon":"-93.0"}`},
		{name: "lat as boolean", json: `{"lat":true,"lon":0}`},
		{name: "lon as boolean", json: `{"lat":0,"lon":false}`},
		{name: "lat as array", json: `{"lat":[45.0],"lon":-93.0}`},
		{name: "lon as array", json: `{"lat":45.0,"lon":[-93.0]}`},
		{name: "lat as object", json: `{"lat":{},"lon":0}`},
		{name: "lon as object", json: `{"lat":0,"lon":{}}`},

		// Malformed JSON
		{name: "invalid JSON syntax", json: `not json`},
		{name: "incomplete object", json: `{"lat":45`},
		{name: "trailing comma", json: `{"lat":45,"lon":90,}`},
		{name: "missing quotes", json: `{lat:45,lon:90}`},
		{name: "single quotes", json: `{'lat':45,'lon':90}`},

		// Wrong root types
		{name: "string instead of object", json: `"coordinates"`},
		{name: "number instead of object", json: `42`},
		{name: "boolean instead of object", json: `true`},
		{name: "array instead of object", json: `[45.0, -93.0]`},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			var coord models.GeoCoordinates
			err := coord.UnmarshalJSON([]byte(tc.json))
			assert.Error(err, "expected error for: %s", tc.json)
			assert.False(coord.IsValid())
		})
	}
}

func BenchmarkGeoCoordinates_MarshalJSON(b *testing.B) {
	coords := models.MustNewGeoCoordinates(45.1234, -93.1234)

	for b.Loop() {
		_, _ = coords.MarshalJSON()
	}
}

func BenchmarkGeoCoordinates_UnmarshalJSON(b *testing.B) {
	data := []byte(`{"lat":45.1234,"lon":-93.1234}`)
	var coords models.GeoCoordinates

	for b.Loop() {
		_ = coords.UnmarshalJSON(data)
	}
}

func BenchmarkGeoCoordinates_MarshalUnmarshalRoundTrip(b *testing.B) {
	coords := models.MustNewGeoCoordinates(45.1234, -93.1234)

	for b.Loop() {
		data, _ := coords.MarshalJSON()
		var c models.GeoCoordinates
		_ = c.UnmarshalJSON(data)
	}
}

func BenchmarkGeoCoordinates_NewGeoCoordinates(b *testing.B) {

	for b.Loop() {
		_, _ = models.NewGeoCoordinates(45.1234, -93.1234)
	}
}
