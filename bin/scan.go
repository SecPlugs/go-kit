package main

import (
	"os"

	"github.com/secplugs/go-kit/filescan"
)

func main() {
	client := filescan.NewDefaultScanClient()
	res, _ := client.ScanFile(os.Args[0])
	if res.Score < 70 {
		os.Exit(1)
	}
}
