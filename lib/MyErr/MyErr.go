package MyErr

import (
	"fmt"
	"os"
)

func HandleErr(err error){
	_, _ = fmt.Fprintln(os.Stdout, err)
	os.Exit(1)
}