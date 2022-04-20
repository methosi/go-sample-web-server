package upload

// Image struct
type Image struct {
	Filename    string `json:"fileName"`
	Size        int64  `json:"fileSize"`
	ContentType string `json:"contentType"`
}
