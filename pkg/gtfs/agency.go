package gtfs

import (
	"fmt"
)

type Agency struct {
	ID          string `json:"id,omitempty" csv:"agency_id"`
	Name        string `json:"name" csv:"agency_name"`
	URL         string `json:"url" csv:"agency_url"`
	Timezone    string `json:"timezone" csv:"agency_timezone"`
	Lang        string `json:"lang,omitempty" csv:"agency_lang"`
	Phone       string `json:"phone,omitempty" csv:"agency_phone"`
	FareURL     string `json:"fareUrl,omitempty" csv:"agency_fare_url"`
	AgencyEmail string `json:"email,omitempty" csv:"agency_email"`
}

func (a Agency) key() string {
	return a.ID
}

func (a Agency) validate() errorList {
	var errs errorList

	if a.Name == "" {
		_ = errs.add(fmt.Errorf("agency name is required"))
	}
	if a.URL == "" {
		_ = errs.add(fmt.Errorf("agency URL is required"))
	}
	if a.Timezone == "" {
		_ = errs.add(fmt.Errorf("agency timezone is required"))
	}

	return errs
}
