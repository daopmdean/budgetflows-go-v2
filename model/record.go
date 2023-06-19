package model

type RecordGet struct {
	Version string `json:"version"`
}

type RecordDelete struct {
	Version string `json:"version"`
	Id      string `json:"id"`
}
