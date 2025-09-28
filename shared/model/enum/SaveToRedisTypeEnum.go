package enum

type SaveToRedisTypeEnum string

const (
	SuccessEarlyEnum        = SaveToRedisTypeEnum("success_early")
	SuccessOntimeEnum       = SaveToRedisTypeEnum("success_ontime")
	SuccessLateEnum         = SaveToRedisTypeEnum("success_late")
	FirstMileRejectEnum     = SaveToRedisTypeEnum("firstmile_reject")
	LastMileFailedEnum      = SaveToRedisTypeEnum("lastmile_failed")
	OnProcessOntimeEnum     = SaveToRedisTypeEnum("on_process_ontime")
	OnProcessLateEnum       = SaveToRedisTypeEnum("on_process_late")
	LastMileUndeliveredEnum = SaveToRedisTypeEnum("lastmile_undelivered")
	FirstMileCancelEnum     = SaveToRedisTypeEnum("firstmile_cancel")
)
