package models

type NameMap struct {
	UUID          string `json:"uuid"`
	Src           string `json:"src"`
	Path          string `json:"path"`
	Size          int64  `json:"size"`
	Type          string `json:"type"`
	DownloadCount int64  `json:"download_count"`
}
