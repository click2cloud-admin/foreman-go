package foreman

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ForemanClient struct {
	URL      string
	client   *http.Client
	username string
	password string
}

type Host struct {
	Name              string `json:"name"`
	Id                int    `json:"id"`
	HostgroupId       int    `json:"hostgroup_id"`
	OperatingSystemId int    `json:"operatingsystem_id"`
}

func NewForemanClient(URL string) *ForemanClient {
	return &ForemanClient{
		URL:    URL,
		client: new(http.Client),
	}
}

func (f *ForemanClient) SetBasicAuth(username, password string) {
	f.username = username
	f.password = password
}

func (f ForemanClient) setBasicAuth(req *http.Request) {
	req.SetBasicAuth(f.username, f.password)
}

func (f ForemanClient) Hosts() []Host {

	url := fmt.Sprintf("%s/api/hosts", f.URL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
	}

	f.setBasicAuth(req)

	resp, err := f.client.Do(req)
	if err != nil {
		log.Println(err)
	}

	respDecoder := json.NewDecoder(resp.Body)

	var hostList []struct {
		Host Host `json:"host"`
	}

	err = respDecoder.Decode(&hostList)
	if err != nil {
		log.Println(err)
		return nil
	}

	hosts := make([]Host, len(hostList))

	for i, host := range hostList {
		hosts[i] = host.Host
	}

	return hosts
}
