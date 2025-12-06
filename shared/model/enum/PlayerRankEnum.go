package enum

type PlayerRankEnum string

const (
	LowRankPlayerRankEnum      = PlayerRankEnum("S")
	MidRankPlayerRankEnum      = PlayerRankEnum("SS")
	HighRankPlayerRankEnum     = PlayerRankEnum("SSS")
	SuperiorPlayerRankEnum     = PlayerRankEnum("superior")
	IntermediatePlayerRankEnum = PlayerRankEnum("intermediate")
	NormalPlayerRankEnum       = PlayerRankEnum("normal")
)
