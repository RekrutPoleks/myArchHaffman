package AuxiliaryTypes

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Counter map[rune]int

func NewCounter() Counter{
	return make(Counter)
}

func (c *Counter)Calculate(text string){
	for _, char := range text {
		(*c)[char]+=1
	}
}

func (c *Counter) reset()  {
	*c = make(Counter)
}

type Char rune

func (c Char) ToByte() []byte {
	res := make([]byte, utf8.RuneLen(rune(c)))
	utf8.EncodeRune(res, rune(c))
	return res
}

func (c Char) ToStringByte() string {
	var buf strings.Builder
	for _, b := range c.ToByte() {
		buf.WriteString(fmt.Sprintf("%08b", b))
	}

	return buf.String()
}

func (c Char) Len() int {
	return utf8.RuneLen(rune(c))
}

