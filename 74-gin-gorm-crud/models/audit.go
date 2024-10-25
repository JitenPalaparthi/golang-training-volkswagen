package models

type Audit struct {
	Action       string `json:"action"`
	Data         string `json:"record"`
	User         string `json:"user"`
	LastModified int64  `json:"last_modified"`
}
