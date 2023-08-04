package models

import (
	"time"
)

type NameMap struct {
	Id            int       `json:"id"`
	UUID          string    `json:"uuid"`
	Src           string    `json:"src"`
	Size          int64     `json:"size"`
	Type          string    `json:"type"`
	DownloadCount int       `json:"download_count"`
	CreatedAt     time.Time `json:"created_at"`
	Temporary     bool      `json:"temporary"`
	ExpireCount   int       `json:"expire_count"`
}
