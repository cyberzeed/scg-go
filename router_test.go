package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/kr/pretty"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader, headers map[string]string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	if headers != nil {
		for key, value := range headers {
			req.Header.Set(key, value)
		}
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

var config = setupConfig("scg")
var router = setupRouter(config)

func TestGetRequestValuesOfScgSeries7Items(t *testing.T) {
	resp := performRequest(router, "GET", "/series?size=7", nil, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data map[string][]int
	body := resp.Body.String()
	err := json.Unmarshal([]byte(body), &data)
	series, exists := data["series"]
	expect := []int{3, 5, 9, 15, 23, 33, 45}

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, expect, series)
}

func TestGetRequestValueOfScgSeriesIndex(t *testing.T) {
	values := []int{3, 5, 9, 15, 23, 33, 45}

	for index, expect := range values {
		uri := fmt.Sprintf("/series/%v", index)
		resp := performRequest(router, "GET", uri, nil, nil)

		assert.Equal(t, http.StatusOK, resp.Code)

		var data map[string]int
		body := resp.Body.String()
		err := json.Unmarshal([]byte(body), &data)
		value, exists := data["value"]

		assert.Nil(t, err)
		assert.True(t, exists)
		assert.Equal(t, expect, value)
	}
}

func TestGetRestaurantInBangSue(t *testing.T) {
	resp := performRequest(router, "GET", "/restaurant/bangsue", nil, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data map[string][]interface{}
	body := resp.Body.String()
	err := json.Unmarshal([]byte(body), &data)
	results, exists := data["results"]

	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Len(t, results, 20)
}

func TestBroadcastMessage(t *testing.T) {
	b, err := json.Marshal(map[string][]string{
		"messages": []string{"broadcast"},
	})
	if err != nil {
		assert.Fail(t, err.Error())
	}

	broadcast := string(b)
	resp := performRequest(router, "POST", "/message/broadcast", strings.NewReader(broadcast), nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data map[string]string
	body := resp.Body.String()
	err = json.Unmarshal([]byte(body), &data)
	_, exists := data["error"]

	assert.Nil(t, err)
	assert.True(t, exists)
	pretty.Println(data)
}
