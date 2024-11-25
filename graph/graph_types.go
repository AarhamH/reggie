package graph

type States struct {
	start       bool
	end         bool // if this is true, the regex matches
	transitions map[uint8][]*States
}
