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
			_, err := fmt.Fprintf(errFile, "%s\n", err)
			if err != nil {
				return fmt.Errorf("Error writing to error file: %s\n", err.Error())
			}
		}
	}

	return nil
}

func printStopNames(col map[string]gtfs.GTFSSchedule) error {
	if len(col) == 0 {
		return nil
	}

	for _, s := range col {
		for _, agency := range s.Agencies {
			fmt.Printf("Stop names for agency %s\n", agency.Name)
			break
		}
		// names := gtfs.StopNames(s.Stops)
		// for _, name := range names {
		// 	fmt.Printf("\t%s\n", name)
		// }

		tree := gtfs.BuildStopTree(s.Stops)
		for _, stop := range tree {
			fmt.Printf("  %s\n", stop.Stop.Name)
			fmt.Printf("    ID: %s\n", stop.Stop.ID)
			for childID, child := range stop.Children {
				fmt.Printf("    Child stop %s: %s\n", childID, child.Name)
			}
		}
	}
	return nil
}

func printStopInfo(col map[string]gtfs.GTFSSchedule, stopID string) error {
	if len(col) == 0 {
		return nil
	}

	for _, s := range col {
		si := gtfs.GetStopInfo(stopID, s)
		fmt.Println(si.String())
	}
	return nil
}
