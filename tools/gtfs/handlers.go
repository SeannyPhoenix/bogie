package main

import (
	"fmt"
	"os"

	"github.com/seannyphoenix/bogie/pkg/gtfs"
)

type handler func(col map[string]gtfs.GTFSSchedule) error

func printScheduleOverviews(col map[string]gtfs.GTFSSchedule) error {
	fmt.Println(gtfs.Overview(col))
	return nil
}

func writeErrors(col map[string]gtfs.GTFSSchedule) error {
	errFile, err := os.Create("gtfs_files/gtfs_errors.txt")
	if err != nil {
		return fmt.Errorf("Error creating error file: %s\n", err.Error())
	}
	defer errFile.Close()

	for _, e := range col {
		for _, err := range e.Errors() {
			_, err := errFile.WriteString(fmt.Sprintf("%s\n", err))
			if err != nil {
				return fmt.Errorf("Error writing to error file: %s\n", err.Error())
			}
		}
	}

	return nil
}
