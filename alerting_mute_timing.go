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

// A WeekdayRange is an inclusive range of weekdays, e.g. "monday" or "tuesday:thursday".
type WeekdayRange string

// A DayOfMonthRange is an inclusive range of days, 1-31, within a month, e.g. "1" or "14:16". Negative values can be used to represent days counting from the end of a month, e.g. "-1".
type DayOfMonthRange string

// A MonthRange is an inclusive range of months, either numerical or full calendar month, e.g "1:3", "december", or "may:august".
type MonthRange string

// A YearRange is a positive inclusive range of years, e.g. "2030" or "2021:2022".
type YearRange string

func (c *Client) MuteTimings() ([]MuteTiming, error) {
	mts := make([]MuteTiming, 0)
	err := c.request("GET", "/api/v1/provisioning/mute-timings", nil, nil, &mts)
	if err != nil {
		return nil, err
	}
	return mts, nil
}
