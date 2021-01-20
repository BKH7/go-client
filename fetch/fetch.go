package fetch

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

// MsgStruct ...
type MsgStruct struct {
	ID     int    `json:"Msg_id"`
	Sender string `json:"Sender"`
	Msg    string `json:"Msg"`
}

// HTTPClienter handler
type HTTPClienter interface {
	Do(req *http.Request) (*http.Response, error)
}

// PostMsg ...
func PostMsg(c HTTPClienter, body *MsgStruct) error {
	j, err := json.Marshal(body)
	if err != nil {
		return err
	}

	r, err := http.NewRequest(http.MethodPost, `http://iambk.ga/kafka/`, bytes.NewBuffer(j))
	// r.Header.Add("Authorization", "Bearer ")
	r.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}

	resp, err := c.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)

	var m map[string]interface{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	logrus.Infof("%s", m)
	return nil
}
