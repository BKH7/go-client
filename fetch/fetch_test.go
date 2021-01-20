package fetch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeClient struct{}

func (fakeClient) Do(req *http.Request) (*http.Response, error) {
	resp := `{"Code":"OK","Created":"2021-01-20T15:49:20Z"}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, resp)
	}))

	defer ts.Close()
	return http.Get(ts.URL)
}

func TestPOST(t *testing.T) {

	fn := PostMsg(&fakeClient{}, &MsgStruct{
		ID:     1,
		Sender: "Tom",
		Msg:    "Hello",
	})
	assert.Nil(t, fn)
}
