package types

import "encoding/json"

type Weapon struct {
	ID         json.Number  `json:"id",string`
	Detail     WeaponDetail `json:"weapon"`
	CoopWeapon WeaponDetail `json:"coop_special_weapon"`
}

func (t Weapon) JapaneseName() (string, bool) {
	jpName, _, isFound := t.getWeaponInfo()
	if !isFound {
		return "", false
	}

	return jpName, true
}

func (t Weapon) StatinkKey() (string, bool) {
	_, key, isFound := t.getWeaponInfo()
	if (!isFound) || (key == "") {
		return "", false
	}

	return key, true
}

func (t Weapon) getWeaponInfo() (string, string, bool) {
	id, err := t.ID.Int64()
	if err != nil {
		return "", "", false
	}

	return getWeaponInfoFromID(id)
}

func getWeaponInfoFromID(id int64) (string, string, bool) {
	switch id {
	case -2:
		return "ランダム", "", true
	case -1:
		return "ランダム", "", true
	case 0:
		return "ボールドマーカー", "bold", true
	case 10:
		return "わかばシューター", "wakaba", true
	case 20:
		return "シャープマーカー", "sharp", true
	case 30:
		return "プロモデラーMG", "promodeler_mg", true
	case 40:
		return "スプラシューター", "sshooter", true
	case 50:
		return ".52ガロン", "52gal", true
	case 60:
		return "N-ZAP85", "nzap85", true
	case 70:
		return "プライムシューター", "prime", true
	case 80:
		return ".96ガロン", "96gal", true
	case 90:
		return "ジェットスイーパー", "jetsweeper", true
	case 200:
		return "ノヴァブラスター", "nova", true
	case 210:
		return "ホットブラスター", "hotblaster", true
	case 220:
		return "ロングブラスター", "longblaster", true
	case 230:
		return "クラッシュブラスター", "clashblaster", true
	case 240:
		return "ラピッドブラスター", "rapid", true
	case 250:
		return "Rブラスターエリート", "rapid_elite", true
	case 300:
		return "L3リールガン", "l3reelgun", true
	case 310:
		return "H3リールガン", "h3reelgun", true
	case 400:
		return "ボトルガイザー", "bottlegeyser", true
	case 1000:
		return "カーボンローラー", "carbon", true
	case 1010:
		return "スプラローラー", "splatroller", true
	case 1020:
		return "ダイナモローラー", "dynamo", true
	case 1030:
		return "ヴァリアブルローラー", "variableroller", true
	case 1100:
		return "パブロ", "pablo", true
	case 1110:
		return "ホクサイ", "hokusai", true
	case 2000:
		return "スクイックリンα", "squiclean_a", true
	case 2010:
		return "スプラチャージャー", "splatcharger", true
	case 2020:
		return "スプラスコープ", "splatscope", true
	case 2030:
		return "リッター4K", "liter4k", true
	case 2040:
		return "4Kスコープ", "liter4k_scope", true
	case 2050:
		return "14式竹筒銃・甲", "bamboo14mk1", true
	case 2060:
		return "ソイチューバー", "soytuber", true
	case 3000:
		return "バケットスロッシャー", "bucketslosher", true
	case 3010:
		return "ヒッセン", "hissen", true
	case 3020:
		return "スクリュースロッシャー", "screwslosher", true
	case 3030:
		return "オーバーフロッシャー", "furo", true
	case 3040:
		return "エクスプロッシャー", "explosher", true
	case 4000:
		return "スプラスピナー", "splatspinner", true
	case 4010:
		return "バレルスピナー", "barrelspinner", true
	case 4020:
		return "ハイドラント", "hydra", true
	case 4030:
		return "クーゲルシュライバー", "kugelschreiber", true
	case 4040:
		return "ノーチラス47", "nautilus47", true
	case 5000:
		return "スパッタリー", "sputtery", true
	case 5010:
		return "スプラマニューバー", "maneuver", true
	case 5020:
		return "ケルビン525", "kelvin525", true
	case 5030:
		return "デュアルスイーパー", "dualsweeper", true
	case 5040:
		return "クアッドホッパーブラック", "quadhopper_black", true
	case 6000:
		return "パラシェルター", "parashelter", true
	case 6010:
		return "キャンピングシェルター", "campingshelter", true
	case 6020:
		return "スパイガジェット", "spygadget", true
	case 20000:
		return "クマサン印のブラスター", "kuma_blaster", true
	case 20010:
		return "クマサン印のシェルター", "kuma_shelter", true
	case 20020:
		return "クマサン印のチャージャー", "kuma_charger", true
	case 20030:
		return "クマサン印のスロッシャー", "kuma_slosher", true
	}
	return "", "", false
}
