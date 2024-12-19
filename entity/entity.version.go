package entity

import "time"

type VersionAdmin struct {
	VersionId       int       `json:"version_id"`
	VersionNumber   string    `json:"version_number"`
	VersionCode     int       `json:"version_code"`
	VersionChagelog string    `json:"version_chagelog"`
	VersionDatetime time.Time `json:"version_datetime"`
}

type VersionShop struct {
	VersionId       int       `json:"version_id"`
	VersionNumber   string    `json:"version_number"`
	VersionCode     int       `json:"version_code"`
	VersionChagelog string    `json:"version_chagelog"`
	VersionDatetime time.Time `json:"version_datetime"`
}
