package model

//go:generate tscriptify -package=model -target=target_ts_file.ts Bank
type Bank struct {
	Name string `json:"name"`
}
