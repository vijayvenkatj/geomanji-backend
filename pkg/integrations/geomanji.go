package integrations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type GeolocateResult struct {
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
	Probability float64 `json:"Probability"`
}

type GeomanjiClient struct {
	BaseURL string
	Client  *http.Client
}

func NewGeomanjiClient() *GeomanjiClient {
	return &GeomanjiClient{
		BaseURL: os.Getenv("GEOMANJI_API_URL"),
		Client:  http.DefaultClient,
	}
}

func (c *GeomanjiClient) Geolocate(filename string, img io.Reader) (*GeolocateResult, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	part, err := w.CreateFormFile("img", filename)
	if err != nil {
		return nil, err
	}
	if _, err := io.Copy(part, img); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseURL+"/geolocate", &body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("geomanji api: unexpected status %d", resp.StatusCode)
	}

	var result GeolocateResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
