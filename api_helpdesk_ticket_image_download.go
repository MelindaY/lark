// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
	"io"
)

// DownloadTicketImage 该接口用于获取服务台工单消息图象。仅支持自建应用。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket/ticket_image
func (r *HelpdeskService) DownloadTicketImage(ctx context.Context, request *DownloadTicketImageReq, options ...MethodOptionFunc) (*DownloadTicketImageResp, *Response, error) {
	if r.cli.mock.mockHelpdeskDownloadTicketImage != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#DownloadTicketImage mock enable")
		return r.cli.mock.mockHelpdeskDownloadTicketImage(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Helpdesk#DownloadTicketImage call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#DownloadTicketImage request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_images",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedHelpdeskAuth:      true,
	}
	resp := new(downloadTicketImageResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#DownloadTicketImage GET https://open.feishu.cn/open-apis/helpdesk/v1/ticket_images failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#DownloadTicketImage GET https://open.feishu.cn/open-apis/helpdesk/v1/ticket_images failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "DownloadTicketImage", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#DownloadTicketImage request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskDownloadTicketImage(f func(ctx context.Context, request *DownloadTicketImageReq, options ...MethodOptionFunc) (*DownloadTicketImageResp, *Response, error)) {
	r.mockHelpdeskDownloadTicketImage = f
}

func (r *Mock) UnMockHelpdeskDownloadTicketImage() {
	r.mockHelpdeskDownloadTicketImage = nil
}

type DownloadTicketImageReq struct {
	TicketID string `query:"ticket_id" json:"-"` // 工单ID, 示例值："12345"
	MsgID    string `query:"msg_id" json:"-"`    // 消息ID, 示例值："12345"
	Index    *int   `query:"index" json:"-"`     // index，当消息类型为post时，需指定图片index，index从0开始。当消息类型为img时，无需index, 示例值：0
}

type downloadTicketImageResp struct {
	IsFile bool                     `json:"is_file,omitempty"`
	Code   int                      `json:"code,omitempty"`
	Msg    string                   `json:"msg,omitempty"`
	Data   *DownloadTicketImageResp `json:"data,omitempty"`
}

func (r *downloadTicketImageResp) IsFileType() bool {
	return r.IsFile
}

func (r *downloadTicketImageResp) SetFile(file io.Reader) {
	r.Data = &DownloadTicketImageResp{
		File: file,
	}
}

type DownloadTicketImageResp struct {
	File io.Reader `json:"file,omitempty"`
}
