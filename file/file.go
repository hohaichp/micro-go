package file

import(
	"github.com/m3o/m3o-go/client"
)

func NewFileService(token string) *FileService {
	return &FileService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type FileService struct {
	client *client.Client
}


func (t *FileService) List(request ListRequest) (*ListResponse, error) {
	rsp := &ListResponse{}
	return rsp, t.client.Call("file", "List", request, rsp)
}

func (t *FileService) Save(request SaveRequest) (*SaveResponse, error) {
	rsp := &SaveResponse{}
	return rsp, t.client.Call("file", "Save", request, rsp)
}




type File struct {
  Created int64 `json:"created"`
  FileContents string `json:"fileContents"`
  IsDirectory bool `json:"isDirectory"`
  Name string `json:"name"`
  Path string `json:"path"`
  Project string `json:"project"`
  Updated int64 `json:"updated"`
}

type ListRequest struct {
  Path string `json:"path"`
  Project string `json:"project"`
}

type ListResponse struct {
  File []File `json:"file"`
}

type SaveRequest struct {
  File []File `json:"file"`
}

type SaveResponse struct {
}

