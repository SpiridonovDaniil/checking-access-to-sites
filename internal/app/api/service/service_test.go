package service

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	server "siteAccess/internal/app/api/http"
	mock_http "siteAccess/internal/app/api/http/mocks"
	"siteAccess/internal/domain"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/stretchr/testify/assert"
)

var rec = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "endpoint_registered",
	Help: "number of endpoints used",
},
	[]string{"endpoints"},
)

func TestService_GetTime(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, serviceAnswer *domain.Answer, expectedError error)
	testTable := []struct {
		name                   string
		serviceAnswer          domain.Answer
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name: "create HTTP status 200",
			serviceAnswer: domain.Answer{
				Time: 720,
			},
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Answer, expectedError error) {
				s.EXPECT().GetTime(gomock.Any(), gomock.Any()).Return(serviceAnswer, expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       `{"time":720}`,
		},
		{
			name:                   "create bad request",
			expectedTestStatusCode: 400,
			expectedResponse:       "[getTimeHandler] search parameters are not specified",
		},
		{
			name: "create internal server error",
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Answer, expectedError error) {
				s.EXPECT().GetTime(gomock.Any(), gomock.Any()).Return(nil, expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[getTime] error in obtaining data about access to the site example.com"),
			expectedResponse:       "[getTimeHandler] [getTime] error in obtaining data about access to the site example.com",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 200" {
				testCase.mockBehavior(service, &testCase.serviceAnswer, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, nil, testCase.expectedError)
			}

			f := server.NewServer(service, rec)
			var url string
			if testCase.name == "create bad request" {
				url = "/site"
			} else {
				url = "/site?site=example.com"
			}
			req, err := http.NewRequest("GET", url, strings.NewReader(""))
			req.Header.Add("content-Type", "application/json")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}

func TestService_GetMinTime(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error)
	testTable := []struct {
		name                   string
		serviceAnswer          domain.Site
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name: "create HTTP status 200",
			serviceAnswer: domain.Site{
				Site: "example.com",
			},
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error) {
				s.EXPECT().GetMinTime(gomock.Any()).Return(serviceAnswer, expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       `{"site":"example.com"}`,
		},
		{
			name: "create internal server error",
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error) {
				s.EXPECT().GetMinTime(gomock.Any()).Return(nil, expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[getMinTime] error getting the site name with minimal access time"),
			expectedResponse:       "[getMinTimeHandler] [getMinTime] error getting the site name with minimal access time",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 200" {
				testCase.mockBehavior(service, &testCase.serviceAnswer, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, nil, testCase.expectedError)
			}

			f := server.NewServer(service, rec)
			url := "/min"
			req, err := http.NewRequest("GET", url, strings.NewReader(""))
			req.Header.Add("content-Type", "application/json")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}

func TestService_GetMaxTime(t *testing.T) {
	type mockBehavior func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error)
	testTable := []struct {
		name                   string
		serviceAnswer          domain.Site
		mockBehavior           mockBehavior
		expectedTestStatusCode int
		expectedError          error
		expectedResponse       string
	}{
		{
			name: "create HTTP status 200",
			serviceAnswer: domain.Site{
				Site: "example.com",
			},
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error) {
				s.EXPECT().GetMaxTime(gomock.Any()).Return(serviceAnswer, expectedError)
			},
			expectedTestStatusCode: 200,
			expectedResponse:       `{"site":"example.com"}`,
		},
		{
			name: "create internal server error",
			mockBehavior: func(s *mock_http.Mockservice, serviceAnswer *domain.Site, expectedError error) {
				s.EXPECT().GetMaxTime(gomock.Any()).Return(nil, expectedError)
			},
			expectedTestStatusCode: 500,
			expectedError:          fmt.Errorf("[getMaxTime] error getting the site name with maximum access time"),
			expectedResponse:       "[getMaxTimeHandler] [getMaxTime] error getting the site name with maximum access time",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			service := mock_http.NewMockservice(c)
			if testCase.name == "create HTTP status 200" {
				testCase.mockBehavior(service, &testCase.serviceAnswer, testCase.expectedError)
			}
			if testCase.name == "create internal server error" {
				testCase.mockBehavior(service, nil, testCase.expectedError)
			}

			f := server.NewServer(service, rec)
			url := "/max"
			req, err := http.NewRequest("GET", url, strings.NewReader(""))
			req.Header.Add("content-Type", "application/json")
			assert.NoError(t, err)

			resp, err := f.Test(req)
			assert.NoError(t, err)

			body, err := io.ReadAll(resp.Body)
			assert.Equal(t, testCase.expectedTestStatusCode, resp.StatusCode)
			assert.Equal(t, testCase.expectedResponse, string(body))
		})
	}
}
