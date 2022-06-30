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

// GetDriveFileList 获取用户云空间中指定文件夹下的文件清单。清单类型包括文件、各种在线文档（文档、电子表格、多维表格、思维笔记）、文件夹和快捷方式。该接口支持分页，但是不会递归的获取子文件夹的清单。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/list
func (r *DriveService) GetDriveFileList(ctx context.Context, request *GetDriveFileListReq, options ...MethodOptionFunc) (*GetDriveFileListResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFileList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFileList mock enable")
		return r.cli.mock.mockDriveGetDriveFileList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "GetDriveFileList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/drive/v1/files",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getDriveFileListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveGetDriveFileList mock DriveGetDriveFileList method
func (r *Mock) MockDriveGetDriveFileList(f func(ctx context.Context, request *GetDriveFileListReq, options ...MethodOptionFunc) (*GetDriveFileListResp, *Response, error)) {
	r.mockDriveGetDriveFileList = f
}

// UnMockDriveGetDriveFileList un-mock DriveGetDriveFileList method
func (r *Mock) UnMockDriveGetDriveFileList() {
	r.mockDriveGetDriveFileList = nil
}

// GetDriveFileListReq ...
type GetDriveFileListReq struct {
	PageSize    *int64  `query:"page_size" json:"-"`    // 分页大小, 示例值：10, 最大值：`200`
	PageToken   *string `query:"page_token" json:"-"`   // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："MTY1NTA3MTA1OXw3MTA4NDc2MDc1NzkyOTI0NjczfDE"
	FolderToken *string `query:"folder_token" json:"-"` // 文件夹的token, 示例值："fldbcRho46N6MQ3mJkOAuPabcef"
}

// getDriveFileListResp ...
type getDriveFileListResp struct {
	Code int64                 `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                `json:"msg,omitempty"`  // 错误描述
	Data *GetDriveFileListResp `json:"data,omitempty"`
}

// GetDriveFileListResp ...
type GetDriveFileListResp struct {
	Files         []*GetDriveFileListRespFile `json:"files,omitempty"`           // 文件夹清单列表
	NextPageToken string                      `json:"next_page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回下一次遍历的page_token，否则则不返回
	HasMore       bool                        `json:"has_more,omitempty"`        // 是否还有更多项
}

// GetDriveFileListRespFile ...
type GetDriveFileListRespFile struct {
	Token        string                                `json:"token,omitempty"`         // 文件标识
	Name         string                                `json:"name,omitempty"`          // 文件名
	Type         string                                `json:"type,omitempty"`          // 文件类型
	ParentToken  string                                `json:"parent_token,omitempty"`  // 父文件夹标识
	URL          string                                `json:"url,omitempty"`           // 在浏览器中查看的链接
	ShortcutInfo *GetDriveFileListRespFileShortcutInfo `json:"shortcut_info,omitempty"` // 快捷方式文件信息
}

// GetDriveFileListRespFileShortcutInfo ...
type GetDriveFileListRespFileShortcutInfo struct {
	TargetType  string `json:"target_type,omitempty"`  // 快捷方式指向的原文件类型
	TargetToken string `json:"target_token,omitempty"` // 快捷方式指向的原文件token
}