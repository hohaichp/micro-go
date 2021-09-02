package image

import(
	"github.com/m3o/m3o-go/client"
)

func NewImageService(token string) *ImageService {
	return &ImageService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type ImageService struct {
	client *client.Client
}


func (t *ImageService) Convert(request ConvertRequest) (*ConvertResponse, error) {
	rsp := &ConvertResponse{}
	return rsp, t.client.Call("image", "Convert", request, rsp)
}

func (t *ImageService) Resize(request ResizeRequest) (*ResizeResponse, error) {
	rsp := &ResizeResponse{}
	return rsp, t.client.Call("image", "Resize", request, rsp)
}

func (t *ImageService) Upload(request UploadRequest) (*UploadResponse, error) {
	rsp := &UploadResponse{}
	return rsp, t.client.Call("image", "Upload", request, rsp)
}




type ConvertRequest struct {
  Base64 string `json:"base64"`
  Name string `json:"name"`
  OutputUrl bool `json:"outputUrl"`
  Url string `json:"url"`
}

type ConvertResponse struct {
  Base64 string `json:"base64"`
  Url string `json:"url"`
}

type CropOptions struct {
  Anchor string `json:"anchor"`
  Height int32 `json:"height"`
  Width int32 `json:"width"`
}

type Point struct {
  X int32 `json:"x"`
  Y int32 `json:"y"`
}

type Rectangle struct {
  Max Point `json:"max"`
  Min Point `json:"min"`
}

type ResizeRequest struct {
  Base64 string `json:"base64"`
  CropOptions CropOptions `json:"cropOptions"`
  Height int64 `json:"height"`
  Name string `json:"name"`
  OutputUrl bool `json:"outputUrl"`
  Url string `json:"url"`
  Width int64 `json:"width"`
}

type ResizeResponse struct {
  Base64 string `json:"base64"`
  Url string `json:"url"`
}

type UploadRequest struct {
  Base64 string `json:"base64"`
  Name string `json:"name"`
  Url string `json:"url"`
}

type UploadResponse struct {
  Url string `json:"url"`
}

