package enum

type FilterByEnum string

const (
	IDFilterByEnum         = FilterByEnum("id")
	UsernameFilterByEnum   = FilterByEnum("username")
	NameFilterByEnum       = FilterByEnum("name")
	PlayerCodeFilterByEnum = FilterByEnum("player_code")
)
