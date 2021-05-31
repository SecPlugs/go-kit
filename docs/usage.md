```
package main

import (
	"fmt"
	"github.com/secplugs/go-kit/filescan"
	"log"
	"os"
)

func main() {
	fmt.Println("Using secplugs.com filescan")
	client := filescan.NewDefaultScanClient()

	// We will use this go source file to scan
	result, err := client.ScanFile(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is the score", result.Score)
}

```
