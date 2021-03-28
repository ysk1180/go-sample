package api_test

import (
	"go-sample/mock-test-sample/api"
	"testing"
)

type HogeClientMock struct {
}

func (hc *HogeClientMock) FetchDataFromHoge() string {
	return "test"
}

func TestFetch(t *testing.T) {
	client := &api.Client{
		HC: &HogeClientMock{},
	}

	res := client.Fetch()
	expected := "data: test"

	if res != expected {
		t.Errorf("Output=%s, Expected=%s", res, expected)
	}
}
