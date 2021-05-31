package filescan_test

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/secplugs/go-kit/filescan"
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

func TestEicarTest(t *testing.T) {
	client := filescan.NewDefaultScanClient()
	mypath, _ := filepath.Abs(os.Args[0])
	dirname := filepath.Dir(mypath)
	// B64 encoded eicar
	encoded_data := "WDVPIVAlQEFQWzRcUFpYNTQoUF4pN0NDKTd9JEVJQ0FSLVNUQU5EQVJELUFOVElWSVJVUy1URVNULUZJTEUhJEgrSCo="
	data, _ := base64.URLEncoding.DecodeString(encoded_data)
	filename := filepath.Join(dirname, "test.bin")
	ioutil.WriteFile(filename, []byte(data), 0644)
	res, err := client.IsClean(filename)
	if err != nil {
		t.Fatal(err)
	}
	if res != false {
		t.Errorf("%s is not clean, but result says %v\n", data, res)
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

func ExampleIsClean() {
	client := filescan.NewDefaultScanClient()
	filename, _ := filepath.Abs(os.Args[0])
	res, _ := client.IsClean(filename)
	fmt.Printf("Result: %v\n", res)
	// Output:
	// Result: true
}
