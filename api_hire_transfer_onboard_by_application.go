// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// MakeHireTransferOnboardByApplication 根据投递 ID 操作候选人入职并创建员工，投递须处于「待入职」阶段
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMzM1YjLzMTN24yMzUjN/hire-v1/application/transfer_onboard
func (r *HireService) MakeHireTransferOnboardByApplication(ctx context.Context, request *MakeHireTransferOnboardByApplicationReq, options ...MethodOptionFunc) (*MakeHireTransferOnboardByApplicationResp, *Response, error) {
	if r.cli.mock.mockHireMakeHireTransferOnboardByApplication != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Hire#MakeHireTransferOnboardByApplication mock enable")
		return r.cli.mock.mockHireMakeHireTransferOnboardByApplication(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Hire",
		API:                   "MakeHireTransferOnboardByApplication",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/hire/v1/applications/:application_id/transfer_onboard",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(makeHireTransferOnboardByApplicationResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHireMakeHireTransferOnboardByApplication(f func(ctx context.Context, request *MakeHireTransferOnboardByApplicationReq, options ...MethodOptionFunc) (*MakeHireTransferOnboardByApplicationResp, *Response, error)) {
	r.mockHireMakeHireTransferOnboardByApplication = f
}

func (r *Mock) UnMockHireMakeHireTransferOnboardByApplication() {
	r.mockHireMakeHireTransferOnboardByApplication = nil
}

type MakeHireTransferOnboardByApplicationReq struct {
	ApplicationID          string `path:"application_id" json:"-"`            // 投递ID, 示例值："12312312312"
	ActualOnboardTime      *int64 `json:"actual_onboard_time,omitempty"`      // 实际入职时间, 示例值：1616428800000
	ExpectedConversionTime *int64 `json:"expected_conversion_time,omitempty"` // 预期转正时间, 示例值：1616428800000
}

type makeHireTransferOnboardByApplicationResp struct {
	Code int64                                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                    `json:"msg,omitempty"`  // 错误描述
	Data *MakeHireTransferOnboardByApplicationResp `json:"data,omitempty"` //
}

type MakeHireTransferOnboardByApplicationResp struct {
	ID                     string `json:"id,omitempty"`                       // 员工ID
	ApplicationID          string `json:"application_id,omitempty"`           // 投递ID
	OnboardStatus          int64  `json:"onboard_status,omitempty"`           // 入职状态, 可选值有: `1`：已入职, `2`：已离职
	ConversionStatus       int64  `json:"conversion_status,omitempty"`        // 转正状态, 可选值有: `1`：未转正, `2`：已转正
	OnboardTime            int64  `json:"onboard_time,omitempty"`             // 实际入职时间
	ExpectedConversionTime int64  `json:"expected_conversion_time,omitempty"` // 预期转正时间
	ActualConversionTime   int64  `json:"actual_conversion_time,omitempty"`   // 实际转正时间
	OverboardTime          int64  `json:"overboard_time,omitempty"`           // 离职时间
	OverboardNote          string `json:"overboard_note,omitempty"`           // 离职原因
}
