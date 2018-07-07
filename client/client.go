package client

import (
	"fmt"
	"github.com/own3dh2so4/telegram/message"
	"net/http"
	"net/url"
	"time"
	"io"
	"io/ioutil"
	"encoding/json"
	"github.com/own3dh2so4/telegram/model"
	"strconv"
)

type Client interface {
	SendMessage(message message.Message) error
	GetUpdate(offset int64) (*model.UpdateResponse, error)
}

func NewClient() Client {
	return &httpClient{
		client: &http.Client{
			Timeout: time.Duration(config.Timeout) * time.Second,
		},
	}
}

type httpClient struct {
	client *http.Client
}

func (hc *httpClient) GetUpdate(offset int64) (*model.UpdateResponse, error) {
	URL := url.URL{
		Scheme: config.Schema,
		Host:   config.TelegramHost,
		Path:   fmt.Sprintf("%s%s/%s", config.TelegramPath, config.BotKey, config.GetUpdatePath),
	}
	if offset != 0 {
		URL.RawQuery = getQueryValues(map[string]string{"offset" : strconv.FormatInt(offset, 10)}).Encode()
	}
	fmt.Println( URL.String())

	req, err := http.NewRequest(http.MethodGet, URL.String(),nil)
	if err != nil {
		panic(err)
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		fmt.Println("Error sending message to telegram ", err)
		return nil, err
	}
	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()
	//TODO maybe handle de response
	body, _ := ioutil.ReadAll(resp.Body)
	var v model.UpdateResponse
	if len(body) > 0 {
		err = json.Unmarshal(body, &v)
	}
	return &v, err
}

func (hc *httpClient) SendMessage(msg message.Message) error {
	URL := url.URL{
		Scheme: config.Schema,
		Host:   config.TelegramHost,
		Path:   fmt.Sprintf("%s%s/%s", config.TelegramPath, config.BotKey, config.SendMessagePath),
	}

	URL.RawQuery = getQueryValues(msg.GetMapValues()).Encode()
	fmt.Println( URL.String())
	req, err := http.NewRequest(http.MethodGet, URL.String(),nil)
	if err != nil {
		panic(err)
	}

	resp, err := hc.client.Do(req)
	if err != nil {
		fmt.Println("Error sending message to telegram ", err)
		return err
	}
	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()
	//TODO maybe handle de response
	body, _ := ioutil.ReadAll(resp.Body)
	var v model.SendMessageResponse
	if len(body) > 0 {
		err = json.Unmarshal(body, &v)
	}
	return err
}

func getQueryValues(values map[string]string) url.Values {
	result := map[string][]string{}
	for k, v := range values {
		result[k] = []string{v}
	}
	return result
}
