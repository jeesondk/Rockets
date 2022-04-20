package tests

import (
	"RocketService/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCanHandleJsonRequest(t *testing.T) {
	request := `{
					"metadata": {
						"channel": "193270a9-c9cf-404a-8f83-838e71d9ae67",
						"messageNumber": 1,    
						"messageTime": "2022-02-02T19:39:05.86337+01:00",                                          
						"messageType": "RocketLaunched"                             
					},
					"message": {                                                    
						"type": "Falcon-9",
						"launchSpeed": 500,
						"mission": "ARTEMIS"  
					}
				}`

	c := controllers.NewController()
	w := httptest.NewRecorder()
	r := gin.Default()

	r.POST("/messages", c.Messages)
	req, _ := http.NewRequest("POST", "/messages", strings.NewReader(request))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
