package models

type Vehicle struct {
	ID          string         `json:"id,omitempty"`
	Agency      string         `json:"agency,omitempty"`
	Orientation string         `json:"orientation,omitempty"`
	Sequence    map[string]int `json:"sequence,omitempty"`
	Length      int            `json:"length,omitempty"`
	Position    int            `json:"position,omitempty"`
}

var (
	Bus = Vehicle{
		ID:     "5738",
		Agency: "MUNI",
	}

	BARTLegacyRun = Vehicle{
		Agency: "BART",
		Sequence: map[string]int{
			"1259": 1,
			"1738": 2,
			"1607": 3,
			"1897": 4,
			"1212": 5,
		},
		Length: 5,
	}

	Plane = Vehicle{
		ID:     "EI-EIM",
		Agency: "Aer Lingus",
	}

	TuesdayCommute = Vehicle{
		ID:     "4496",
		Agency: "BART",
		Sequence: map[string]int{
			"4496": 3,
		},
		Length: 8,
	}

	ChurchToKing = Vehicle{
		ID:     "1501",
		Agency: "MUNI",
	}
)
