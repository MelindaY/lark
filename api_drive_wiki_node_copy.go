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

// CopyWikiNode 此接口用于创建节点副本到指定地点。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/wiki-v2/space-node/copy
func (r *DriveService) CopyWikiNode(ctx context.Context, request *CopyWikiNodeReq, options ...MethodOptionFunc) (*CopyWikiNodeResp, *Response, error) {
	if r.cli.mock.mockDriveCopyWikiNode != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#CopyWikiNode mock enable")
		return r.cli.mock.mockDriveCopyWikiNode(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "CopyWikiNode",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/wiki/v2/spaces/:space_id/nodes/:node_token/copy",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(copyWikiNodeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockDriveCopyWikiNode mock DriveCopyWikiNode method
func (r *Mock) MockDriveCopyWikiNode(f func(ctx context.Context, request *CopyWikiNodeReq, options ...MethodOptionFunc) (*CopyWikiNodeResp, *Response, error)) {
	r.mockDriveCopyWikiNode = f
}

// UnMockDriveCopyWikiNode un-mock DriveCopyWikiNode method
func (r *Mock) UnMockDriveCopyWikiNode() {
	r.mockDriveCopyWikiNode = nil
}

// CopyWikiNodeReq ...
type CopyWikiNodeReq struct {
	SpaceID           string  `path:"space_id" json:"-"`             // 知识空间id, 示例值："6946843325487912356"
	NodeToken         string  `path:"node_token" json:"-"`           // 节点token, 示例值："wikcnKQ1k3pcuo5uSK4t8Vabcef"
	TargetParentToken *string `json:"target_parent_token,omitempty"` // 目标父节点token, 示例值："wikcnKQ1k3pcuo5uSK4t8Vabcef"
	TargetSpaceID     *string `json:"target_space_id,omitempty"`     // 目标知识空间id, 示例值："6946843325487912356"
	Title             *string `json:"title,omitempty"`               // 复制后的新标题。如果填空，则新标题为空。如果不填，则使用原节点标题。, 示例值："新标题"
}

// copyWikiNodeResp ...
type copyWikiNodeResp struct {
	Code int64             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string            `json:"msg,omitempty"`  // 错误描述
	Data *CopyWikiNodeResp `json:"data,omitempty"`
}

// CopyWikiNodeResp ...
type CopyWikiNodeResp struct {
	Node *CopyWikiNodeRespNode `json:"node,omitempty"` // copy后的节点
}

// CopyWikiNodeRespNode ...
type CopyWikiNodeRespNode struct {
	SpaceID         string `json:"space_id,omitempty"`          // 知识库id
	NodeToken       string `json:"node_token,omitempty"`        // 节点token
	ObjToken        string `json:"obj_token,omitempty"`         // 文档token，可以根据obj_type判断是属于doc、sheet还是mindnote的token(对于快捷方式，该字段是对应的实体的obj_token)
	ObjType         string `json:"obj_type,omitempty"`          // 文档类型，对于快捷方式，该字段是对应的实体的obj_type, 可选值有: ,<md-enum>,<md-enum-item key="doc" >doc</md-enum-item>,<md-enum-item key="sheet" >sheet</md-enum-item>,<md-enum-item key="mindnote" >mindnote</md-enum-item>,<md-enum-item key="bitable" >bitable</md-enum-item>,<md-enum-item key="file" >file</md-enum-item>,<md-enum-item key="docx" >docx</md-enum-item>,</md-enum>
	ParentNodeToken string `json:"parent_node_token,omitempty"` // 节点的父亲token。当节点为一级节点时，父亲token为空。
	NodeType        string `json:"node_type,omitempty"`         // 节点类型, 可选值有: ,<md-enum>,<md-enum-item key="origin" >实体</md-enum-item>,<md-enum-item key="shortcut" >快捷方式</md-enum-item>,</md-enum>
	OriginNodeToken string `json:"origin_node_token,omitempty"` // 快捷方式对应的实体node_token，当创建节点为快捷方式时，需要传该值
	OriginSpaceID   string `json:"origin_space_id,omitempty"`   // 快捷方式对应的实体所在的spaceid
	HasChild        bool   `json:"has_child,omitempty"`         // 是否有子节点
	Title           string `json:"title,omitempty"`             // 文档标题
	ObjCreateTime   string `json:"obj_create_time,omitempty"`   // 文档创建时间
	ObjEditTime     string `json:"obj_edit_time,omitempty"`     // 文档最近编辑时间
	NodeCreateTime  string `json:"node_create_time,omitempty"`  // 节点创建时间
	Creator         string `json:"creator,omitempty"`           // 节点创建者
	Owner           string `json:"owner,omitempty"`             // 节点所有者
}