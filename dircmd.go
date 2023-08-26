package main

import (
	"log"

	"github.com/spf13/cobra"
)

var FileId string

func init() {
	downloadCmd := downloadCmd()
	downloadCmd.Flags().StringVarP(&FileId, "fileid", "f", "", "fileid for download file")

	if err := downloadCmd.MarkFlagRequired("fileid"); err != nil {
		log.Println(err)
	}
	rootCmd.AddCommand(downloadCmd)

}

func downloadCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "download",
		Short: "download file from server",
		Run: func(cmd *cobra.Command, args []string) { // Download 폴더에 파일 다운로드
			//TODO Download 요청
		},
	}
}
