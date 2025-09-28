package enum

type ServiceTypeEnum string

const (
	PaxelRecycleEnum     = ServiceTypeEnum("PAXEL RECYCLE")
	InstantMultidropEnum = ServiceTypeEnum("INSTANT MULTIDROP")
	InstantEnum = ServiceTypeEnum("INSTANT GOSEND")
)
