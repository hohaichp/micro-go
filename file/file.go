package file

import (
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

func (t *FileService) Delete(request DeleteRequest) (*DeleteResponse, error) {
	rsp := &DeleteResponse{}
	return rsp, t.client.Call("file", "Delete", request, rsp)
}

func (t *FileService) List(request ListRequest) (*ListResponse, error) {
	rsp := &ListResponse{}
	return rsp, t.client.Call("file", "List", request, rsp)
}

func (t *FileService) Read(request ReadRequest) (*ReadResponse, error) {
	rsp := &ReadResponse{}
	return rsp, t.client.Call("file", "Read", request, rsp)
}

func (t *FileService) Save(request SaveRequest) (*SaveResponse, error) {
	rsp := &SaveResponse{}
	return rsp, t.client.Call("file", "Save", request, rsp)
}

type BatchSaveRequest struct {
	Files []Record `json:"files"`
}

type BatchSaveResponse struct {
}

type DeleteRequest struct {
	Path    string `json:"path"`
	Project string `json:"project"`
}

type DeleteResponse struct {
}

type ListRequest struct {
	Path    string `json:"path"`
	Project string `json:"project"`
}

type ListResponse struct {
	Files []Record `json:"files"`
}

type ReadRequest struct {
	Path    string `json:"path"`
	Project string `json:"project"`
}

type ReadResponse struct {
	File Record `json:"file"`
}

type Record struct {
	Content  string            `json:"content"`
	Created  string            `json:"created"`
	Metadata map[string]string `json:"metadata"`
	Path     string            `json:"path"`
	Project  string            `json:"project"`
	Updated  string            `json:"updated"`
}

type SaveRequest struct {
	File Record `json:"file"`
}

type SaveResponse struct {
}
