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
	key, isFound := weapon.StatinkKey()
	if !isFound {
		return ""
	}

	// FIXME: shimiclean サーバー専用データを分離
	switch key {
	case "52gal":
		return "<:52gal:639796764916056064>"
	case "96gal":
		return "<:96gal:639796988942221317>"
	case "bamboo14mk1":
		return "<:bamboo14mk1:639797026757804042>"
	case "barrelspinner":
		return "<:barrelspinner:639797048635424788>"
	case "bold":
		return "<:bold:639797078687612950>"
	case "bottlegeyser":
		return "<:bottlegeyser:639797128344109056>"
	case "bucketslosher":
		return "<:bucketslosher:639797156252876821>"
	case "campingshelter":
		return "<:campingshelter:639797178696466442>"
	case "carbon":
		return "<:carbon:639797205233958924>"
	case "clashblaster":
		return "<:clashblaster:639797233847369729>"
	case "dualsweeper":
		return "<:dualsweeper:639797345789411338>"
	case "dynamo":
		return "<:dynamo:639797372721037342>"
	case "explosher":
		return "<:explosher:639797393889427456>"
	case "furo":
		return "<:furo:639797413619433483>"
	case "h3reelgun":
		return "<:h3reelgun:639797444502355970>"
	case "hissen":
		return "<:hissen:639797542464258048>"
	case "hokusai":
		return "<:hokusai:639797567491932172>"
	case "hotblaster":
		return "<:hotblaster:639797594842988555>"
	case "hydra":
		return "<:hydra:639797620684095488>"
	case "jetsweeper":
		return "<:jetsweeper:639797669518376960>"
	case "kelvin525":
		return "<:kelvin525:639797698039513089>"
	case "kugelschreiber":
		return "<:kugelschreiber:639797759456837632>"
	case "l3reelgun":
		return "<:l3reelgun:639797797578866718>"
	case "liter4k":
		return "<:liter4k:639797863785693234>"
	case "liter4k_scope":
		return "<:liter4k_scope:639797925102223390>"
	case "longblaster":
		return "<:longblaster:639797959751630859>"
	case "maneuver":
		return "<:maneuver:639798030798815288>"
	case "nautilus47":
		return "<:nautilus47:639798061610172425>"
	case "nova":
		return "<:nova:639798094673739788>"
	case "nzap85":
		return "<:nzap85:639798151342981120>"
	case "pablo":
		return "<:pablo:639798184985755698>"
	case "parashelter":
		return "<:parashelter:639798214865715200>"
	case "prime":
		return "<:prime:639798256473341952>"
	case "promodeler_mg":
		return "<:promodeler_mg:639798308000497705>"
	case "quadhopper_black":
		return "<:quadhopper_black:639798377508372490>"
	case "rapid":
		return "<:rapid:639798423003856897>"
	case "rapid_elite":
		return "<:rapid_elite:639798462086512651>"
	case "screwslosher":
		return "<:screwslosher:639798487239753732>"
	case "sharp":
		return "<:sharp:639798520190205953>"
	case "soytuber":
		return "<:soytuber:639798548598095874>"
	case "splatcharger":
		return "<:splatcharger:639798589065003008>"
	case "splatroller":
		return "<:splatroller:639798616353013760>"
	case "splatscope":
		return "<:splatscope:639798647999168512>"
	case "splatspinner":
		return "<:splatspinner:639798699282923520>"
	case "sputtery":
		return "<:sputtery:639798732958728192>"
	case "spygadget":
		return "<:spygadget:639798758967869441>"
	case "squiclean_a":
		return "<:squiclean_a:639798785131806730>"
	case "sshooter":
		return "<:sshooter:639798811828682772>"
	case "variableroller":
		return "<:variableroller:639798835387826196>"
	case "wakaba":
		return "<:wakaba:639798855931789347>"
	default:
		return ""
	}
}

func getWeaponName(weapon types.Weapon) string {
	name, isFound := weapon.JapaneseName()
	if !isFound {
		return fmt.Sprintf("(Unknwon #%s)", weapon.ID)
	}

	return name
}
