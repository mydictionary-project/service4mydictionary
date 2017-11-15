package cache4mydictionary

// ItemStruct : item struct
type ItemStruct struct {
	QueryString  string   `json:"queryString"`
	Word         string   `json:"word"`
	Define       []string `json:"define"`
	Status       string   `json:"status"`
	CreationTime int64    `json:"creationTime"`
}
