package regerrors

import "fmt"

type RegexError struct {
	Code    string
	Message string
	Pos     int
}

func (p *RegexError) Error() string {
	return fmt.Sprintf("code=%s, message=%s, pos=%d", p.Code, p.Message, p.Pos)
}
