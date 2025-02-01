package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/seannyphoenix/bogie/pkg/gtfs"
	"github.com/seannyphoenix/bogie/pkg/util"
)

func main() {
	tt := util.TrackTime("create GTFS collection")

	handlers := parseFlags()

	defer tt()

	gtfsDir := "gtfs_files"

	if _, err := os.Stat(gtfsDir); err != nil {
		log.Fatalf("Error finding %s: %s \n", gtfsDir, err.Error())
	}

	zipFiles, err := filepath.Glob(filepath.Join(gtfsDir, "*.zip"))
	if err != nil {
		log.Fatalf("Malformed file path: %s\n", err.Error())
	}

	col, err := gtfs.CreateGTFSCollection(zipFiles)
	if err != nil {
		log.Fatalf("Error creating GTFS schedule collection: %s\n", err)
		tt()
	}

	for _, h := range handlers {
		err := h(col)
		if err != nil {
			log.Fatalf("Error running handler: %s\n", err)
		}
	}
}

func parseFlags() []handler {
	var hh []handler

	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			switch arg {
			case "-e":
				hh = append(hh, writeErrors)
			case "-o":
				hh = append(hh, printScheduleOverviews)
			default:
			}
		}
	}

	return hh
}
