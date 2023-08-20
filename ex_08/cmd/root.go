package cmd

import (
	"log"

	"github.com/hoangtk0100/dc-go-23/ex_08/pkg/handler"
	"github.com/hoangtk0100/dc-go-23/ex_08/pkg/util"
	"github.com/spf13/cobra"
)

var rootCMD = &cobra.Command{
	Use:   "app",
	Short: "Start the app",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := util.LoadConfig(".env")
		if err != nil {
			log.Fatal("Cannot load config:", err)
		}

		server := handler.NewServer(config)
		server.RunDBMigration()
		server.Start()
	},
}

func Execute() {
	if err := rootCMD.Execute(); err != nil {
		log.Fatal(err)
	}
}
