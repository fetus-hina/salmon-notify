package types

import "time"

type CoopScheduleDetail struct {
	Stage     Stage    `json:"stage"`
	Weapons   []Weapon `json:"weapons"`
	StartTime int64    `json:"start_time"`
	EndTime   int64    `json:"end_time"`
}

func (t CoopScheduleDetail) IsOpened() bool {
	now := time.Now().Unix()
	return t.StartTime <= now && now < t.EndTime
}

func (t CoopScheduleDetail) IsOpenRecently() bool {
	if !t.IsOpened() {
		return false
	}

	now := time.Now().Unix()
	return t.StartTime >= now-30*60
}
