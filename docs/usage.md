```
package main

import (
        "fmt"
        "log"

        "github.com/secplugs/go-kit/filescan"
)

func main() {
	fmt.Println("Using secplugs.com filescan")
	client := filescan.NewDefaultScanClient()

	result, err := client.ScanFile("/path/to/file/to/be/scanned")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d is the score", result.Score)
}
```
