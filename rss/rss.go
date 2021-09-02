package rss

import(
	"github.com/m3o/m3o-go/client"
)

func NewRssService(token string) *RssService {
	return &RssService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type RssService struct {
	client *client.Client
}


func (t *RssService) Add(request AddRequest) (*AddResponse, error) {
	rsp := &AddResponse{}
	return rsp, t.client.Call("rss", "Add", request, rsp)
}

func (t *RssService) Feed(request FeedRequest) (*FeedResponse, error) {
	rsp := &FeedResponse{}
	return rsp, t.client.Call("rss", "Feed", request, rsp)
}

func (t *RssService) List(request ListRequest) (*ListResponse, error) {
	rsp := &ListResponse{}
	return rsp, t.client.Call("rss", "List", request, rsp)
}

func (t *RssService) Remove(request RemoveRequest) (*RemoveResponse, error) {
	rsp := &RemoveResponse{}
	return rsp, t.client.Call("rss", "Remove", request, rsp)
}




type AddRequest struct {
  Category string `json:"category"`
  Name string `json:"name"`
  Url string `json:"url"`
}

type AddResponse struct {
}

type Entry struct {
  Content string `json:"content"`
  Date string `json:"date"`
  Feed string `json:"feed"`
  Id string `json:"id"`
  Link string `json:"link"`
  Summary string `json:"summary"`
  Title string `json:"title"`
}

type Feed struct {
  Category string `json:"category"`
  Id string `json:"id"`
  Name string `json:"name"`
  Url string `json:"url"`
}

type FeedRequest struct {
  Limit int64 `json:"limit"`
  Name string `json:"name"`
  Offset int64 `json:"offset"`
}

type FeedResponse struct {
  Entries []Entry `json:"entries"`
}

type ListRequest struct {
}

type ListResponse struct {
  Feeds []Feed `json:"feeds"`
}

type RemoveRequest struct {
  Name string `json:"name"`
}

type RemoveResponse struct {
}

