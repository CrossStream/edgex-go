/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package metadata

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/core/metadata/interfaces"
	"github.com/edgexfoundry/edgex-go/internal/core/metadata/interfaces/mocks"
	"github.com/edgexfoundry/edgex-go/internal/pkg/config"
	"github.com/edgexfoundry/go-mod-core-contracts/clients"
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/pkg/errors"
)

func TestGetAllDevices(t *testing.T) {
	Configuration = &ConfigurationStruct{Service: config.ServiceInfo{MaxResultCount: 1}}
	req, err := http.NewRequest(http.MethodGet, clients.ApiBase+"/"+DEVICE, nil)
	if err != nil {
		t.Error(err)
		return
	}

	tests := []struct {
		name           string
		dbMock         interfaces.DBClient
		expectedStatus int
	}{
		{"OK", createGetDeviceLoaderMock(1), http.StatusOK},
		{"MaxExceeded", createGetDeviceLoaderMock(2), http.StatusRequestEntityTooLarge},
		{"Unexpected", createGetDeviceLoaderMockFail(), http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbClient = tt.dbMock
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(restGetAllDevices)

			handler.ServeHTTP(rr, req)
			response := rr.Result()
			if response.StatusCode != tt.expectedStatus {
				t.Errorf("status code mismatch -- expected %v got %v", tt.expectedStatus, response.StatusCode)
				return
			}
		})
	}
	Configuration = &ConfigurationStruct{}
}

func createGetDeviceLoaderMock(howMany int) interfaces.DBClient {
	devices := []contract.Device{}
	for i := 0; i < howMany; i++ {
		devices = append(devices, contract.Device{Name: "Some Device"})
	}

	dbMock := &mocks.DBClient{}
	dbMock.On("GetAllDevices").Return(devices, nil)
	return dbMock
}

func createGetDeviceLoaderMockFail() interfaces.DBClient {
	dbMock := &mocks.DBClient{}
	dbMock.On("GetAllDevices").Return(nil, errors.New("unexpected error"))
	return dbMock
}
