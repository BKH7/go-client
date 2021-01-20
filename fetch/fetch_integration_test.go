package fetch

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchIntegration(t *testing.T) {
	fn := PostMsg(&http.Client{}, &MsgStruct{
		ID:     1,
		Sender: "Tom",
		Msg:    "Hello",
	})
	assert.Nil(t, fn)
}
