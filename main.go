package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

var (
	Server string
	SPort  string
	Host   string
	HPort  string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	startCmd := StartCmd()
	// startCmd.Flags().StringVarP(&Server, "server", "s", "", "server-IP for make connection")
	// startCmd.Flags().StringVarP(&SPort, "port", "p", "", "server-Port for make connection")

	// if err := startCmd.MarkFlagRequired("server"); err != nil {
	// 	fmt.Println(err)
	// }
	// if err := startCmd.MarkFlagRequired("port"); err != nil {
	// 	fmt.Println(err)
	// }

	rootCmd.AddCommand(startCmd)
}

func main() {
	os.Exit(Run())

}

func StartCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "start Client Server ",
		Run: func(cmd *cobra.Command, args []string) {

			DirWatchStart()

		},
	}
}
