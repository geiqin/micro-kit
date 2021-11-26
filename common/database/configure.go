package database

type Configure struct {
	Driver      string `json:"driver"`
	Source      string `json:"source"`
	Database    string `json:"database"`
	TablePrefix string `json:"table_prefix"`
	IsPools     bool   `json:"is_pools"`
}
