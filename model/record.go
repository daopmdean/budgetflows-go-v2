package model

type RecordGet struct {
	Version string `json:"version"`

	From string `json:"from"`
	To   string `json:"to"`
}

type RecordReport struct {
	Version string `json:"version"`

	From string `json:"from"`
	To   string `json:"to"`
}

type RecordDelete struct {
	Version string `json:"version"`
	Id      string `json:"id"`
}
