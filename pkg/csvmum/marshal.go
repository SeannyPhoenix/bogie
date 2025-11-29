package csvmum

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
)

type CSVMarshaler[T any] struct {
	writer    *csv.Writer
	fieldList []int
}

func NewMarshaler[T any](w io.Writer, opts ...marshalerOption) (*CSVMarshaler[T], error) {
	c := csv.NewWriter(w)

	return NewCSVMarshaler[T](c, opts...)
}

func NewCSVMarshaler[T any](w *csv.Writer, opts ...marshalerOption) (*CSVMarshaler[T], error) {
	m := &CSVMarshaler[T]{writer: w}

	for _, opt := range opts {
		opt(w)
	}

	var t T
	hm, err := buildFieldMap(reflect.TypeOf(t))
	if err != nil {
		return m, fmt.Errorf("cannot marshal: %w", err)
	}

	hh, fl := getOrderedHeaders(hm)
	if err = m.writer.Write(hh); err != nil {
		return m, fmt.Errorf("cannot marshal: %w", err)
	}

	m.fieldList = fl

	return m, nil
}

type marshalerOption func(*csv.Writer)

func WithDelimiter(delim rune) marshalerOption {
	return func(w *csv.Writer) {
		w.Comma = delim
	}
}

func (m *CSVMarshaler[T]) Marshal(record T) error {
	v := reflect.ValueOf(record)
	row := []string{}

	for _, fi := range m.fieldList {
		field := v.Field(fi)

		f, err := marshalValue(field)
		if err != nil {
			return fmt.Errorf("cannot marshal field %d: %w", fi, err)
		}
		row = append(row, f)
	}

	if err := m.writer.Write(row); err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}
	return nil
}

func (m *CSVMarshaler[T]) Flush() error {
	m.writer.Flush()

	if err := m.writer.Error(); err != nil {
		return fmt.Errorf("cannot marshal: %w", err)
	}

	return nil
}
