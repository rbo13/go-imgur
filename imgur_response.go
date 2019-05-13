package imgur

// Response is the data given from imgur API.
type Response struct {
	Data    ImageData `json:"data"`
	Status  int       `json:"status"`
	Success bool      `json:"success"`
}

// ImageData contains the metada of the images from imgur.
type ImageData struct {
	AccountID  int    `json:"account_id"`
	Animated   bool   `json:"animated"`
	Bandwidth  int    `json:"bandwidth"`
	DateTime   int    `json:"datetime"`
	Deletehash string `json:"deletehash"`
	Favorite   bool   `json:"favorite"`
	Height     int    `json:"height"`
	ID         string `json:"id"`
	InGallery  bool   `json:"in_gallery"`
	IsAd       bool   `json:"is_ad"`
	Link       string `json:"link"`
	Name       string `json:"name"`
	Size       int    `json:"size"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Views      int    `json:"views"`
	Width      int    `json:"width"`
}

// GetImageLink returns the url of the uploaded image from imgur.
func (res *Response) GetImageLink() string {
	return res.Data.Link
}

// Clipboard copies the link to your clipboard.
// TODO:: implement this;
func (res *Response) Clipboard() error {
	return nil
}
