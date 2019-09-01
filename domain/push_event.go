package domain

type PushEvent struct {
	BeforeCommitID string `json:"before"`
	AfterCommitID  string `json:"after"`
}
