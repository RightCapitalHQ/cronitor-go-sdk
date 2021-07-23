package cronitor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const apiEndpoint = "https://cronitor.io/api"
const apiVersion = "2020-09-01"

type Cronitor struct {
	ApiKey string
}

func (c Cronitor) PutMonitors(monitors []Monitor) error {
	postData, err := json.Marshal(MonitorRequest{
		Monitors: monitors,
	})

	if err != nil {
		return err
	}

	request, err := http.NewRequest(http.MethodPut, apiEndpoint+"/monitors", bytes.NewReader(postData))

	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := c.sendHttpRequest(request)

	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to put monitor. response: %s", respBody)
	}

	return nil
}

func (c Cronitor) DeleteMonitor(key string) error {
	request, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/monitors/%s", apiEndpoint, key), nil)
	if err != nil {
		return err
	}

	resp, err := c.sendHttpRequest(request)

	if err != nil {
		return err
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return fmt.Errorf("failed to delete monitor: %s. response: %s", key, respBody)
	}

	return nil
}

func (c Cronitor) GetMonitor(key string) (*Monitor, error) {
	return nil, nil
}

func (c Cronitor) sendHttpRequest(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(c.ApiKey, "")
	request.Header.Set("Cronitor-Version", apiVersion)

	return http.DefaultClient.Do(request)
}
