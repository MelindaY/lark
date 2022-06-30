// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// FinishUploadDriveMedia 触发完成上传。
//
// 该接口不支持太高的并发，且调用频率上限为5QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/media/upload_finish
func (r *DriveService) FinishUploadDriveMedia(ctx context.Context, request *FinishUploadDriveMediaReq, options ...MethodOptionFunc) (*FinishUploadDriveMediaResp, *Response, error) {
	if r.cli.mock.mockDriveFinishUploadDriveMedia != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#FinishUploadDriveMedia mock enable")
		return r.cli.mock.mockDriveFinishUploadDriveMedia(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "FinishUploadDriveMedia",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/medias/upload_finish",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(finishUploadDriveMediaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveFinishUploadDriveMedia mock DriveFinishUploadDriveMedia method
func (r *Mock) MockDriveFinishUploadDriveMedia(f func(ctx context.Context, request *FinishUploadDriveMediaReq, options ...MethodOptionFunc) (*FinishUploadDriveMediaResp, *Response, error)) {
	r.mockDriveFinishUploadDriveMedia = f
}

// UnMockDriveFinishUploadDriveMedia un-mock DriveFinishUploadDriveMedia method
func (r *Mock) UnMockDriveFinishUploadDriveMedia() {
	r.mockDriveFinishUploadDriveMedia = nil
}

// FinishUploadDriveMediaReq ...
type FinishUploadDriveMediaReq struct {
	UploadID string `json:"upload_id,omitempty"` // 分片上传事务ID, 示例值："1233456"
	BlockNum int64  `json:"block_num,omitempty"` // 分片数量, 示例值：1
}

// finishUploadDriveMediaResp ...
type finishUploadDriveMediaResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *FinishUploadDriveMediaResp `json:"data,omitempty"`
}

// FinishUploadDriveMediaResp ...
type FinishUploadDriveMediaResp struct {
	FileToken string `json:"file_token,omitempty"` // 新创建文件的 token
}