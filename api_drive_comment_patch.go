// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateCommentPatch 解决或恢复云文档中的评论。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file-comment/patch
func (r *DriveService) UpdateCommentPatch(ctx context.Context, request *UpdateCommentPatchReq, options ...MethodOptionFunc) (*UpdateCommentPatchResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateCommentPatch != nil {
		r.cli.logDebug(ctx, "[lark] Drive#UpdateCommentPatch mock enable")
		return r.cli.mock.mockDriveUpdateCommentPatch(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Drive#UpdateCommentPatch call api")
	r.cli.logDebug(ctx, "[lark] Drive#UpdateCommentPatch request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateCommentPatchResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Drive#UpdateCommentPatch PATCH https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Drive#UpdateCommentPatch PATCH https://open.feishu.cn/open-apis/drive/v1/files/:file_token/comments/:comment_id failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "UpdateCommentPatch", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Drive#UpdateCommentPatch request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveUpdateCommentPatch(f func(ctx context.Context, request *UpdateCommentPatchReq, options ...MethodOptionFunc) (*UpdateCommentPatchResp, *Response, error)) {
	r.mockDriveUpdateCommentPatch = f
}

func (r *Mock) UnMockDriveUpdateCommentPatch() {
	r.mockDriveUpdateCommentPatch = nil
}

type UpdateCommentPatchReq struct {
	FileType  FileType `query:"file_type" json:"-"` // 文档类型, 示例值："doc", 可选值有: `doc`：文档, `sheet`：表格, `file`：文件
	FileToken string   `path:"file_token" json:"-"` // 文档token, 示例值："doccnGp4UK1UskrOEJwBXd3****"
	CommentID string   `path:"comment_id" json:"-"` // 评论ID, 示例值："6916106822734578184"
	IsSolved  bool     `json:"is_solved,omitempty"` // 评论解决标志, 示例值：true
}

type updateCommentPatchResp struct {
	Code int                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *UpdateCommentPatchResp `json:"data,omitempty"`
}

type UpdateCommentPatchResp struct{}
