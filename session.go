package orientrest

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"mime/multipart"
	"log"
	"io/ioutil"

	"github.com/ympons/napping"
)

// Hack based on Napping Send method to allow uploading files
func (s *OSession) Upload(r *napping.Request, file []byte) (response string, err error) {
	r.Method = strings.ToUpper(r.Method)
	//
	// Create a URL object from the raw url string.  This will allow us to compose
	// query parameters programmatically and be guaranteed of a well-formed URL.
	//
	u, err := url.Parse(r.Url)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	//
	// Default query parameters
	//
	p := url.Values{}
	//
	// Encode parameters
	//
	u.RawQuery = p.Encode()
	//
	// Create a Request object; if populated, Data field is JSON encoded as
	// request body
	//
	header := http.Header{}
	if s.Header != nil {
		for k, _ := range *s.Header {
			v := s.Header.Get(k)
			header.Set(k, v)
		}
	}
	var req *http.Request

	buf := new(bytes.Buffer)
	w := multipart.NewWriter(buf)
	part, err := w.CreateFormFile("filename", "dbimport")
	if err != nil {
		log.Printf("FormFile %v", err)
	}
	part.Write(file)

	err = w.Close()

	req, err = http.NewRequest(r.Method, u.String(), buf)

	header.Set("Content-Type", w.FormDataContentType())
	//
	// Merge Session and Request options
	//
	var userinfo *url.Userinfo
	if u.User != nil {
		userinfo = u.User
	}
	if s.Userinfo != nil {
		userinfo = s.Userinfo
	}
	// Prefer Request's user credentials
	if r.Userinfo != nil {
		userinfo = r.Userinfo
	}
	if r.Header != nil {
		for k, v := range *r.Header {
			header.Set(k, v[0]) // Is there always guarnateed to be at least one value for a header?
		}
	}
	if header.Get("Accept") == "" {
		header.Add("Accept", "application/json") // Default, can be overridden with Opts
	}
	req.Header = header
	//
	// Set HTTP Basic authentication if userinfo is supplied
	//
	if userinfo != nil {
		pwd, _ := userinfo.Password()
		req.SetBasicAuth(userinfo.Username(), pwd)
		if u.Scheme != "https" {
			//s.log("WARNING: Using HTTP Basic Auth in cleartext is insecure.")
		}
	}
	//
	// Execute the HTTP request
	//

	var client *http.Client
	if s.Client != nil {
		client = s.Client
	} else {
		client = &http.Client{}
		s.Client = client
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Body: %v", err)
		return
	}
	response = string(body)

	return
}
