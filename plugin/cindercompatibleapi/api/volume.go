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

/*
This module implements a entry into the OpenSDS northbound service.

*/

package api

import (
	"encoding/json"
	"fmt"

	log "github.com/golang/glog"
	OpenSDSAPI "github.com/opensds/opensds/pkg/api"
	"github.com/opensds/opensds/pkg/api/policy"
	"github.com/opensds/opensds/plugin/cindercompatibleapi/cindermodel"
	"github.com/opensds/opensds/plugin/cindercompatibleapi/converter"

	"github.com/opensds/opensds/pkg/model"
)

// VolumePortal ...
type VolumePortal struct {
	OpenSDSAPI.BasePortal
}

// ListVolumeDetail ...
func (portal *VolumePortal) ListVolumeDetail() {
	if !policy.Authorize(portal.Ctx, "volume:list_detail") {
		return
	}

	volumes, err := client.ListVolumes()
	if err != nil {
		reason := fmt.Sprintf("List accessible volumes with details failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result := converter.ListVolumeDetailResp(volumes)
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("List accessible volumes with details, marshal result failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorInternalServer)
		portal.Ctx.Output.Body(model.ErrorInternalServerStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusOK)
	portal.Ctx.Output.Body(body)
	return
}

// CreateVolume ...
func (portal *VolumePortal) CreateVolume() {
	if !policy.Authorize(portal.Ctx, "volume:create") {
		return
	}
	var cinderReq = cindermodel.CreateVolumeReqSpec{}

	if err := json.NewDecoder(portal.Ctx.Request.Body).Decode(&cinderReq); err != nil {
		reason := fmt.Sprintf("Create a volume, parse request body failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	volume, err := converter.CreateVolumeReq(&cinderReq)
	if err != nil {
		reason := fmt.Sprintf("Create a volume failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	volume, err = client.CreateVolume(volume)
	if err != nil {
		reason := fmt.Sprintf("Create a volume failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result := converter.CreateVolumeResp(volume)
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Create a volume, marshal result failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorInternalServer)
		portal.Ctx.Output.Body(model.ErrorInternalServerStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusAccepted)
	portal.Ctx.Output.Body(body)
	return
}

// ListVolume ...
func (portal *VolumePortal) ListVolume() {
	if !policy.Authorize(portal.Ctx, "volume:list") {
		return
	}

	volumes, err := client.ListVolumes()
	if err != nil {
		reason := fmt.Sprintf("List accessible volumes failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result := converter.ListVolumeResp(volumes)
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("List accessible volumes, marshal result failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorInternalServer)
		portal.Ctx.Output.Body(model.ErrorInternalServerStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusOK)
	portal.Ctx.Output.Body(body)
	return
}

// GetVolume ...
func (portal *VolumePortal) GetVolume() {
	if !policy.Authorize(portal.Ctx, "volume:get") {
		return
	}

	id := portal.Ctx.Input.Param(":volumeId")
	volume, err := client.GetVolume(id)

	if err != nil {
		reason := fmt.Sprintf("Show a volume’s details failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result := converter.ShowVolumeResp(volume)
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Show a volume’s details, marshal result failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorInternalServer)
		portal.Ctx.Output.Body(model.ErrorInternalServerStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusOK)
	portal.Ctx.Output.Body(body)
	return
}

// UpdateVolume ...
func (portal *VolumePortal) UpdateVolume() {
	if !policy.Authorize(portal.Ctx, "volume:update") {
		return
	}

	id := portal.Ctx.Input.Param(":volumeId")
	var cinderReq = cindermodel.UpdateVolumeReqSpec{}

	if err := json.NewDecoder(portal.Ctx.Request.Body).Decode(&cinderReq); err != nil {
		reason := fmt.Sprintf("Update a volume, parse request body failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	volume, err := converter.UpdateVolumeReq(&cinderReq)
	if err != nil {
		reason := fmt.Sprintf("Update a volume failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	volume, err = client.UpdateVolume(id, volume)

	if err != nil {
		reason := fmt.Sprintf("Update a volume failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	result := converter.UpdateVolumeResp(volume)
	body, err := json.Marshal(result)
	if err != nil {
		reason := fmt.Sprintf("Update a volume, marshal result failed: %s", err.Error())
		portal.Ctx.Output.SetStatus(model.ErrorInternalServer)
		portal.Ctx.Output.Body(model.ErrorInternalServerStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusOK)
	portal.Ctx.Output.Body(body)
	return
}

// DeleteVolume ...
func (portal *VolumePortal) DeleteVolume() {
	if !policy.Authorize(portal.Ctx, "volume:delete") {
		return
	}

	id := portal.Ctx.Input.Param(":volumeId")
	volume := model.VolumeSpec{}

	err := client.DeleteVolume(id, &volume)

	if err != nil {
		reason := fmt.Sprintf("Delete a volume failed: %v", err)
		portal.Ctx.Output.SetStatus(model.ErrorBadRequest)
		portal.Ctx.Output.Body(model.ErrorBadRequestStatus(reason))
		log.Error(reason)
		return
	}

	portal.Ctx.Output.SetStatus(StatusAccepted)
	return
}
