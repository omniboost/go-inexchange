package inexchange

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/http/httputil"
	"net/textproto"
	"net/url"
	"path"
	"strings"

	"github.com/hashicorp/go-multierror"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-inexchange/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "testapi.inexchange.se",
		Path:   "/v1/api/",
	}
)

// NewClient returns a new InvoiceXpress Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{
		http: httpClient,
	}

	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with InvoiceXpress Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	apiKey      string
	clientToken string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c *Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) APIKey() string {
	return c.apiKey
}

func (c *Client) SetAPIKey(apiKey string) {
	c.apiKey = apiKey
}

func (c *Client) ClientToken() string {
	return c.clientToken
}

func (c *Client) SetClientToken(clientToken string) {
	c.clientToken = clientToken
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c *Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c *Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c *Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) GetEndpointURL(relative string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	relativeURL, err := url.Parse(relative)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = path.Join(clientURL.Path, relativeURL.Path)

	query := url.Values{}
	for k, v := range clientURL.Query() {
		query[k] = append(query[k], v...)
	}
	for k, v := range relativeURL.Query() {
		query[k] = append(query[k], v...)
	}
	clientURL.RawQuery = query.Encode()

	tmpl, err := template.New("endpoint_url").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}
	clientURL.Path = buf.String()

	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	contentType := fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset())

	// convert body struct to json
	buf := new(bytes.Buffer)

	if req.RequestBodyInterface() != nil {
		// Request has a method that returns a request body
		if r, ok := req.RequestBodyInterface().(io.Reader); ok {
			// request body is a io.Reader
			_, err := io.Copy(buf, r)
			if err != nil {
				return nil, err
			}
		} else {
			// request body is a struct/slice; marshal to json
			err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
			if err != nil {
				return nil, err
			}
		}
	} else if i, ok := req.(interface{ FormParamsInterface() Form }); ok {
		if i.FormParamsInterface().IsMultiPart() {
			// @TODO implement this as RequestBodyInterface()
			// Request has a form as body
			var err error
			w := multipart.NewWriter(buf)

			for k, f := range i.FormParamsInterface().Files() {
				var part io.Writer
				if x, ok := f.Content.(io.Closer); ok {
					defer x.Close()
				}

				if part, err = CreateFormFile(w, f.Content, k, f.Filename); err != nil {
					return nil, err
				}

				if _, err = io.Copy(part, f.Content); err != nil {
					return nil, err
				}
			}

			for k := range i.FormParamsInterface().Values() {
				var part io.Writer

				// Add other fields
				if part, err = w.CreateFormField(k); err != nil {
					return nil, err
				}

				fv := strings.NewReader(i.FormParamsInterface().Values().Get(k))
				if _, err = io.Copy(part, fv); err != nil {
					return nil, err
				}
			}

			// Don't forget to close the multipart writer.
			// If you don't close it, your request will be missing the terminating boundary.
			w.Close()

			// Don't forget to set the content type, this will contain the boundary.
			contentType = w.FormDataContentType()

		} else {
			buf.WriteString(i.FormParamsInterface().Values().Encode())
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", contentType)
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("ClientToken", c.ClientToken())
	r.Header.Add("APIKey", c.APIKey())

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, responseBody any) (*http.Response, error) {
	if c.debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	// c.SleepUntilRequestRate()
	// c.RegisterRequestTimestamp(time.Now())

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	// interface implements io.Writer: write Body to it
	// if w, ok := response.Envelope.(io.Writer); ok {
	// 	_, err := io.Copy(w, httpResp.Body)
	// 	return httpResp, err
	// }

	// try to decode body into interface parameter
	if responseBody == nil {
		return httpResp, nil
	}

	errorResponse := &ErrorResponse{Response: httpResp}
	statusErrResponse := &StatusErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, []any{}, []any{responseBody, errorResponse, statusErrResponse})
	if err != nil {
		return httpResp, err
	}

	if errorResponse.Error() != "" {
		return httpResp, errorResponse
	}

	if statusErrResponse.Error() != "" {
		return httpResp, statusErrResponse
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv []any, optionalVv []any) error {
	if len(vv) == 0 && len(optionalVv) == 0 {
		return nil
	}

	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			return err
		}
	}

	for _, v := range optionalVv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)

		_ = dec.Decode(v)
	}

	return nil
}

// func (c *Client) RegisterRequestTimestamp(t time.Time) {
// 	if len(requestTimestamps) >= 5 {
// 		requestTimestamps = requestTimestamps[1:5]
// 	}
// 	requestTimestamps = append(requestTimestamps, t)
// }

// func (c *Client) SleepUntilRequestRate() {
// 	// Requestrate is 5r/1s

// 	// if there are less then 5 registered requests: execute the request
// 	// immediately
// 	if len(requestTimestamps) < 4 {
// 		return
// 	}

// 	// is the first item within 1 second? If it's > 1 second the request can be
// 	// executed imediately
// 	diff := time.Now().Sub(requestTimestamps[0])
// 	if diff >= time.Second {
// 		return
// 	}

// 	// Sleep for the time it takes for the first item to be > 1 second old
// 	time.Sleep(time.Second - diff)
// }

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; (c >= 200 && c <= 299) || c == 400 {
		return nil
	}

	err := checkContentType(r)
	if err != nil {
		return errors.New(r.Status)
	}

	// read data and copy it back
	data, err := io.ReadAll(r.Body)
	r.Body = io.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	if len(data) == 0 {
		return nil
	}

	// convert json to struct
	err = json.Unmarshal(data, errorResponse)
	if err != nil {
		return err
	}

	return errorResponse
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func CreateFormFile(w *multipart.Writer, data io.Reader, fieldname, filename string) (io.Writer, error) {
	var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

	escapeQuotes := func(s string) string {
		return quoteEscaper.Replace(s)
	}

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes(fieldname), escapeQuotes(filename)))

	// contentType, err := GetFileContentType(data)
	// if err != nil {
	// 	return nil, err
	// }
	// h.Set("Content-Type", contentType)

	h.Set("Content-Type", "text/xml")
	return w.CreatePart(h)
}

// func GetFileContentType(file io.Reader) (string, error) {
//
// 	// Only the first 512 bytes are used to sniff the content type.
// 	buffer := make([]byte, 512)
//
// 	_, err := file.Read(buffer)
// 	if err != nil {
// 		return "", err
// 	}
//
// 	// Use the net/http package's handy DectectContentType function. Always returns a valid
// 	// content-type by returning "application/octet-stream" if no others seemed to match.
// 	contentType := http.DetectContentType(buffer)
//
// 	return contentType, nil
// }

type StatusErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
}

func (r *StatusErrorResponse) Error() string {
	if r.Response.StatusCode != 0 && (r.Response.StatusCode < 200 || r.Response.StatusCode > 299) {
		return r.Response.Status
	}

	return ""
}

//	{
//	  "message": "The request is invalid.",
//	  "modelState": {
//	    "search.PartyId": [
//	      "The PartyId field is required."
//	    ]
//	  }
//	}
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response `json:"-"`

	Message    string              `json:"message"`
	ModelState map[string][]string `json:"modelState,omitempty"`
}

func (r *ErrorResponse) Error() string {
	if r.Message == "" && len(r.ModelState) == 0 {
		return ""
	}

	var err *multierror.Error
	if r.Message != "" {
		err = multierror.Append(err, errors.New(r.Message))
	}

	for _, msgs := range r.ModelState {
		for _, msg := range msgs {
			err = multierror.Append(err, errors.New(msg))
		}
	}

	return err.Error()
}
