// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CheckMemberPermission 该接口用于根据 filetoken 判断当前登录用户是否具有某权限。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uYzN3UjL2czN14iN3cTN
func (r *DriveAPI) CheckMemberPermission(ctx context.Context, request *CheckMemberPermissionReq, options ...MethodOptionFunc) (*CheckMemberPermissionResp, *Response, error) {
	if r.cli.mock.mockDriveCheckMemberPermission != nil {
		r.cli.logDebug(ctx, "[lark] Drive#CheckMemberPermission mock enable")
		return r.cli.mock.mockDriveCheckMemberPermission(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#CheckMemberPermission call api")
	r.cli.logDebug(ctx, "[lark] Drive#CheckMemberPermission request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/member/permitted",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(checkMemberPermissionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#CheckMemberPermission POST https://open.feishu.cn/open-apis/drive/permission/member/permitted failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#CheckMemberPermission POST https://open.feishu.cn/open-apis/drive/permission/member/permitted failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "CheckMemberPermission", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#CheckMemberPermission request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveCheckMemberPermission(f func(ctx context.Context, request *CheckMemberPermissionReq, options ...MethodOptionFunc) (*CheckMemberPermissionResp, *Response, error)) {
	r.mockDriveCheckMemberPermission = f
}

func (r *Mock) UnMockDriveCheckMemberPermission() {
	r.mockDriveCheckMemberPermission = nil
}

type CheckMemberPermissionReq struct {
	Token string `json:"token,omitempty"` // 文件的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type  string `json:"type,omitempty"`  // 文档类型  "doc"  or  "sheet" or "file"
	Perm  string `json:"perm,omitempty"`  // 权限，"view" or "edit" or "share"
}

type checkMemberPermissionResp struct {
	Code int                        `json:"code,omitempty"`
	Msg  string                     `json:"msg,omitempty"`
	Data *CheckMemberPermissionResp `json:"data,omitempty"`
}

type CheckMemberPermissionResp struct {
	IsPermitted bool `json:"is_permitted,omitempty"` // 是否具有指定权限
}
