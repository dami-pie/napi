package responses

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type JSONTestCase struct {
	name string
	req  *http.Request
	want int
}

func TestDecodeJSON(t *testing.T) {
	tests := []JSONTestCase{
		{
			name: "Content-Type inválido",
			req:  &http.Request{Header: http.Header{"Content-Type": []string{"application/xml"}}, Body: nil},
			want: http.StatusUnsupportedMediaType,
		},
		{
			name: "JSON mal formado",
			req:  &http.Request{Body: io.NopCloser(io.Reader(strings.NewReader("{\"email: \"aaaa\"}")))},
			want: http.StatusBadRequest,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := DecodeJSON(&httptest.ResponseRecorder{}, test.req, "")

			var mr *MalformedRequest
			if errors.As(err, &mr) { //unwrap error
				if test.want != mr.Status {
					t.Errorf("Deveria receber %d, recebido foi %d", test.want, mr.Status)
				}
			}
		})
	}
}
