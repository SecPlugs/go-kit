package filescan

type ScanResult struct {
	Score        int    `json:"score"`
	JSONReport   string `json:"json_report"`
	ThreatObject struct {
		Sha256 string `json:"sha256"`
	} `json:"threat_object"`
	ScanContext struct {
	} `json:"scan_context"`
	Datetime  float64 `json:"datetime"`
	Status    string  `json:"status"`
	AccountID string  `json:"account_id"`
	APIKey    string  `json:"api_key"`
	MetaData  struct {
		PluginInfo struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"plugin_info"`
		Capability string `json:"capability"`
		VendorInfo struct {
			VendorConfigName string `json:"vendor_config_name"`
			EntitlementLevel string `json:"entitlement_level"`
			Vendor           string `json:"vendor"`
			Credits          int    `json:"credits"`
			Params           struct {
			} `json:"params"`
		} `json:"vendor_info"`
	} `json:"meta_data"`
	ReportID string  `json:"report_id"`
	Verdict  string  `json:"verdict"`
	Duration float64 `json:"duration"`
}

type ProxyScanResult struct {
	Score        int    `json:"score"`
	JSONReport   string `json:"json_report"`
	ThreatObject struct {
		Sha256 string `json:"sha256"`
	} `json:"threat_object"`
	ScanContext struct {
	} `json:"scan_context"`
	Datetime  float64 `json:"datetime"`
	Status    string  `json:"status"`
	AccountID string  `json:"account_id"`
	APIKey    string  `json:"api_key"`
	MetaData  struct {
		PluginInfo struct {
			Type string `json:"type"`
			Name string `json:"name"`
		} `json:"plugin_info"`
		Capability string `json:"capability"`
		VendorInfo struct {
			VendorConfigName string `json:"vendor_config_name"`
			EntitlementLevel string `json:"entitlement_level"`
			Vendor           string `json:"vendor"`
			Credits          int    `json:"credits"`
			Params           struct {
			} `json:"params"`
		} `json:"vendor_info"`
	} `json:"meta_data"`
	ReportID      string  `json:"report_id"`
	Verdict       string  `json:"verdict"`
	Duration      float64 `json:"duration"`
	UserReportURL string  `json:"user_report_url"`
}
