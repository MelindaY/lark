// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdatePublicPermission 该接口用于根据 filetoken 更新文档的公共设置。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ukTM3UjL5EzN14SOxcTN
func (r *DriveService) UpdatePublicPermission(ctx context.Context, request *UpdatePublicPermissionReq, options ...MethodOptionFunc) (*UpdatePublicPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveUpdatePublicPermission != nil {
		r.cli.logDebug(ctx, "[lark] Drive#UpdatePublicPermission mock enable")
		return r.cli.mock.mockDriveUpdatePublicPermission(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#UpdatePublicPermission call api")
	r.cli.logDebug(ctx, "[lark] Drive#UpdatePublicPermission request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/public/update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updatePublicPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#UpdatePublicPermission POST https://open.feishu.cn/open-apis/drive/permission/public/update failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#UpdatePublicPermission POST https://open.feishu.cn/open-apis/drive/permission/public/update failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "UpdatePublicPermission", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#UpdatePublicPermission request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveUpdatePublicPermission(f func(ctx context.Context, request *UpdatePublicPermissionReq, options ...MethodOptionFunc) (*UpdatePublicPermissionResp, *Response, error)) {
	r.mockDriveUpdatePublicPermission = f
}

func (r *Mock) UnMockDriveUpdatePublicPermission() {
	r.mockDriveUpdatePublicPermission = nil
}

type UpdatePublicPermissionReq struct {
	Token                 string  `json:"token,omitempty"`                    // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type                  string  `json:"type,omitempty"`                     // 文档类型  "doc"  or  "sheet"
	CopyPrintExportStatus *bool   `json:"copy_print_export_status,omitempty"` // 可创建副本/打印/导出/复制设置（不传则保持原值）：<br>true - 所有可访问此文档的用户<br>false - 有编辑权限的用户
	Comment               *bool   `json:"comment,omitempty"`                  // 可评论设置（不传则保持原值）：<br>true - 所有可访问此文档的用户<br>false - 有编辑权限的用户
	TenantShareable       *bool   `json:"tenant_shareable,omitempty"`         // 租户内用户是否有共享权限（不传则保持原值）
	LinkShareEntity       *string `json:"link_share_entity,omitempty"`        // 链接共享（不传则保持原值）：<br>"tenant_readable" - 组织内获得链接的人可阅读<br>"tenant_editable" - 组织内获得链接的人可编辑<br>"anyone_readable" - 获得链接的任何人可阅读<br>"anyone_editable" - 获得链接的任何人可编辑
	ExternalAccess        *bool   `json:"external_access,omitempty"`          // 是否允许分享到租户外开关（不传则保持原值）
	InviteExternal        *bool   `json:"invite_external,omitempty"`          // 非owner是否允许邀请外部人（不传则保持原值）
}

type updatePublicPermissionResp struct {
	Code int                         `json:"code,omitempty"`
	Msg  string                      `json:"msg,omitempty"`
	Data *UpdatePublicPermissionResp `json:"data,omitempty"`
}

type UpdatePublicPermissionResp struct {
	IsSuccess bool `json:"is_success,omitempty"` // 是否成功
}
