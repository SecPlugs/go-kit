package filescan

// ScanClient is the struct on which all scan features are defined.
type ScanClient struct {
	ApiKey   string
	Vendor   string
	FileName string
	UniqueId string
}

const SECPLUGS_API_BASE_URL string = "https://api.live.secplugs.com/security"
const SECPLUGS_API_FILE_UPLOAD_URL string = SECPLUGS_API_BASE_URL + "/file/upload"
const SECPLUGS_API_FILE_QUICKSCAN string = SECPLUGS_API_BASE_URL + "/file/quickscan"
const SECPLUGS_DEFAULT_API_KEY string = "r2iKI4q7Lu91Nu5uPl3eW3BPmRo4XK1ZbhLWtOKd"
const SECPLUGS_CLEAN_MID_SCORE int = 70
