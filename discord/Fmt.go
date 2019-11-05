package discord

import (
	"fmt"
	"github.com/fetus-hina/salmon-notify/types"
	"strings"
	"time"
)

func FormatMessage(schedule types.CoopScheduleDetail) string {
	return fmt.Sprintf("%s %s\n%s",
		formatTermRange(schedule.StartTime, schedule.EndTime),
		formatStage(schedule.Stage),
		formatWeapons(schedule.Weapons),
	)
}

func formatTermRange(startTime, endTime int64) string {
	return fmt.Sprintf("%s～%s",
		formatDateTime(startTime),
		formatDateTime(endTime),
	)
}

func formatDateTime(timestamp int64) string {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	t := time.Unix(timestamp, 0).In(jst)
	return t.Format("2006/01/02 15:04")
}

func formatStage(stage types.Stage) string {
	jpName, isFound := stage.JapaneseName()
	if !isFound {
		return stage.EnglishName
	}

	return jpName
}

func formatWeapons(weapons []types.Weapon) string {
	names := make([]string, 0, 4)
	for _, weapon := range weapons {
		names = append(names, formatWeapon(weapon))
	}
	return strings.Join(names, "\n")
}

func formatWeapon(weapon types.Weapon) string {
	return strings.TrimLeft(
		fmt.Sprintf(
			"%s %s",
			getWeaponEmoji(weapon),
			getWeaponName(weapon),
		),
		" ",
	)
}

func getWeaponEmoji(weapon types.Weapon) string {
	// 絵文字をサポートするにはあらかじめ絵文字のリストを取得して、
	// 特殊な形式で送信しなければならないらしい。
	// https://github.com/Rapptz/discord.py/issues/390

	return ""

	// key, isFound := weapon.StatinkKey()
	// if !isFound {
	// 	return ""
	// }

	// return ":" + key + ":"
}

func getWeaponName(weapon types.Weapon) string {
	name, isFound := weapon.JapaneseName()
	if !isFound {
		return fmt.Sprintf("(Unknwon #%s)", weapon.ID)
	}

	return name
}
