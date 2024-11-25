package graph

type States struct {
	Start       bool
	End         bool // if this is true, the regex matches
	Transitions map[uint8][]*States
}
