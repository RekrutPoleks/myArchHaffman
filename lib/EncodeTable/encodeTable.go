package EncodeTable

import (
	"fmt"
	AT "myArchHafm/lib/AuxiliaryTypes"
	"myArchHafm/lib/MyErr"
	"myArchHafm/lib/packTree"
	"strings"
)

type ETable map[rune]string

func (et ETable) EncodeTextToStringByte(text string) string {
	var buf strings.Builder
	for _, char := range text {
		if code, ok := et[char]; ok{
			buf.WriteString(code)
		}else{
			MyErr.HandleErr(fmt.Errorf("Can't encode text"))
		}
	}
	return buf.String()
}


func ConvertTreeToTable(tree *packTree.Tree) ETable {
	code := ""
	table := make(ETable)
	bildTableFromTree(code, tree, table)
	return table
}

func  bildTableFromTree(cd string, node *packTree.Tree, table ETable){

	if node.Char != AT.Char(0){

		table[rune(node.Char)]=cd
	} else {
		if node.One != nil{
			bildTableFromTree(cd+"1", node.One, table)
		}
		if node.Zero != nil{
			bildTableFromTree(cd+"0", node.Zero, table)
		}
	}
}