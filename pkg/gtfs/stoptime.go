package gtfs

import (
	"fmt"
)

type StopTime struct {
	TripID                   string   `json:"tripId" csv:"trip_id"`
	ArrivalTime              Time     `json:"arrivalTime,omitempty" csv:"arrival_time"`
	DepartureTime            Time     `json:"departureTime,omitempty" csv:"departure_time"`
	StopID                   string   `json:"stopId,omitempty" csv:"stop_id"`
	LocationGroupID          string   `json:"locationGroupId,omitempty" csv:"location_group_id"`
	LocationID               string   `json:"locationId,omitempty" csv:"location_id"`
	StopSequence             int      `json:"stopSequence" csv:"stop_sequence"`
	StopHeadsign             string   `json:"stopHeadsign,omitempty" csv:"stop_headsign"`
	StartPickupDropOffWindow Time     `json:"startPickupDropOffWindow,omitempty" csv:"start_pickup_drop_off_window"`
	EndPickupDropOffWindow   Time     `json:"endPickupDropOffWindow,omitempty" csv:"end_pickup_drop_off_window"`
	PickupType               *int     `json:"pickupType,omitempty" csv:"pickup_type"`
	DropOffType              *int     `json:"dropOffType,omitempty" csv:"drop_off_type"`
	ContinuousPickup         *int     `json:"continuousPickup,omitempty" csv:"continuous_pickup"`
	ContinuousDropOff        *int     `json:"continuousDropOff,omitempty" csv:"continuous_drop_off"`
	ShapeDistTraveled        *float64 `json:"shapeDistTraveled,omitempty" csv:"shape_dist_traveled"`
	Timepoint                *int     `json:"timepoint,omitempty" csv:"timepoint"`
	PickupBookingRuleId      string   `json:"pickupBookingRuleId,omitempty" csv:"pickup_booking_rule_id"`
	DropOffBookingRuleId     string   `json:"dropOffBookingRuleId,omitempty" csv:"drop_off_booking_rule_id"`
}

func (st StopTime) key() string {
	return fmt.Sprintf("%s-%d", st.TripID, st.StopSequence)
}

func (st StopTime) validate() errorList {
	var errs errorList

	if st.TripID == "" {
		errs.add(fmt.Errorf("trip ID is required"))
	}
	if st.StopSequence < 0 {
		errs.add(fmt.Errorf("stop sequence must be greater than or equal to 0"))
	}
	if st.StopID == "" {
		errs.add(fmt.Errorf("stop ID is required"))
	}

	return errs
}
