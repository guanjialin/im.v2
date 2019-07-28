package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"im.v2/enum"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// 以json body请求url，other为nil表示不增加额外的header
// 注意：不支持other重写content-type
func HTTPPostJSON(url string, body interface{}, other http.Header) (*http.Response, error) {
	buf, err := json.Marshal(&body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	req.Header = other
	req.Header.Set(enum.ContentType, enum.ContentTypeJSON)

	return http.DefaultClient.Do(req)
}

func HTTPReadRespToJson(resp *http.Response, model interface{}) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("resp body: %s", body)

	err = json.Unmarshal(body, &model)

	return err
}

// 以json body请求url，并将response解析为model
func HTTPPostJsonParse(url string, body interface{}, model interface{}) error {
	resp, err := HTTPPostJSON(url, body, http.Header{enum.Accept: []string{enum.ContentTypeJSON}})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("status should be 200, but is " + strconv.Itoa(resp.StatusCode))
	}

	err = HTTPReadRespToJson(resp, model)

	return err
}
