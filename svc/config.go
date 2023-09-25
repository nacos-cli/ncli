package svc

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	uparam "net/url"
	"strconv"
)

type ConfigQuery struct {
	Group       string `json:"group"`
	DataId      string `json:"dataId"`
	NamespaceId string `json:"tenant"`
}

type ConfigPost struct {
	Group       string
	DataId      string
	NamespaceId string
	Content     string
	Type        string
	App         string
	Tags        string
	Description string
}

type Config struct {
	Group       string `json:"group"`
	DataId      string `json:"dataId"`
	NamespaceId string `json:"tenant"`
	Content     string `json:"content"`
	Type        string `json:"type"`
	App         string `json:"appName"`
	Tags        string `json:"configTags"`
	Description string `json:"desc"`
	Md5         string `json:"md5,omitempty"`
	CreateTime  uint64 `json:"createTime,omitempty"`
	ModifyTime  uint64 `json:"modifyTime,omitempty"`
	CreateUser  string `json:"createUser,omitempty"`
	CreateIp    string `json:"createIp,omitempty"`
}

func ConfigGet(login *LoginResult, s *ConfigQuery) (*Config, error) {

	const path = "/v1/cs/configs"
	url := login.Server.url() + path
	query := uparam.Values{}
	query.Set("group", s.Group)
	query.Set("dataId", s.DataId)
	query.Set("namespaceId", s.NamespaceId)
	query.Set("tenant", s.NamespaceId)
	query.Set("show", "all")
	query.Set("accessToken", login.Token)
	query.Set("username", login.Username)

	cfg := new(Config)
	cfg.Group = s.Group
	cfg.DataId = s.DataId
	cfg.NamespaceId = s.NamespaceId

	furl := url + "?" + query.Encode()
	verboseFln("url: %s", furl)
	resp, err := http.Get(furl)
	if err != nil {
		return cfg, errors.New("http get failed, caused by:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return cfg, errors.New("read response body error:" + err.Error())
		}

		verboseFln("response: %s", string(bytes))
		err = json.Unmarshal(bytes, &cfg)
		if err != nil {
			return cfg, errors.New(fmt.Sprintf("decode response body %s failed, error: %s", string(bytes), err.Error()))
		}
		if GVerbose {
			fmt.Printf("config: %+v\n", cfg)
		}
		return cfg, nil
	} else {
		return cfg, errors.New("http response status:" + resp.Status)
	}

}

func ConfigCreateOrUpdate(login *LoginResult, p *ConfigPost) (bool, error) {

	const path = "/v1/cs/configs"
	url := login.Server.url() + path + "?accessToken=" + login.Token

	postBody := uparam.Values{}
	postBody.Set("dataId", p.DataId)
	postBody.Set("group", p.Group)
	postBody.Set("content", p.Content)
	postBody.Set("desc", p.Description)
	postBody.Set("type", p.Type)
	postBody.Set("appName", p.App)
	postBody.Set("tenant", p.NamespaceId)
	postBody.Set("namespaceId", p.NamespaceId)

	if err := cfgNamespaceOk(login, p); err != nil {
		return false, errors.New("Automatically create namespace '" + p.NamespaceId + "' failed, caused by:" + err.Error())
	}
	if GVerbose {
		fmt.Printf("url: %s, post:%+v, postData: %s\n", url, postBody, postBody.Encode())
	}

	resp, err := http.PostForm(url, postBody)
	if err != nil {
		return false, errors.New("http post failed, caused by:" + err.Error())
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, errors.New("read response body error:" + err.Error())
		}

		body := string(bytes)
		verboseFln("response: %s", body)
		ok, err := strconv.ParseBool(body)
		if err != nil {
			return false, errors.New(fmt.Sprintf("decode response body %s failed, error: %s", body, err.Error()))
		}
		if ok {
			return ok, nil
		} else {
			return false, errors.New("server response op result false")
		}

	} else {
		return false, errors.New("http response status:" + resp.Status)
	}

}

func cfgNamespaceOk(login *LoginResult, cfg *ConfigPost) error {
	app := cfg.App
	if len(app) == 0 {
		app = cfg.DataId
	}
	name := "auto-by-app-" + app
	desc := cfg.Description
	if len(desc) == 0 {
		desc = "auto-by-app-" + app
	}
	_, err := NsCreate(login, cfg.NamespaceId, name, desc)
	return err
}
