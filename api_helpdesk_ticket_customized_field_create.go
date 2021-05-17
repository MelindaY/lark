// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateTicketCustomizedField
//
// 该接口用于创建自定义字段
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/create-ticket-customized-field
func (r *HelpdeskService) CreateTicketCustomizedField(ctx context.Context, request *CreateTicketCustomizedFieldReq, options ...MethodOptionFunc) (*CreateTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskCreateTicketCustomizedField != nil {
		r.cli.logDebug(ctx, "[lark] Helpdesk#CreateTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskCreateTicketCustomizedField(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Helpdesk#CreateTicketCustomizedField call api")
	r.cli.logDebug(ctx, "[lark] Helpdesk#CreateTicketCustomizedField request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:              "POST",
		URL:                 "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(createTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Helpdesk#CreateTicketCustomizedField POST https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Helpdesk#CreateTicketCustomizedField POST https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Helpdesk", "CreateTicketCustomizedField", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Helpdesk#CreateTicketCustomizedField request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockHelpdeskCreateTicketCustomizedField(f func(ctx context.Context, request *CreateTicketCustomizedFieldReq, options ...MethodOptionFunc) (*CreateTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskCreateTicketCustomizedField = f
}

func (r *Mock) UnMockHelpdeskCreateTicketCustomizedField() {
	r.mockHelpdeskCreateTicketCustomizedField = nil
}

type CreateTicketCustomizedFieldReq struct {
	HelpdeskID            string                  `json:"helpdesk_id,omitempty"`             // 服务台ID, 示例值："1542164574896126"
	KeyName               string                  `json:"key_name,omitempty"`                // 键名, 示例值："test dropdown"
	DisplayName           string                  `json:"display_name,omitempty"`            // 名称, 示例值："test dropdown"
	Position              string                  `json:"position,omitempty"`                // 字段在列表后台管理列表中的位置, 示例值："3"
	FieldType             string                  `json:"field_type,omitempty"`              // 类型, 示例值："dropdown"
	Description           string                  `json:"description,omitempty"`             // 描述, 示例值："下拉示例"
	Visible               bool                    `json:"visible,omitempty"`                 // 是否可见, 示例值：true
	Editable              bool                    `json:"editable,omitempty"`                // 是否可以修改, 示例值：true
	Required              bool                    `json:"required,omitempty"`                // 是否必填, 示例值：false
	DropdownOptions       *HelpdeskDropdownOption `json:"dropdown_options,omitempty"`        // 下拉列表选项
	DropdownAllowMultiple *bool                   `json:"dropdown_allow_multiple,omitempty"` // 是否支持多选，仅在字段类型是dropdown的时候有效, 示例值：true
}

type createTicketCustomizedFieldResp struct {
	Code int                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *CreateTicketCustomizedFieldResp `json:"data,omitempty"`
}

type CreateTicketCustomizedFieldResp struct{}
