package client

// Locked enable/disable to edit Pipeline over the UI
type Locked struct {
	UI            bool   `json:"ui"`
	Description   string `json:"description"`
	AllowUnlockUI bool   `json:"allowUnlockUi"`
}
