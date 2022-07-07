package handlers

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"net/http/httptest"
	"testing"
	"url-shortener/internal/service"
	mock_service "url-shortener/internal/service/mocks"
)

type CreateRequest struct {
	LongUrl string `json:"long_url"`
}

type FindRequest struct {
	ShortUrl string `uri:"short_url"`
}

func TestHandler_Create(t *testing.T) {

	// Init Test Table
	type mockBehavior func(r *mock_service.MockUrlService, createRequest CreateRequest)

	tests := []struct {
		name                 string
		inputBody            string
		inputCreateRequest   CreateRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			inputBody: `{"long_url": "https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies"}`,
			inputCreateRequest: CreateRequest{
				LongUrl: "https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies",
			},
			mockBehavior: func(r *mock_service.MockUrlService, createRequest CreateRequest) {
				r.EXPECT().Create(createRequest.LongUrl).Return("adw12afwqv", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"short_url":"adw12afwqv"}`,
		},
		{
			name:      "Wrong Input",
			inputBody: `{"long_url": "https//hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies"}`,
			inputCreateRequest: CreateRequest{
				LongUrl: "https//hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies",
			},
			mockBehavior:         func(r *mock_service.MockUrlService, createRequest CreateRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid URI for request"}`,
		},
		{
			name:      "Service Error",
			inputBody: `{"long_url": "https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies"}`,
			inputCreateRequest: CreateRequest{
				LongUrl: "https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies",
			},
			mockBehavior: func(r *mock_service.MockUrlService, createRequest CreateRequest) {
				r.EXPECT().Create(createRequest.LongUrl).Return("adw12afwqv", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			serviceMock := mock_service.NewMockUrlService(c)
			test.mockBehavior(serviceMock, test.inputCreateRequest)

			services := &service.Service{UrlService: serviceMock}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/api", handler.create)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api",
				bytes.NewBufferString(test.inputBody))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestHandler_Find(t *testing.T) {

	// Init Test Table
	type mockBehavior func(r *mock_service.MockUrlService, findRequest FindRequest)

	tests := []struct {
		name                 string
		inputUri             string
		inputFindRequest     FindRequest
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			inputUri: "adw12afwqv",
			inputFindRequest: FindRequest{
				ShortUrl: "adw12afwqv",
			},
			mockBehavior: func(r *mock_service.MockUrlService, findRequest FindRequest) {
				r.EXPECT().Find("adw12afwqv").Return("https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"long_url":"https://hh.ru/vacancy/66610729?hhtmFrom=employer_vacancies"}`,
		},
		{
			name:     "Wrong Input",
			inputUri: "adw12afwqv!",
			inputFindRequest: FindRequest{
				ShortUrl: "adw12afwqv!",
			},
			mockBehavior:         func(r *mock_service.MockUrlService, findRequest FindRequest) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"message":"invalid input type"}`,
		},
		{
			name:     "Service Error",
			inputUri: "adw12afwqv",
			inputFindRequest: FindRequest{
				ShortUrl: "adw12afwqv",
			},
			mockBehavior: func(r *mock_service.MockUrlService, findRequest FindRequest) {
				r.EXPECT().Find("adw12afwqv").Return("", errors.New("something went wrong"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			serviceMock := mock_service.NewMockUrlService(c)
			test.mockBehavior(serviceMock, test.inputFindRequest)

			services := &service.Service{UrlService: serviceMock}
			handler := Handler{services}

			// Init Endpoint
			r := gin.New()
			r.POST("/api/:short_url", handler.find)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/"+test.inputUri,
				bytes.NewBufferString(""))

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
