package user

import (
	"github.com/m3o/m3o-go/client"
)

func NewUserService(token string) *UserService {
	return &UserService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type UserService struct {
	client *client.Client
}

func (t *UserService) Create(request CreateRequest) (*CreateResponse, error) {
	rsp := &CreateResponse{}
	return rsp, t.client.Call("user", "Create", request, rsp)
}

func (t *UserService) Delete(request DeleteRequest) (*DeleteResponse, error) {
	rsp := &DeleteResponse{}
	return rsp, t.client.Call("user", "Delete", request, rsp)
}

func (t *UserService) Login(request LoginRequest) (*LoginResponse, error) {
	rsp := &LoginResponse{}
	return rsp, t.client.Call("user", "Login", request, rsp)
}

func (t *UserService) Logout(request LogoutRequest) (*LogoutResponse, error) {
	rsp := &LogoutResponse{}
	return rsp, t.client.Call("user", "Logout", request, rsp)
}

func (t *UserService) Read(request ReadRequest) (*ReadResponse, error) {
	rsp := &ReadResponse{}
	return rsp, t.client.Call("user", "Read", request, rsp)
}

func (t *UserService) ReadSession(request ReadSessionRequest) (*ReadSessionResponse, error) {
	rsp := &ReadSessionResponse{}
	return rsp, t.client.Call("user", "ReadSession", request, rsp)
}

func (t *UserService) SendVerificationEmail(request SendVerificationEmailRequest) (*SendVerificationEmailResponse, error) {
	rsp := &SendVerificationEmailResponse{}
	return rsp, t.client.Call("user", "SendVerificationEmail", request, rsp)
}

func (t *UserService) UpdatePassword(request UpdatePasswordRequest) (*UpdatePasswordResponse, error) {
	rsp := &UpdatePasswordResponse{}
	return rsp, t.client.Call("user", "UpdatePassword", request, rsp)
}

func (t *UserService) Update(request UpdateRequest) (*UpdateResponse, error) {
	rsp := &UpdateResponse{}
	return rsp, t.client.Call("user", "Update", request, rsp)
}

func (t *UserService) VerifyEmail(request VerifyEmailRequest) (*VerifyEmailResponse, error) {
	rsp := &VerifyEmailResponse{}
	return rsp, t.client.Call("user", "VerifyEmail", request, rsp)
}

type Account struct {
	Created          int64             `json:"created"`
	Email            string            `json:"email"`
	Id               string            `json:"id"`
	Profile          map[string]string `json:"profile"`
	Updated          int64             `json:"updated"`
	Username         string            `json:"username"`
	VerificationDate int64             `json:"verificationDate"`
	Verified         bool              `json:"verified"`
}

type CreateRequest struct {
	Email    string            `json:"email"`
	Id       string            `json:"id"`
	Password string            `json:"password"`
	Profile  map[string]string `json:"profile"`
	Username string            `json:"username"`
}

type CreateResponse struct {
	Account Account `json:"account"`
}

type DeleteRequest struct {
	Id string `json:"id"`
}

type DeleteResponse struct {
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	Session Session `json:"session"`
}

type LogoutRequest struct {
	SessionId string `json:"sessionId"`
}

type LogoutResponse struct {
}

type ReadRequest struct {
	Email    string `json:"email"`
	Id       string `json:"id"`
	Username string `json:"username"`
}

type ReadResponse struct {
	Account Account `json:"account"`
}

type ReadSessionRequest struct {
	SessionId string `json:"sessionId"`
}

type ReadSessionResponse struct {
	Session Session `json:"session"`
}

type SendVerificationEmailRequest struct {
	Email              string `json:"email"`
	FailureRedirectUrl string `json:"failureRedirectUrl"`
	FromName           string `json:"fromName"`
	RedirectUrl        string `json:"redirectUrl"`
	Subject            string `json:"subject"`
	TextContent        string `json:"textContent"`
}

type SendVerificationEmailResponse struct {
}

type Session struct {
	Created int64  `json:"created"`
	Expires int64  `json:"expires"`
	Id      string `json:"id"`
	UserId  string `json:"userId"`
}

type UpdatePasswordRequest struct {
	ConfirmPassword string `json:"confirmPassword"`
	NewPassword     string `json:"newPassword"`
	OldPassword     string `json:"oldPassword"`
	UserId          string `json:"userId"`
}

type UpdatePasswordResponse struct {
}

type UpdateRequest struct {
	Email    string            `json:"email"`
	Id       string            `json:"id"`
	Profile  map[string]string `json:"profile"`
	Username string            `json:"username"`
}

type UpdateResponse struct {
}

type VerifyEmailRequest struct {
	Token string `json:"token"`
}

type VerifyEmailResponse struct {
}
