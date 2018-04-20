// Copyright (c) 2018 Huawei Technologies Co., Ltd. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/astaxie/beego"
	c "github.com/opensds/opensds/client"
	"github.com/opensds/opensds/plugin/cindercompatibleapi/cindermodel"
)

func init() {
	beego.Router("/v3/volumes/:volumeId", &VolumePortal{},
		"get:GetVolume;delete:DeleteVolume;put:UpdateVolume")
	beego.Router("/v3/volumes/detail", &VolumePortal{},
		"get:ListVolumeDetail")
	beego.Router("/v3/volumes", &VolumePortal{},
		"post:CreateVolume;get:ListVolume")
	if false == IsFakeClient {
		client = NewFakeClient(&c.Config{Endpoint: TestEp})
	}
}

////////////////////////////////////////////////////////////////////////////////
//                            Tests for Volume                              //
////////////////////////////////////////////////////////////////////////////////
func TestGetVolume(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v3/volumes/bd5b12a8-a101-11e7-941e-d77981b584d8", nil)

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output cindermodel.ShowVolumeRespSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	expectedJSON := `{
    	"volume": {
		"id": "bd5b12a8-a101-11e7-941e-d77981b584d8",
        "name": "sample-volume",
		"description": "This is a sample volume for testing",
		"metadata":{},
		"size": 1
		}
	}`

	var expected cindermodel.ShowVolumeRespSpec
	json.Unmarshal([]byte(expectedJSON), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	expected.Volume.Attachments = make([]cindermodel.AttachmentOfVolumeResp, 0, 0)
	expected.Volume.Status = "available"

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}

func TestListVolume(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v3/volumes", nil)

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output cindermodel.ListVolumeRespSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	expectedJSON := `{"volumes":
		[
		{"id":"bd5b12a8-a101-11e7-941e-d77981b584d8",
		"metadata":{},
		 "name":"sample-volume"
		}
		]
	}`

	var expected cindermodel.ListVolumeRespSpec
	json.Unmarshal([]byte(expectedJSON), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}

func TestListVolumeDetail(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v3/volumes/detail", nil)

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output cindermodel.ListVolumeDetailRespSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	expectedJSON := `{"volumes":
		[
		{"id":"bd5b12a8-a101-11e7-941e-d77981b584d8",
			"size":1,"status":"available",
			"description":"This is a sample volume for testing",
			"metadata":{},
			"name":"sample-volume"
			}
			]
	}`

	var expected cindermodel.ListVolumeDetailRespSpec
	json.Unmarshal([]byte(expectedJSON), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	expected.Volumes[0].Attachments = make([]cindermodel.AttachmentOfVolumeResp, 0, 0)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}

func TestCreateVolume(t *testing.T) {
	RequestBodyStr := `{
    	"volume": {
        "name": "sample-volume",
		"description": "This is a sample volume for testing",
		"size": 1
		}
	}`

	var jsonStr = []byte(RequestBodyStr)
	r, _ := http.NewRequest("POST", "/v3/volumes", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output cindermodel.CreateVolumeRespSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	var expected cindermodel.CreateVolumeRespSpec
	json.Unmarshal([]byte(RequestBodyStr), &expected)

	if w.Code != StatusAccepted {
		t.Errorf("Expected %v, actual %v", StatusAccepted, w.Code)
	}

	expected.Volume.Attachments = make([]cindermodel.AttachmentOfVolumeResp, 0, 0)
	expected.Volume.Status = "available"
	expected.Volume.ID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected.Volume.Metadata = make(map[string]string)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}

}

func TestDeleteVolume(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/v3/volumes/bd5b12a8-a101-11e7-941e-d77981b584d8", nil)

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	if w.Code != 202 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}
}

func TestUpdateVolume(t *testing.T) {
	RequestBodyStr := `{
    	"volume": {
        "name": "sample-volume",
		"multiattach": false,
		"description": "This is a sample volume for testing"
		}
	}`

	var jsonStr = []byte(RequestBodyStr)
	r, _ := http.NewRequest("PUT", "/v3/volumes/bd5b12a8-a101-11e7-941e-d77981b584d8", bytes.NewBuffer(jsonStr))

	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	var output cindermodel.UpdateVolumeRespSpec
	json.Unmarshal(w.Body.Bytes(), &output)

	var expected cindermodel.UpdateVolumeRespSpec
	json.Unmarshal([]byte(RequestBodyStr), &expected)

	if w.Code != 200 {
		t.Errorf("Expected 200, actual %v", w.Code)
	}

	expected.Volume.Attachments = make([]cindermodel.AttachmentOfVolumeResp, 0, 0)
	expected.Volume.Status = "available"
	expected.Volume.ID = "bd5b12a8-a101-11e7-941e-d77981b584d8"
	expected.Volume.Size = 1
	expected.Volume.Metadata = make(map[string]string)
	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected %v, actual %v", expected, output)
	}
}
