package orientrest

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
)

const (
	libVersion     = "1.0"
	defaultBaseURL = "http://localhost:2480/"
	userAgent      = "go-orientrest/" + libVersion

	defaultMediaType = "application/octet-stream"
)

// A Client manages communications with the OrientDB Server API.
type Client struct {
	// HTTP client
	client *http.Client

	// Base URL for API request. It should be specified with a trailing slash
	BaseUrl *url.URL

	// User agent used when communicating with the OrientDB Server API
	UserAgent string

	// Optional
	UserInfo *url.Userinfo
}

// New returns a new OrientDB API client.
func New(uri string) (*Client, error) {
	if uri == "" {
		uri = defaultBaseURL
	}

	if !strings.HasSuffix(uri, "/") {
		uri += "/"
	}

	baseUrl, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient

	c := &Client{client: httpClient, BaseUrl: baseUrl, UserAgent: userAgent}

	return c, nil
}

func (c *Client) Open(dbname, username, password string) (*Database, error) {
	if username != "" && password != "" {
		if dbname == "" {
			return nil, fmt.Errorf("orientrest: Database name cannot be empty")
		}
		c.UserInfo = url.UserPassword(username, password)

		db := &Database{name: dbname, client: c}
		return db.connect()
	}
	return nil, fmt.Errorf("orientrest: Username and Password cannot be empty")
}

func (c *Client) Auth(username, password string) (*Admin, error) {
	if username != "" && password != "" {
		c.UserInfo = url.UserPassword(username, password)
		return &Admin{client: c}, nil
	}
	return nil, fmt.Errorf("orientrest: Username and Password cannot be empty")
}

// NewRequest generates an API request It handles encoding
// parameters and attaching the appropriate headers.
func (c *Client) NewRequest(method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	uri := c.BaseUrl.ResolveReference(rel)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, uri.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept-Encoding", "gzip,deflate")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	// Set HTTP Basic authentication if UserInfo is supplied
	if userInfo := c.UserInfo; userInfo != nil {
		pwd, _ := userInfo.Password()
		req.SetBasicAuth(userInfo.Username(), pwd)
	}

	return req, nil
}

func (c *Client) NewUploadRequest(path string, reader io.Reader, mediaType string, size int64) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	uri := c.BaseUrl.ResolveReference(rel)

	req, err := http.NewRequest("POST", uri.String(), reader)
	if err != nil {
		return nil, err
	}
	req.ContentLength = size

	if len(mediaType) == 0 {
		mediaType = defaultMediaType
	}

	req.Header.Add("Content-Type", mediaType)
	req.Header.Add("Accept", "application/json")
	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

//  Do is used to send an API request and parse the response.
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	err = CheckResponse(resp)
	if err != nil {
		return err
	}

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if v != nil {
		return json.Unmarshal(resBody, v)
	}

	return nil
}

func CheckResponse(r *http.Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	var errorR interface{}
	data, err := ioutil.ReadAll(r.Body)
	if err == nil && data != nil {
		json.Unmarshal(data, &errorR)
	}
	return fmt.Errorf("%+v", errorR)
}
