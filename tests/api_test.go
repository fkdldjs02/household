package tests

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

func TestHousehold(t *testing.T) {
	data := url.Values{}
	data.Add("typeof", "지출")
	data.Add("categoryName", "식비")
	data.Add("content", "저녁 식사")
	data.Add("money", "10000")
	data.Add("author", "testUser")

	suit := &testSuit.TestSuit{
		Router: App,
		Method: "POST",
		URL:    "/household",
		Data:   data,
	}

	resp := suit.Do()
	body, _ := ioutil.ReadAll(resp.Body)
	t.Log("가계부 생성 응답: " + string(body))
	assert.Equal(t, http.StatusCreated, resp.Code)

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "GET",
		URL:    "/household/list/식비",
		Data:   data,
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	t.Log("변경 전 식비에 관한 기록: " + string(body))
	assert.Equal(t, http.StatusOK, resp.Code)

	data = url.Values{}
	data.Add("typeof", "지출")
	data.Add("categoryName", "식비")
	data.Add("content", "아침 식사")
	data.Add("money", "13000")
	data.Add("author", "testUser")

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "PUT",
		URL:    "/household/1",
		Data:   data,
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.Code)

	suit = &testSuit.TestSuit{
		Router: App,
		Method: "GET",
		URL:    "/household/list/식비",
		Data:   data,
	}

	resp = suit.Do()
	body, _ = ioutil.ReadAll(resp.Body)
	t.Log("변경 후 식비에 관한 기록: " + string(body))
	assert.Equal(t, http.StatusOK, resp.Code)

	os.Remove("./test.db")
}
