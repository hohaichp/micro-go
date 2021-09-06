package db

import (
	"github.com/m3o/m3o-go/client"
)

func NewDbService(token string) *DbService {
	return &DbService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type DbService struct {
	client *client.Client
}

func (t *DbService) Create(request CreateRequest) (*CreateResponse, error) {
	rsp := &CreateResponse{}
	return rsp, t.client.Call("db", "Create", request, rsp)
}

func (t *DbService) Delete(request DeleteRequest) (*DeleteResponse, error) {
	rsp := &DeleteResponse{}
	return rsp, t.client.Call("db", "Delete", request, rsp)
}

func (t *DbService) Read(request ReadRequest) (*ReadResponse, error) {
	rsp := &ReadResponse{}
	return rsp, t.client.Call("db", "Read", request, rsp)
}

func (t *DbService) Truncate(request TruncateRequest) (*TruncateResponse, error) {
	rsp := &TruncateResponse{}
	return rsp, t.client.Call("db", "Truncate", request, rsp)
}

func (t *DbService) Update(request UpdateRequest) (*UpdateResponse, error) {
	rsp := &UpdateResponse{}
	return rsp, t.client.Call("db", "Update", request, rsp)
}

type CreateRequest struct {
	Record map[string]interface{} `json:"record"`
	Table  string                 `json:"table"`
}

type CreateResponse struct {
	Id string `json:"id"`
}

type DeleteRequest struct {
	Id    string `json:"id"`
	Table string `json:"table"`
}

type DeleteResponse struct {
}

type ReadRequest struct {
	Id      string `json:"id"`
	Limit   int32  `json:"limit"`
	Offset  int32  `json:"offset"`
	Order   string `json:"order"`
	OrderBy string `json:"orderBy"`
	Query   string `json:"query"`
	Table   string `json:"table"`
}

type ReadResponse struct {
	Records map[string]interface{} `json:"records"`
}

type TruncateRequest struct {
	Table string `json:"table"`
}

type TruncateResponse struct {
	Table string `json:"table"`
}

type UpdateRequest struct {
	Id     string                 `json:"id"`
	Record map[string]interface{} `json:"record"`
	Table  string                 `json:"table"`
}

type UpdateResponse struct {
}
