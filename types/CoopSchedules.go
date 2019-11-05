package types

type CoopSchedules struct {
	Details   []CoopScheduleDetail `json:"details"`
	Schedules []CoopSchedule       `json:"schedules"`
}
