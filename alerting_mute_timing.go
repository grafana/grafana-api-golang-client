package gapi

// MuteTiming represents a Grafana Alerting mute timing.
type MuteTiming struct {
	Name          string         `json:"name"`
	TimeIntervals []TimeInterval `json:"time_intervals"`
	Provenance    string         `json:"provenance,omitempty"`
}

// TimeInterval describes intervals of time using a Prometheus-defined standard.
type TimeInterval struct {
	Times       []TimeRange       `json:"times,omitempty"`
	Weekdays    []WeekdayRange    `json:"weekdays,omitempty"`
	DaysOfMonth []DayOfMonthRange `json:"days_of_month,omitempty"`
	Months      []MonthRange      `json:"months,omitempty"`
	Years       []YearRange       `json:"years,omitempty"`
}

// TimeRange represents a range of minutes within a 1440 minute day, exclusive of the End minute.
type TimeRange struct {
	StartMinute int
	EndMinute   int
}

// A WeekdayRange is an inclusive range between [0, 6] where 0 = Sunday.
type WeekdayRange struct {
	InclusiveRange
}

// A DayOfMonthRange is an inclusive range that may have negative Beginning/End values that represent distance from the End of the month Beginning at -1.
type DayOfMonthRange struct {
	InclusiveRange
}

// A MonthRange is an inclusive range between [1, 12] where 1 = January.
type MonthRange struct {
	InclusiveRange
}

// A YearRange is a positive inclusive range.
type YearRange struct {
	InclusiveRange
}

// InclusiveRange is used to hold the Beginning and End values of many time interval components.
type InclusiveRange struct {
	Begin int
	End   int
}

func (c *Client) MuteTimings() ([]MuteTiming, error) {
	mts := make([]MuteTiming, 0)
	err := c.request("GET", "/api/v1/provisioning/mute-timings", nil, nil, &mts)
	if err != nil {
		return nil, err
	}
	return mts, nil
}
