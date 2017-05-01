//Copyright 2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ottenwbe/golook/broker/api"
	"github.com/ottenwbe/golook/broker/models"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

var (
	Host string
)

func GetFiles(searchString string) (string, error) {
	return get(fmt.Sprintf("%s/%s", api.FileEndpoint, searchString))
}

func ReportFiles(file string) (string, error) {

	logrus.Info("Report for: " + fmt.Sprintf("%s%s/", Host, api.FileEndpoint))

	c := http.Client{}

	b, err := json.Marshal(models.FileReport{Path: file})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", Host, api.FileEndpoint), bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error retrieving result from server: %d", resp.StatusCode)
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func GetInfo() (string, error) {
	return get(api.InfoEndpoint)
}

func GetConfig() (string, error) {
	return get(api.ConfigEndpoint)
}

func get(ep string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", Host, ep))
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Error retrieving result from server: %d %s", resp.StatusCode, resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
