package cmd

import (
	"github.com/spf13/cobra"
	"myArchHafm/lib/MyErr"
)

var rootCmd = &cobra.Command{
	Short: "Arhivator for hofman",

}

func Execute(){
	if err := rootCmd.Execute(); err !=nil {
		MyErr.HandleErr(err)
	}
}



