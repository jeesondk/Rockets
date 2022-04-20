package tests

import (
	"RocketService/controllers"
	"RocketService/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

	messageServiceMock := messageServiceMock{}
	messageServiceMock.On("HandleMessage", mock.Anything).Return(DTO.RocketLaunched{
		Type:        "Falcon-9",
		LaunchSpeed: 500,
		Mission:     "ARTEMIS",
	}, nil)

	c := controllers.Controller{MessageService: &messageServiceMock}
	w := httptest.NewRecorder()
	r := gin.Default()

	r.POST("/messages", c.ReceiveMessage)
	req, _ := http.NewRequest("POST", "/messages", strings.NewReader(request))
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
func TestCanHandleInvalidRequest(t *testing.T) {
	request := `{
						"mojo": {
							"what": "193270a9-c9cf-404a-8f83-838e71d9ae67",
							"goesbump": 1,    
							"at": "2022-02-02T19:39:05.86337+01:00",                                          
							"inthenight": "RocketLaunched"                             
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
	r.POST("/messages", c.ReceiveMessage)
	req, _ := http.NewRequest("POST", "/messages", strings.NewReader(request))
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assert.Equal(t, "\"Invalid request\"", w.Body.String())
}
