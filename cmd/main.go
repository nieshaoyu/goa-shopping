package main

import (
	"log"

	"github.com/spf13/cobra"

	shoppingCmd "goa-shopping/cmd/shopping"
	"goa-shopping/config"
	"goa-shopping/dao"
)

var (
	cfgFile string
)

func serverCmd() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "runserver",
		Short: "run server",
		Long:  "run http  server",
		RunE: func(cmd *cobra.Command, args []string) error {
			shoppingCmd.RunServer(config.C)
			return nil
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			config.Init(cfgFile)
			dao.InitDB(config.C)
			dao.AutoMigrateDB()

			return nil
		},
	}

	serverCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file")
	_ = serverCmd.MarkPersistentFlagRequired("config")

	return serverCmd
}

var RootCmd = cobra.Command{
	Use: "shopping",
}

func main() {
	RootCmd.AddCommand(serverCmd())

	if err := RootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
