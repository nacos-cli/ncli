package svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	uparam "net/url"
)

type LoginPost struct {
	Username string
	Password string
}

type LoginResult struct {
	Server
	Username string
	Token    string
	TokenTTl uint64
}

func Login(srv *Server, post *LoginPost) (*LoginResult, error) {

	result := new(LoginResult)
	result.Schema = srv.Schema
	result.Host = srv.Host
	result.Port = srv.Port
	result.Context = srv.Context

	login := "/v1/auth/users/login"
	url := srv.url() + login
	loginBody := uparam.Values{}
	loginBody.Set("username", post.Username)
	loginBody.Set("password", post.Password)

	resp, err := http.PostForm(url, loginBody)
	if err != nil {
		return result, errors.New("http post failed, caused by:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return result, errors.New("read response body error:" + err.Error())
		}

		type jsonResult struct {
			Token    string `json:"accessToken"`
			Ttl      uint64 `json:"tokenTtl"`
			Admin    bool   `json:"globalAdmin"`
			Username string `json:"username"`
		}
		var jResult jsonResult
		err = json.Unmarshal(bytes, &jResult)
		if err != nil {
			return result, errors.New(fmt.Sprintf("decode response body %s failed, error: %s", string(bytes), err.Error()))
		}

		if len(jResult.Token) == 0 {
			return result, errors.New("empty token")
		}
		result.Username = post.Username
		result.Token = jResult.Token
		result.TokenTTl = jResult.Ttl
		return result, nil

	} else {
		return result, errors.New("http response status:" + resp.Status)
	}

}
