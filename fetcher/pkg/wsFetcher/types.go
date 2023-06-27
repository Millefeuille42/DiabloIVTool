package wsFetcher

type socketData struct {
	Type string `json:"t"`
	Data struct {
		Body struct {
			Status string      `json:"s,omitempty"`
			Data   interface{} `json:"d"`
			Title  string      `json:"p"`
		} `json:"b,omitempty"`
	} `json:"d"`
}
