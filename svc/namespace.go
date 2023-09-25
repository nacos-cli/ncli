package svc

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	uparam "net/url"
	"strconv"
)

func NsCreate(login *LoginResult, namespaceId, namespaceName, description string) (bool, error) {

	exist, err := NsExist(login, namespaceId)
	if err != nil {
		return false, err
	}
	if exist {
		return true, nil
	}

	ok, err := create(login, namespaceId, namespaceName, description)
	if err != nil {
		return false, err
	}
	return ok, nil

}

func NsExist(login *LoginResult, namespaceId string) (bool, error) {

	const path = "/v1/console/namespaces"
	url := login.Server.url() + path
	query := uparam.Values{}
	query.Set("checkNamespaceIdExist", "true")
	query.Set("accessToken", login.Token)
	query.Set("customNamespaceId", namespaceId)
	empty := ""
	query.Set("namespaceId", empty)

	resp, err := http.Get(url + "?" + query.Encode())
	if err != nil {
		return false, errors.New("http get failed, caused by:" + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, errors.New("read response body error:" + err.Error())
		}
		body := string(bytes)
		exist, err := strconv.ParseBool(body)
		if err != nil {
			return false, errors.New(fmt.Sprintf("decode response body %s failed, error: %s", body, err.Error()))
		}
		return exist, nil
	} else {
		return false, errors.New("http response status:" + resp.Status)
	}

}

func create(login *LoginResult, namespaceId, namespaceName, description string) (bool, error) {

	const path = "/v1/console/namespaces"
	url := login.Server.url() + path
	postBody := uparam.Values{}
	postBody.Set("checkNamespaceIdExist", "true")
	postBody.Set("accessToken", login.Token)
	postBody.Set("customNamespaceId", namespaceId)
	postBody.Set("namespaceName", namespaceName)
	postBody.Set("namespaceDesc", description)

	resp, err := http.PostForm(url, postBody)
	if err != nil {
		return false, errors.New("http post failed, caused by:" + err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, errors.New("read response body error:" + err.Error())
		}

		body := string(bytes)
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
