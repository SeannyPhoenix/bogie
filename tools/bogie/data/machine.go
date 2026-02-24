package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"

	"github.com/google/uuid"
	"github.com/seannyphoenix/bogie/pkg/csvmum"
)

type machine struct {
	in    *os.File
	unm   *csvmum.CSVUnmarshaler[EventV2CSV]
	tab   *os.File
	mt    *csvmum.CSVMarshaler[EventV2CSV]
	comma *os.File
	mc    *csvmum.CSVMarshaler[EventV2CSV]
	json  *os.File

	closers []func() error
}

func newIO() machine {
	var m machine
	var err error

	m.in, err = os.Open("in/events.csv")
	if err != nil {
		panic(err)
	}
	m.closers = append(m.closers, m.in.Close)

	m.unm, err = csvmum.NewUnmarshaler[EventV2CSV](m.in)
	if err != nil {
		panic(err)
	}

	m.tab, err = os.Create("out/events.tsv")
	if err != nil {
		panic(err)
	}
	m.closers = append(m.closers, m.tab.Close)

	m.mt, err = csvmum.NewMarshaler[EventV2CSV](m.tab, csvmum.WithDelimiter('\t'))
	if err != nil {
		panic(err)
	}
	m.closers = append([]func() error{m.mt.Flush}, m.closers...)

	m.comma, err = os.Create("out/events.csv")
	if err != nil {
		panic(err)
	}
	m.closers = append(m.closers, m.comma.Close)

	m.mc, err = csvmum.NewMarshaler[EventV2CSV](m.comma)
	if err != nil {
		panic(err)
	}
	m.closers = append([]func() error{m.mc.Flush}, m.closers...)

	m.json, err = os.Create("out/events.jsonl")
	if err != nil {
		panic(err)
	}
	m.closers = append(m.closers, m.json.Close)

	return m
}

func (m *machine) closeAll() error {
	var errs []error
	for _, closer := range m.closers {
		if err := closer(); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func (m *machine) process() error {
	var last EventV2CSV
	var lastID uuid.UUID

	for {
		current, err := m.processCurrent(last)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		err = m.processJSON(current, &lastID)
		if err != nil {
			return err
		}

		last = current
	}

	return nil
}

func (m *machine) processCurrent(last EventV2CSV) (EventV2CSV, error) {
	var current EventV2CSV
	err := m.unm.Unmarshal(&current)
	if err != nil {
		// handle EOF above
		return current, err
	}

	current.fillDate(last)
	current.setTimestamp()
	current.checkOrder(last)

	err = m.mt.Marshal(current)
	if err != nil {
		return current, err
	}

	err = m.mc.Marshal(current)
	if err != nil {
		return current, err
	}

	return current, nil
}

func (m *machine) processJSON(current EventV2CSV, lastID *uuid.UUID) error {
	e := event{
		ID:        uuid.New(),
		Previous:  lastID,
		Begin:     current.Begin,
		EventType: current.EventType,
		Timestamp: current.binarytime,
		Location: location{
			Name:     current.Location,
			Platform: current.Platform,
			StopID:   current.StopID,
			Agency:   current.Agency,
		},
		Run: run{
			Route:     current.Route,
			Run:       current.Run,
			Operator:  current.Operator,
			VehicleID: current.VehicleID,
			Units:     getUnits(current),
		},
		Notes: getNotes(current),
	}

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	_, err = m.json.Write(append(b, '\n'))
	if err != nil {
		return err
	}

	*lastID = e.ID
	return nil
}
