package packTree

import (
	"container/heap"
	"encoding/binary"
	"fmt"
	"io"
	AT "myArchHafm/lib/AuxiliaryTypes"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)




type Tree struct {
	Char AT.Char
	One *Tree
	Zero *Tree
}


func (tr *Tree) ConvertToStringBite() string {
	// Произвести обход дерева в ширину в результате обхода должна вернуться строка вида
	//  два бита наличие узлов 0 or 1 - 0: узел Zero 1: узел One, 0 or 1 1: Tree.Char != nil
	//если предыдущий бит 1 => 2 бита размер символа в байтах 00 :1 байт, 01 2 байт, 10 3 байта, 11 4 байта.
	// провести обход в ширину вернуть строк
	var buf strings.Builder
	strTree := widthTraversal(tr)
	buf.WriteString(fmt.Sprintf("%016b", len(strTree)))
	buf.WriteString(strTree)
	return buf.String()

}

func widthTraversal(tree *Tree) string {
	var buf strings.Builder
	travSlice := &([]*Tree{tree})
	for i:=0; i != len(*travSlice);i++{
		node := (*travSlice)[i]
		if node.One != nil{
			buf.WriteString("1")
			tmp := append(*travSlice, node.One)
			travSlice = &tmp
		}else{
			buf.WriteString("0")
		}
		if node.Zero != nil{
			buf.WriteString("1")
			tmp := append(*travSlice, node.Zero)
			travSlice = &tmp
		} else {
			buf.WriteString("0")
		}
		if rune(node.Char) != rune(0){
			buf.WriteString("1")
			buf.WriteString(EncodeLenChar(node.Char.Len()))
			buf.WriteString(node.Char.ToStringByte())
		}else {
			buf.WriteString("0")
		}
	}
	return buf.String()
}
func EncodeLenChar(len int) string{
	var res string
	switch len{
	case 1:
		res = "00"
	case 2:
		res = "01"
	case 3:
		res = "10"
	case 4:
		res = "11"
	}
	return res
}


func BinaryTreeEncodeFromText(text string) *Tree {
	count := AT.NewCounter()
	count.Calculate(text)
	pq := make(PriorityQueue, len(count))
	ind := 0
	for char, pr := range count {
		c := AT.Char(char)
		pq[ind]= &Item{
			Value: &Tree{
				c,
				nil,
				nil,
			},
			Priority: pr,
			Index:    ind,
		}
		ind++
	}
	heap.Init(&pq)
	for pq.Len() > 1 {
		one, zero := heap.Pop(&pq).(*Item), heap.Pop(&pq).(*Item)
		item := &Item{
			Value:&Tree{
				AT.Char(0),
				one.Value,
				zero.Value,
			},
			Priority: one.Priority + zero.Priority,
		}
		heap.Push(&pq, item)
	}
	return heap.Pop(&pq).(*Item).Value
}




func NewTreeFromBites(reader *AT.ReaderBites) *Tree{

	tree := bildTreeForBites(reader)
	return tree
}

func decodeLenHeadTree(bites []AT.Bite) int{
	var buf strings.Builder
	for _, bite := range bites {
		buf.WriteString(bite.ToString())
	}
	res, _ := strconv.ParseUint(buf.String(), 2, len(bites))
	return int(res)
}

func bildTreeForBites(reader *AT.ReaderBites) *Tree{
	tree := &Tree{}
	blenTree, _ := reader.Read(AT.SizeTreeInBites)
	lenTree := decodeLenHeadTree(blenTree)
	sliceBild := []*Tree{tree}
	for l,i :=0, 0;l < lenTree; i,l = i+1, l+AT.LenCodeNode {
		node := sliceBild[i]
		codeNode, _ := reader.Read(AT.LenCodeNode)
		if codeNode[0] == AT.Bite1{
			node.One = &Tree{AT.Char(rune(0)), nil, nil}
			sliceBild = append(sliceBild, node.One)

		}
		if codeNode[1] == AT.Bite1{
			node.Zero = &Tree{AT.Char(rune(0)), nil, nil}
			sliceBild = append(sliceBild, node.Zero)

		}
		if codeNode[2] == AT.Bite1{
			//считать длину из 2 бит 00 : 1байт, 01 ,10, 11
			bites, _ := reader.Read(2)
			//присвоить переменой char Char
			lenChar := convertCodeBiteInLenRune(bites)
			sbChar, _ := reader.ReadToString(lenChar)
			node.Char = convertBiteStringInChar(sbChar)
			l+=lenChar+2

		}

	}

	return tree

}

func convertCodeBiteInLenRune(bite []AT.Bite)int {
	if bite[0] == AT.Bite0{
		if bite[1] == AT.Bite0 {
			return 8
		} else {
			return 16
		}
	}else{
		if bite[1] == AT.Bite0 {
			return 24
		} else {
			return 32
		}
	}
}

func convertBiteStringInChar(bites string) AT.Char{
	bt, _  :=  strconv.ParseUint(bites, 2, len(bites))
	bChar := make([]byte, 8)
	binary.BigEndian.PutUint64(bChar, bt)
	r,_ := utf8.DecodeRune(bChar[8-len(bites)/8:])
	return AT.Char(r)
}

func (tr *Tree) Decode(bites *AT.ReaderBites) string {
	node := tr
	var buf strings.Builder
	flagCapital := false
	for bite, err := bites.Read(1); err != io.EOF; bite, err = bites.Read(1) {
		//fmt.Printf("node:%p one:%p\tzero: %p\n", node, node.One, node.Zero)
		switch bite[0]{
		case  AT.Bite0:
			node = node.Zero
		case AT.Bite1:
			node = node.One
		}
		if node.Char != AT.Char(0) {
			if node.Char == AT.Char(AT.FlagIsCapital){
				flagCapital = true
				node = tr
				continue
			}
			r := rune(node.Char)
			if flagCapital{
				r = unicode.ToUpper(r)
				flagCapital = false
			}
			buf.WriteRune(r)
			node = tr
		}
	}

	return buf.String()
}
