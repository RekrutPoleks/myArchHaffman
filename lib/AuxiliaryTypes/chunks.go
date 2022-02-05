package AuxiliaryTypes

import (
	"fmt"
	"myArchHafm/lib/MyErr"
	"strconv"
)


type Chunk string

func (c Chunk) Byte() byte {
	num, err := strconv.ParseUint(string(c),2, SizeChunks)
	if err != nil{
		MyErr.HandleErr(fmt.Errorf("Can't covert chunk %v", c))
	}
	return byte(num)
}

