package models

type Unit struct {
	Document

	Agency string   `json:"agency,omitempty"`
	UnitID string   `json:"unitID,omitempty"`
	Notes  []string `json:"notes,omitempty"`
}

func NewUnitDocument() Document {
	return newDocument(DocTypeUnit, nil)
}
