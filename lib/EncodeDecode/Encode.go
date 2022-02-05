package EncodeDecode

import (
	"bytes"
	AT "myArchHafm/lib/AuxiliaryTypes"
	"myArchHafm/lib/EncodeTable"
	"myArchHafm/lib/packTree"
	"strings"
	"unicode"
)



func Encode(text string) []byte {
	var buf strings.Builder
	text = prepare(text)
	tree := packTree.BinaryTreeEncodeFromText(text)
	buf.WriteString(tree.ConvertToStringBite())
	table := EncodeTable.ConvertTreeToTable(tree)
	buf.WriteString(table.EncodeTextToStringByte(text))
	return splitStringToByte(buf.String())
}

func prepare(text string) string {
	var buf strings.Builder
	for _, char := range text {
		if unicode.IsUpper(char) {
			buf.WriteRune(AT.FlagIsCapital)
			char = unicode.ToLower(char)
		}
		buf.WriteRune(char)
	}
	return buf.String()
}

func splitStringToByte(str string)[]byte{
	var bufB bytes.Buffer
	var bufS strings.Builder
	for i, s:= range str {
		bufS.WriteRune(s)
		if (i+1)%AT.SizeChunks == 0{
			bufB.WriteByte(AT.Chunk(bufS.String()).Byte())
			bufS.Reset()
		}
	}
	if bufS.Len() !=0{
		str = bufS.String()+strings.Repeat("0", AT.SizeChunks-bufS.Len())
		bufB.WriteByte(AT.Chunk(bufS.String()).Byte())
	}

	return bufB.Bytes()
}
