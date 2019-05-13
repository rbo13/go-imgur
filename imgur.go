package imgur

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const imgurUploadURL = "https://api.imgur.com/3/image"

// Imgur sets the field for the required fields when uploading images to imgur.
type Imgur struct {
	apiKey string
	client *http.Client
}

// New returns pointer to imgur
func New(apiKey string) *Imgur {

	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	return &Imgur{
		apiKey: apiKey,
		client: &http.Client{
			Timeout:   time.Second * 10,
			Transport: netTransport,
		},
	}
}

// Upload uploads an image to imgur.
func (imgr *Imgur) Upload(filename string) (*Response, error) {
	fileEncoded, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	parameters := url.Values{"image": {base64.StdEncoding.EncodeToString(fileEncoded)}}

	req, err := http.NewRequest("POST", imgurUploadURL, strings.NewReader(parameters.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "https://github.com/rbo13/go-imgur")
	req.Header.Set("Authorization", "Client-ID "+imgr.apiKey)

	res, err := imgr.client.Do(req)
	if err != nil {
		return nil, err
	}

	img := &Response{}
	json.NewDecoder(res.Body).Decode(img)

	return img, nil
}
