package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"simple-webapp/page/model"
	"strings"
	"testing"
)

type MockStorage struct {
	testName string
}

func (s *MockStorage) SavePage(p *model.Page) error {
	return nil
}

func (s *MockStorage) GetPage(title string) (*model.Page, error) {
	switch s.testName {
	case "OK":
		return &model.Page{Title: title, Data: []byte("test body")}, nil
	case "GetPageError":
		return nil, errors.New("Can not find page")
	}
	panic("no cases in mock function GetPage")
}

func TestRead(t *testing.T) {
	//Arrange
	tests := []struct {
		testName       string
		title          string
		expectedTitle  string
		expectedData   string
		expectedStatus int
	}{
		{
			testName:       "OK",
			title:          "Good test",
			expectedTitle:  "<h1>Good test</h1>",
			expectedData:   "<div>test body</div>",
			expectedStatus: 200,
		},
		{
			testName:       "GetPageError",
			title:          "Bad test",
			expectedStatus: 302,
		},
	}

	for _, tt := range tests {
		//Mocks
		rw := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Error(err)
		}
		ms := &MockStorage{testName: tt.testName}

		//Act
		Read(rw, req, ms, tt.title)

		//Assert
		if rw.Result().StatusCode != tt.expectedStatus {
			t.Errorf("testName=%s: expected status %d, but got %d", tt.testName, tt.expectedStatus, rw.Result().StatusCode)
		}

		body := rw.Body.String()
		if !strings.Contains(body, tt.expectedTitle) {
			t.Errorf("testName=%s: expected title not found in body", tt.testName)
		}
		if !strings.Contains(body, tt.expectedData) {
			t.Errorf("testName=%s: expected data not found in body", tt.testName)
		}

	}

}
