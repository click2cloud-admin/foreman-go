package foreman

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestHosts(t *testing.T) {
	testJson, err := os.Open("testdata/all_hosts.json")
	if err != nil {
		t.Error(err)
	}
	defer testJson.Close()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.Copy(w, testJson)
	}))
	defer ts.Close()

	client := NewForemanClient(ts.URL)
	expected := []Host{
		Host{
			Name:              "node1",
			Id:                100,
			HostgroupId:       200,
			OperatingSystemId: 300,
		},
		Host{
			Name:              "node2",
			Id:                101,
			HostgroupId:       201,
			OperatingSystemId: 301,
		},
	}

	actual := client.Hosts()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}
