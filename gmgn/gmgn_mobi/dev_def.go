package gmgn_mobi

type DeviceInfo struct {
	DeviceID string `json:"device_id"`
	ClientID string `json:"client_id"`
	FromApp  string `json:"from_app"`
	AppVer   string `json:"app_ver"`
	Pkg      string `json:"pkg"`
	AppLang  string `json:"app_lang"`
	SysLang  string `json:"sys_lang"`
	Brand    string `json:"brand"`
	Model    string `json:"model"`
	Os       string `json:"os"`
	OsAPI    string `json:"os_api"`
	TzName   string `json:"tz_name"`
	TzOffset string `json:"tz_offset"`
	Gpv      string `json:"gpv"`
}
