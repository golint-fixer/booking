package booking

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebApp(t *testing.T) {
	service := testService{}

	app := newWebApp(service)
	ts := httptest.NewServer(app)
	defer ts.Close()

	// new test client
	client := newTestClient(ts.URL)

	// get homepage
	client.Get("/")
	if code := client.response.StatusCode; code != http.StatusOK {
		t.Error("want", http.StatusOK)
		t.Error("got ", code)
	}

	// ensure dates are listed
	want := []byte("February 1, 2015")
	if !bytes.Contains(client.body, want) {
		t.Error("want contains", string(want))
		t.Error("got", string(client.body))
	}

	want = []byte("February 2, 2015")
	if !bytes.Contains(client.body, want) {
		t.Error("want contains", string(want))
		t.Error("got", string(client.body))
	}

	// select two checkboxes
	// click Book
	// see registration
}

type testClient struct {
	url      string
	cookies  *http.CookieJar
	visited  []string       // list of visited urls. head is first, tail is last.
	response *http.Response // last response
	body     []byte
	code     int
}

func newTestClient(url string) *testClient {
	c := &testClient{}
	c.url = url
	return c
}

func (c *testClient) Response() *http.Response {
	return c.response
}

func (c *testClient) Get(path string) error {
	var err error
	c.response, err = http.Get(c.url + path)
	if err != nil {
		c.body = nil
		c.code = 0
		return err
	}
	c.body, err = ioutil.ReadAll(c.response.Body)
	defer c.response.Body.Close()
	if err != nil {
		return err
	}
	c.code = c.response.StatusCode

	return nil
}

type testService struct{}

func (ts testService) AvailableDays() ([]time.Time, error) {
	return []time.Time{
		time.Date(2015, 2, 1, 0, 0, 0, 0, time.UTC),
	}, nil
}
