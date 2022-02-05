package EncodeDecode

import (
	"bufio"
	AT "myArchHafm/lib/AuxiliaryTypes"
	"myArchHafm/lib/packTree"
	"os"

)

func Decode(file *os.File) string{
	ReaderBites := AT.NewReaderBites(bufio.NewReader(file))
	tree := packTree.NewTreeFromBites(ReaderBites)
	return tree.Decode(ReaderBites)
}

