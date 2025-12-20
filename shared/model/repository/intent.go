package repository

type Intent struct {
	TeamCount      int
	PlayersPerTeam int
	Composition    map[string]int // roleCode -> jumlah
}
