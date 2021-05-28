package filescan_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"secplugs.com/filescan"
)

func TestNewDefaultScanClient(t *testing.T) {
	client := filescan.NewDefaultScanClient()
	if client.ApiKey != "r2iKI4q7Lu91Nu5uPl3eW3BPmRo4XK1ZbhLWtOKd" {
		t.Errorf("ApiKey is incorrect: %s", client.ApiKey)
	}
}

func TestNewScanClient(t *testing.T) {
	client := filescan.NewScanClient("testapikey")
	if client.ApiKey != "testapikey" {
		t.Errorf("ApiKey is incorrect: %s", client.ApiKey)
	}
}

func ExampleScanFile() {
	client := filescan.NewDefaultScanClient()
	filename, err := filepath.Abs(os.Args[0])
	result, err := client.ScanFile(filename)
	if err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
	fmt.Printf("Score: %d\n", result.Score)
	// Output:
	// Score: 70
}
