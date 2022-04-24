---
layout: none
---

{brand-name} powered, ready to use tooling written in Go.

The plugin is open source so you can modify as you wish.

## Installation
Get the go module by simply running the commands below.
```console
GO111MODULE=on go get github.com/secplugs/go-kit/filescan
```
You'll now have the module in your `$GOPATH`

## Usage
Usage pattern is instanciate a client and then use its methods to scan items.

### Scan A File
Here, a very simple example of how to scan a file
```go
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
### Use Your Own API Key
To use additional features and the privacy of your own account, after registering with {brand-name}, sign in with your username and [create an API key](docs?doc=docs/HowTo/CreateKey) 

After creating a key, the only change to the code sample above would be

```
client := filescan.NewScanClient{"my-api-key"}
```

Everything else remains the same.

## Contact
Having trouble? [Contact {brand-name} ](https://{brand-root-domain}/contacts)
