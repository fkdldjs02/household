package tests

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/fkdldjs02/household/api"
	"github.com/fkdldjs02/household/conf"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/usjeong/testSuit"
)

var (
	caseOne = conf.NewCaseOne("develop")
	App     = setApp()
)

func setApp() *gin.Engine {
	r := testSuit.GetGinEngine()
	api.NewApp(caseOne)
	api.SetRouter(r)
	return r
}

func TestSetUser(t *testing.T) {
	data := url.Values{}
	data.Add("name", "testUser")
	data.Add("passwd", "veryverySeceret")

	suit := &testSuit.TestSuit{
		Router: App,
		Method: "POST",
		URL:    "/auth/user",
		Data:   data,
	}

	resp := suit.Do()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log(string(body))
	assert.Equal(t, http.StatusCreated, resp.Code)
}

func TestCategory(t *testing.T) {
	data := url.Values{}
	data.Add("name", "test")
	data.Add("budget", "400000")

	suit := &testSuit.TestSuit{
		Router: App,
		Method: "POST",
		URL:    "/category",
		Data:   data,
	}

	resp := suit.Do()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("식비 항목 생성 응답:  " + string(body))
	assert.Equal(t, http.StatusCreated, resp.Code)

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "GET",
		URL:    "/category",
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	t.Log("현재 생성된 항목들: " + string(body))
	assert.Equal(t, http.StatusOK, resp.Code)

	data = url.Values{}
	data.Add("name", "식비")
	data.Add("budget", "1000000")

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "PUT",
		URL:    "/category/1",
		Data:   data,
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.Code)

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "GET",
		URL:    "/category",
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	t.Log("수정 후 항목들: " + string(body))
	assert.Equal(t, http.StatusOK, resp.Code)
}
