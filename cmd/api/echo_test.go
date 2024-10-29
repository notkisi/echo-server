package api

import (
	"echo-server/util"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoGet(t *testing.T) {
	message := util.RandomMessage(16)

	testCases := []struct {
		name          string
		message       string
		buildRequest  func() (*http.Request, error)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:    "OK",
			message: message,
			buildRequest: func() (*http.Request, error) {
				url := fmt.Sprintf("/echo?message=%s", message)
				request, err := http.NewRequest(http.MethodGet, url, nil)
				return request, err
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)

				var response map[string]string
				err := json.Unmarshal(recorder.Body.Bytes(), &response)
				require.NoError(t, err)
				require.Equal(t, message, response["message"])
			},
		},
		{
			name:    "NoMessageParameter",
			message: "",
			buildRequest: func() (*http.Request, error) {
				url := fmt.Sprintf("/echo?message=%s", "")
				request, err := http.NewRequest(http.MethodGet, url, nil)
				return request, err
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			server := newTestServer(t)
			recorder := httptest.NewRecorder()

			request, err := tc.buildRequest()
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
