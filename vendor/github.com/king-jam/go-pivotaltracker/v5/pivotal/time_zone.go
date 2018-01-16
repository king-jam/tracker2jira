package pivotal

// TimeZone is ..
type TimeZone struct {
	OlsonName string `json:"olson_name,omitempty"`
	Offset    string `json:"offset,omitempty"`
}
