package database

type Configure struct {
	Driver      string `json:"driver"`
	Source      string `json:"source"`
	Database    string `json:"database"`
	TablePrefix string `json:"table_prefix"`
	IsPools     bool   `json:"is_pools"`
	/*
		Server      string `json:"server"`
		Port        int32  `json:"port"`
		Charset     string `json:"charset"`
		DbPrefix    string `json:"db_prefix"`
	*/
}
