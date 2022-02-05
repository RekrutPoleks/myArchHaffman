package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"myArchHafm/lib/EncodeDecode"
	"myArchHafm/lib/MyErr"
	"os"
	"path/filepath"
	"strings"
)

var cmdUnpack = &cobra.Command{
			Use : "unpack",
			Short: "Unpack according to the hofmann algorithm",
			Run: unpack,
}

func init()  {
	rootCmd.AddCommand(cmdUnpack)
}

func unpack(_ *cobra.Command, args []string)  {
	if len(args[0]) == 0 || args[0]==""{
		MyErr.HandleErr(fmt.Errorf("path error"))
	}
	file, err := os.Open(args[0])
	if err != nil{
		MyErr.HandleErr(err)
	}
	defer file.Close()
	data :=EncodeDecode.Decode(file)
	err = os.WriteFile(fileUnpackNewName(args[0]), []byte(data), 0677)
}

func fileUnpackNewName(path string) string{
	bname := filepath.Base(path)
	return strings.TrimSuffix(path, filepath.Ext(bname)) + ".unpack"
}