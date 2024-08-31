package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/NatalNW7/cliqtree/pkg/config"
	"github.com/NatalNW7/cliqtree/pkg/server/handler"
	"github.com/gin-gonic/gin"
)

var (
	linkID string
)

func initConfig() error {
	err := config.Init()
	if err != nil {
		return fmt.Errorf("Error to initialize project configuration: %v", err)
	}
	handler.Init()
	return nil
}

func createGinContext(recorder *httptest.ResponseRecorder) (*gin.Context, error) {
	err := initConfig()
	if err != nil {
		return nil, err
	}
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL: &url.URL{},
	}

	return ctx, nil
}

func MakeGet(ctx *gin.Context, params gin.Params){
	ctx.Request.Method = "GET"
	ctx.Params = params
}

func MakePost(ctx *gin.Context, body interface{}){
	ctx.Request.Method = "POST"
	jsonBody, _ := json.Marshal(body)
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(jsonBody))
	ctx.Request.Header.Set("Content-Type", "application/json")
}

func TestCreateLink(t *testing.T){
	recorder := httptest.NewRecorder()
	ctx, err := createGinContext(recorder)
	if err != nil {
		t.Errorf("Something went wrong to create gin context: %v", err)
	}
	body := map[string]string{
		"url": "www.github.com",
	}

	MakePost(ctx, body)
	handler.CreateShortLink(ctx)

	if recorder.Code != http.StatusCreated {
		t.Errorf("expect %d and got: %d", http.StatusCreated, recorder.Code)
	}

	var response map[string]interface{}
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	if err != nil {
		t.Errorf("Not possible to unmarshal body response: %s", err.Error())
	}
	data := response["data"].(map[string]interface{})
	linkID = data["linkId"].(string)
}

func TestFindLink(t *testing.T){
	recorder := httptest.NewRecorder()
	ctx, err := createGinContext(recorder)
	if err != nil {
		t.Errorf("Something went wrong to create gin context: %v", err)
	}
	params := []gin.Param{
		{
			Key: "linkId",
			Value: linkID,
		},
	}

	MakeGet(ctx, params)
	handler.FindShortLinkByLinkId(ctx)

	if recorder.Code != http.StatusOK {
		t.Errorf("expect %d and got: %d",http.StatusOK , recorder.Code)
	}
}

func TestStatus(t *testing.T){
	recorder := httptest.NewRecorder()
	ctx, err := createGinContext(recorder)
	if err != nil {
		t.Errorf("Something went wrong to create gin context: %v", err)
	}
	params := []gin.Param{}

	MakeGet(ctx, params)
	handler.Status(ctx)

	if recorder.Code != http.StatusOK {
		t.Errorf("expect %d and got: %d",http.StatusOK , recorder.Code)
	}
	
	if recorder.Body == nil {
		t.Errorf("No body returned: %v", recorder.Body)
	}
}

func TestMethodNotAllowed(t *testing.T){
	recorder := httptest.NewRecorder()
	ctx, err := createGinContext(recorder)
	if err != nil {
		t.Errorf("Something went wrong to create gin context: %v", err)
	}
	ctx.Request.Method = http.MethodDelete
	ctx.Params = []gin.Param{
		{
			Key: "linkId",
			Value: "testId",
		},
	}

	handler.FindShortLinkByLinkId(ctx)

	if recorder.Code != http.StatusMethodNotAllowed {
		t.Errorf("expect %d and got: %d", http.StatusMethodNotAllowed, recorder.Code)
	}
}