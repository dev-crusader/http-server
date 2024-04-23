package client

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/dev-crusader404/http-server/models"
	"github.com/dev-crusader404/http-server/startup"
	props "github.com/dev-crusader404/http-server/startup"
)

var client *rClient

func init() {
	client = &rClient{
		c: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns: 10,
			},
			Timeout: 60 * time.Second,
		},
	}
}

type rClient struct {
	c *http.Client
}

type restClient interface {
	do(*http.Request) (*http.Response, error)
}

func (cl *rClient) do(r *http.Request) (*http.Response, error) {
	return cl.c.Do(r)
}

func GetClient() *rClient {
	return client
}

func RestCall(ctx context.Context, client restClient, user, msg string) error {
	URL := props.GetAll().GetString("HOST", "")
	if URL == "" {
		log.Panic("no url found")
	}

	u, err := url.Parse(fmt.Sprintf("%s/%s", URL, "message"))
	if err != nil {
		log.Panic("error parsing url: " + URL)
	}

	reqBody := models.MessageBody{
		User: user,
		Msg:  msg,
	}

	byt, err := json.Marshal(reqBody)
	if err != nil {
		log.Print("error while marshalling")
		return err
	}
	req := &http.Request{
		Method: http.MethodGet,
		URL:    u,
		Header: map[string][]string{
			"Accept":        {"application/json"},
			"Content-Type":  {"application/json"},
			"Authorization": {GetBasicAuth()},
		},
		Body: io.NopCloser(bytes.NewBuffer(byt)),
	}

	resp, err := client.do(req)
	if err != nil {
		log.Printf("error during call: %s", err.Error())
		return err
	}

	if resp == nil || resp.Body == nil {
		err := fmt.Errorf("nil respose/body received")
		log.Println(err)
		return err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("unable to read response body")
		log.Println(err)
		return err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("\nunexpected status code: %d", resp.StatusCode)
		log.Println(err)
		return err
	}

	resMsg := models.HTTPResponse{}
	err = json.Unmarshal(b, &msg)
	if err != nil {
		err := fmt.Errorf("error while unmarshalling")
		log.Println(err)
		return err
	}
	log.Printf("Received Response: %+v", msg)
	log.Printf("\nMessage: %s", resMsg.Message.Text)
	return nil
}

func GetBasicAuth() string {
	user := startup.GetAll().MustGetString("BASIC-LOGIN")
	password := startup.GetAll().MustGetString("BASIC-PASSWORD")

	authByte := []byte(user + ":" + password)
	encodedAuth := base64.StdEncoding.EncodeToString(authByte)
	return "Basic " + encodedAuth
}

func MakeHTTPCall(user, msg string) {
	err := RestCall(context.TODO(), client, user, msg)
	if err != nil {
		log.Print(err)
	}
}
