package phone

import (
	"fmt"
	"regexp"
)

type PhoneNumber struct {
	RawText string
	Number  string
}

func (p *PhoneNumber) Normalize() {
	r := regexp.MustCompile(`[^0-9.]`)
	p.Number = r.ReplaceAllString(p.RawText, "")
}

func (p PhoneNumber) PrintNumber() {
	fmt.Println(p.Number)
}
