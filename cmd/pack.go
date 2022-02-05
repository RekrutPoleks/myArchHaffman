package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"myArchHafm/lib/AuxiliaryTypes"
	"myArchHafm/lib/EncodeDecode"
	"myArchHafm/lib/MyErr"
	"os"
	"path/filepath"
	"strings"
)



var packCmd = &cobra.Command{
	Use : "pack",
	Short: "Pack according to the hofmann algorithm",
	Run: pack,
}

func init(){
	rootCmd.AddCommand(packCmd)
}

func pack(_ *cobra.Command, args[]string) {
	if len(args[0]) == 0 || args[0]==""{
		MyErr.HandleErr(fmt.Errorf("path error"))
	}
	file, err := os.Open(args[0])
	if err != nil{
		MyErr.HandleErr(err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil{
		MyErr.HandleErr(err)
	}
	encodeText := EncodeDecode.Encode(string(data))

	err = os.WriteFile(fileNewName(args[0]), encodeText, 0644)
	if err != nil{
		MyErr.HandleErr(err)
	}
}

func fileNewName(path string) string{
	bname := filepath.Base(path)
	return strings.TrimSuffix(path, filepath.Ext(bname)) + AuxiliaryTypes.PackExtension
}