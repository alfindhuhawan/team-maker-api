package enum

type StatusProcessEnum string

const (
	// USED FOR INTERACTOR, DONT USE IT AS LATEST STATUS
	SuccessStatusProcessEnum   = StatusProcessEnum("success")
	OnProcessStatusProcessEnum = StatusProcessEnum("onprocess")
)
