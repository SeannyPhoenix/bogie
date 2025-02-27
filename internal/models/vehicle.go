package models

type VehicleUnit struct {
	Position    int    `json:"position,omitempty"`
	Orientation string `json:"orientation,omitempty"`
}

type Vehicle struct {
	ID       string                 `json:"id,omitempty"`
	Sequence map[string]VehicleUnit `json:"sequence,omitempty"`
	Length   int                    `json:"length,omitempty"`
}

var (
	MUNIBus = Vehicle{
		ID: "5738",
	}

	BARTLegacyRun = Vehicle{
		Sequence: map[string]VehicleUnit{
			"1259": {Position: 1},
			"1738": {Position: 2},
			"1607": {Position: 3},
			"1897": {Position: 4},
			"1212": {Position: 5},
		},
		Length: 5,
	}

	AerLingusPlane = Vehicle{
		ID: "EI-EIM",
	}

	TuesdayBARTCommute = Vehicle{
		ID: "4496",
		Sequence: map[string]VehicleUnit{
			"4496": {Position: 3},
		},
		Length: 8,
	}

	ChurchToKing = Vehicle{
		ID: "1501",
	}
)
