package types

import "encoding/json"

type WeaponDetail struct {
	ID            json.Number `json:"id",string`
	ImagePath     string      `json:"image"`
	ThumbnailPath string      `json:"thumbnail"`
	EnglishName   string      `json:"name"`
}
