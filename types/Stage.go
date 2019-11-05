package types

import "strings"

type Stage struct {
	ImagePath   string `json:"image"`
	EnglishName string `json:"name"`
}

func (t Stage) JapaneseName() (string, bool) {
	switch {
	case strings.Contains(t.ImagePath, "50064ec6e97aac91e70df5fc2cfecf61ad8615fd"):
		return "朽ちた箱舟 ポラリス", true
	case strings.Contains(t.ImagePath, "65c68c6f0641cc5654434b78a6f10b0ad32ccdee"):
		return "シェケナダム", true
	case strings.Contains(t.ImagePath, "6d68f5baa75f3a94e5e9bfb89b82e7377e3ecd2c"):
		return "海上集落シャケト場", true
	case strings.Contains(t.ImagePath, "e07d73b7d9f0c64e552b34a2e6c29b8564c63388"):
		return "難破船ドン・ブラコ", true
	case strings.Contains(t.ImagePath, "e9f7c7b35e6d46778cd3cbc0d89bd7e1bc3be493"):
		return "トキシラズいぶし工房", true
	default:
		return "", false
	}
}
