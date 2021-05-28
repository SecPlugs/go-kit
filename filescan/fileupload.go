// Package filescan provides the basic methods and functions required to
// scan a file at secplugs.com.
package filescan

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// UploadInfo represents the Object form of the JSON response returned by the
// /file/upload API meant to upload a file for malware scan.
type UploadInfo struct {
	UploadPost struct {
		URL    string `json:"url"`
		Fields struct {
			Key               string `json:"key"`
			XAmzAlgorithm     string `json:"x-amz-algorithm"`
			XAmzCredential    string `json:"x-amz-credential"`
			XAmzDate          string `json:"x-amz-date"`
			XAmzSecurityToken string `json:"x-amz-security-token"`
			Policy            string `json:"policy"`
			XAmzSignature     string `json:"x-amz-signature"`
		} `json:"fields"`
	} `json:"upload_post"`
}

// NewScanClient creates a new ScanClient with the specified API key.
// The API key can be created and obtained from the Secplugs portal.
func NewScanClient(api string) *ScanClient {
	client := &ScanClient{}
	client.ApiKey = api
	return client
}

// NewDefaultScanClient creates a new ScanClient with the default API key.
// No choice of vendors and limited features available with this API key.
func NewDefaultScanClient() *ScanClient {
	client := &ScanClient{}
	client.ApiKey = SECPLUGS_DEFAULT_API_KEY
	return client
}

// ScanFile is the high-level to submit a file to Secplugs for a quickscan.
// It generates a pre-signed URL, uploads the file to that URL and the triggers
// a quick scan.
func (client *ScanClient) ScanFile(filename string) (ScanResult, error) {
	client.FileName = filename
	cksum, err := client.CalculateChecksum(filename)
	if err != nil {
		return ScanResult{}, err
	}

	client.UniqueId = cksum
	info, err := client.GetPresignedUrl(cksum)
	if err != nil {
		return ScanResult{}, err
	}

	fields, err := info.GetFields()
	if err != nil {
		return ScanResult{}, err
	}
	url := info.GetUrl()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for k, v := range fields {
		fw, _ := writer.CreateFormField(k)
		fw.Write([]byte(v))
	}
	file, _ := os.Open(filename)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	httpclient := &http.Client{}
	writer.Close()
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return ScanResult{}, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())
	resp, err := httpclient.Do(req)
	if err != nil {
		return ScanResult{}, err
	}

	if resp.StatusCode >= 300 {
		return ScanResult{}, &HTTPResponseStatusError{resp.StatusCode, "Error"}
	} else {
		result, err := client.QuickScan()
		if err != nil {
			return ScanResult{}, err
		}
		return result, nil
	}
}

// CalculateChecksum generates the SHA256 sum for a file. It is meant for internal consumption by
// the ScanFile API. But this can also used by clients using the other low-level APIs.
func (client *ScanClient) CalculateChecksum(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", &ChecksumGenerationError{}
	}
	defer f.Close()

	sha256sum := sha256.New()
	if _, err := io.Copy(sha256sum, f); err != nil {
		return "", &ChecksumGenerationError{}
	}
	return fmt.Sprintf("%x", sha256sum.Sum(nil)), nil
}

// GetPresignedUrl is used to obtain an upload URL and some unique fields to subsequenly upload
// a file for analysis. This is a low-level API and can be used by clients that do not want to
// use the high-level ScanFile API.
func (client *ScanClient) GetPresignedUrl(cksum string) (UploadInfo, error) {
	secplugsBaseUrl := SECPLUGS_API_FILE_UPLOAD_URL + "?sha256=" + cksum
	httpclient := &http.Client{}
	req, err := http.NewRequest("GET", secplugsBaseUrl, nil)
	if err != nil {
		return UploadInfo{}, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("x-api-key", client.ApiKey)
	resp, err := httpclient.Do(req)
	if err != nil {
		return UploadInfo{}, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return UploadInfo{}, err
	}
	var info UploadInfo
	jsonerr := json.Unmarshal(respBytes, &info)
	if jsonerr != nil {
		return UploadInfo{}, jsonerr
	}
	return info, nil
}

func (info UploadInfo) GetFields() (map[string]string, error) {
	b, _ := json.Marshal(info.UploadPost.Fields)
	fields := make(map[string]string)
	err := json.Unmarshal(b, &fields)
	if err != nil {
		return fields, err
	}
	return fields, nil
}

func (info UploadInfo) GetUrl() string {
	return info.UploadPost.URL
}

// QuickScan performs a Secplugs file quickscan, which is sufficient for most
// usecases of a file scan. The most common type of file scans at Secplugs.
func (client *ScanClient) QuickScan() (ScanResult, error) {
	scanContext := map[string]string{
		"filename":    client.FileName,
		"client_uuid": "test-uuid",
	}
	ctxt, err := json.Marshal(scanContext)
	if err != nil {
		return ScanResult{}, err
	}
	url := SECPLUGS_API_FILE_QUICKSCAN
	sha256 := client.UniqueId

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ScanResult{}, err
	}
	req.Header.Add("x-api-key", client.ApiKey)
	q := req.URL.Query()
	q.Add("sha256", sha256)
	q.Add("scancontext", string(ctxt))
	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return ScanResult{}, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ScanResult{}, err
	}
	var result ScanResult
	jsonerr := json.Unmarshal(respBytes, &result)
	if jsonerr != nil {
		return ScanResult{}, jsonerr
	}
	return result, nil
}
