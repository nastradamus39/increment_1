package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPostUrl(t *testing.T) {
	type want struct {
		response   string
		statusCode int
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name   string
		url    string
		method string
		body   io.Reader
		want   want
	}{
		{
			name:   "Получение полной ссылки по короткой",
			url:    "/",
			method: http.MethodPost,
			body:   strings.NewReader("http://ya.ru?x=fljdlfsdf&y=rweurowieur&z=sdkfhsdfisdf"),
			want: want{
				statusCode: 201,
				response:   "63333133353437386339383238396137393962663130396134383266323336656562353838396462393733306437343339633431646630333866383638626534",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			t.Logf(tt.name)

			request := httptest.NewRequest(tt.method, tt.url, tt.body)

			w := httptest.NewRecorder()
			h := http.HandlerFunc(Short)

			h.ServeHTTP(w, request)
			result := w.Result()
			response, _ := ioutil.ReadAll(result.Body)

			t.Logf("HTTP code - %d", result.StatusCode)
			t.Logf("HTTP body - %s", fmt.Sprintf("%x", response))

			assert.Equal(t, tt.want.statusCode, result.StatusCode)
			assert.Equal(t, tt.want.response, fmt.Sprintf("%x", response))
		})
	}
}
