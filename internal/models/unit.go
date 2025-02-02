package models

type Unit struct {
	Record

	Agency string   `json:"agency,omitempty"`
	UnitID string   `json:"unitID,omitempty"`
	Notes  []string `json:"notes,omitempty"`
}

func NewUnitRecord() Record {
	return newRecord(DocTypeUnit, nil)
}
