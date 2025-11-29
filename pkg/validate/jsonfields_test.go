package validate_test

import (
	"encoding/json"
	"testing"

	"github.com/seannyphoenix/bogie/pkg/validate"
	"github.com/stretchr/testify/assert"
)

func TestRequiredFields_Success(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		json   string
		fields []string
	}{
		{
			name:   "single field present",
			json:   `{"name":"John"}`,
			fields: []string{"name"},
		},
		{
			name:   "multiple fields present",
			json:   `{"name":"John","age":30,"city":"NYC"}`,
			fields: []string{"name", "age", "city"},
		},
		{
			name:   "field with null value is present",
			json:   `{"name":null,"age":30}`,
			fields: []string{"name", "age"},
		},
		{
			name:   "field with empty string is present",
			json:   `{"name":"","age":30}`,
			fields: []string{"name", "age"},
		},
		{
			name:   "field with zero value is present",
			json:   `{"age":0,"active":false}`,
			fields: []string{"age", "active"},
		},
		{
			name:   "nested object field present",
			json:   `{"user":{"name":"John"},"age":30}`,
			fields: []string{"user", "age"},
		},
		{
			name:   "array field present",
			json:   `{"tags":["go","rust"],"count":2}`,
			fields: []string{"tags", "count"},
		},
		{
			name:   "no fields required",
			json:   `{"name":"John"}`,
			fields: []string{},
		},
		{
			name:   "extra fields ignored",
			json:   `{"name":"John","age":30,"extra":"data"}`,
			fields: []string{"name", "age"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			err := validate.RequiredFields([]byte(tc.json), tc.fields...)
			assert.NoError(err)
		})
	}
}

func TestRequiredFields_MissingFields(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name         string
		json         string
		fields       []string
		missingField string
	}{
		{
			name:         "single field missing",
			json:         `{"age":30}`,
			fields:       []string{"name"},
			missingField: "name",
		},
		{
			name:         "one of multiple fields missing",
			json:         `{"name":"John","city":"NYC"}`,
			fields:       []string{"name", "age", "city"},
			missingField: "age",
		},
		{
			name:         "all fields missing",
			json:         `{}`,
			fields:       []string{"name", "age"},
			missingField: "name",
		},
		{
			name:         "field missing from nested object",
			json:         `{"user":{"age":30}}`,
			fields:       []string{"name"},
			missingField: "name",
		},
		{
			name:         "case sensitive - wrong case",
			json:         `{"Name":"John"}`,
			fields:       []string{"name"},
			missingField: "name",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			err := validate.RequiredFields([]byte(tc.json), tc.fields...)
			assert.Error(err)
			assert.ErrorContains(err, "missing required field")
			assert.ErrorContains(err, tc.missingField)
		})
	}
}

func TestRequiredFields_InvalidJSON(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name   string
		json   string
		fields []string
	}{
		{
			name:   "malformed JSON",
			json:   `{invalid}`,
			fields: []string{"name"},
		},
		{
			name:   "incomplete JSON",
			json:   `{"name":"John"`,
			fields: []string{"name"},
		},
		{
			name:   "trailing comma",
			json:   `{"name":"John",}`,
			fields: []string{"name"},
		},
		{
			name:   "not an object - string",
			json:   `"not an object"`,
			fields: []string{"name"},
		},
		{
			name:   "not an object - number",
			json:   `42`,
			fields: []string{"name"},
		},
		{
			name:   "not an object - array",
			json:   `["not","an","object"]`,
			fields: []string{"name"},
		},
		{
			name:   "not an object - boolean",
			json:   `true`,
			fields: []string{"name"},
		},
		{
			name:   "empty string",
			json:   ``,
			fields: []string{"name"},
		},
		{
			name:   "null at root",
			json:   `null`,
			fields: []string{"name"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			assert := assert.New(t)

			err := validate.RequiredFields([]byte(tc.json), tc.fields...)
			assert.Error(err)
			assert.ErrorContains(err, "validate fields")
		})
	}
}

func TestRequiredFields_EdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("empty field name allowed", func(t *testing.T) {
		assert := assert.New(t)

		json := `{"":"value","name":"John"}`
		err := validate.RequiredFields([]byte(json), "", "name")
		assert.NoError(err)
	})

	t.Run("unicode field names", func(t *testing.T) {
		assert := assert.New(t)

		json := `{"名前":"太郎","age":30}`
		err := validate.RequiredFields([]byte(json), "名前", "age")
		assert.NoError(err)
	})

	t.Run("special characters in field names", func(t *testing.T) {
		assert := assert.New(t)

		json := `{"field-name":"value","field.name":"value2","field_name":"value3"}`
		err := validate.RequiredFields([]byte(json), "field-name", "field.name", "field_name")
		assert.NoError(err)
	})

	t.Run("duplicate fields in requirements", func(t *testing.T) {
		assert := assert.New(t)

		json := `{"name":"John"}`
		err := validate.RequiredFields([]byte(json), "name", "name", "name")
		assert.NoError(err)
	})

	t.Run("numeric field names", func(t *testing.T) {
		assert := assert.New(t)

		json := `{"123":"value","456":"value2"}`
		err := validate.RequiredFields([]byte(json), "123", "456")
		assert.NoError(err)
	})
}

func TestRequiredFields_RealWorldExamples(t *testing.T) {
	t.Parallel()

	t.Run("user registration", func(t *testing.T) {
		assert := assert.New(t)

		validUser := `{"username":"john","email":"john@example.com","password":"secret"}`
		err := validate.RequiredFields([]byte(validUser), "username", "email", "password")
		assert.NoError(err)

		missingEmail := `{"username":"john","password":"secret"}`
		err = validate.RequiredFields([]byte(missingEmail), "username", "email", "password")
		assert.Error(err)
		assert.ErrorContains(err, "email")
	})

	t.Run("API request validation", func(t *testing.T) {
		assert := assert.New(t)

		validRequest := `{"method":"POST","url":"https://api.example.com","headers":{},"body":null}`
		err := validate.RequiredFields([]byte(validRequest), "method", "url")
		assert.NoError(err)
	})

	t.Run("configuration file", func(t *testing.T) {
		assert := assert.New(t)

		validConfig := `{"host":"localhost","port":8080,"database":"mydb","timeout":30}`
		err := validate.RequiredFields([]byte(validConfig), "host", "port", "database")
		assert.NoError(err)
	})
}

func TestRequiredFields_TypePreservation(t *testing.T) {
	t.Parallel()

	t.Run("does not validate field types", func(t *testing.T) {
		assert := assert.New(t)

		// RequiredFields only checks presence, not type correctness
		jsonWithWrongTypes := `{"age":"not a number","active":"not a boolean"}`
		err := validate.RequiredFields([]byte(jsonWithWrongTypes), "age", "active")
		assert.NoError(err, "RequiredFields should only check presence, not types")
	})

	t.Run("validates presence before type unmarshaling", func(t *testing.T) {
		assert := assert.New(t)

		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		// First check required fields
		jsonData := []byte(`{"name":"John"}`)
		err := validate.RequiredFields(jsonData, "name", "age")
		assert.Error(err, "should catch missing 'age' before attempting to unmarshal")

		// If we skip validation, unmarshaling would succeed with zero value
		var user User
		err = json.Unmarshal(jsonData, &user)
		assert.NoError(err, "standard unmarshal allows missing fields")
		assert.Equal("John", user.Name)
		assert.Equal(0, user.Age, "missing field gets zero value")
	})
}
