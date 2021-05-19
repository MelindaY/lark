// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetSheetValue
//
// 该接口用于根据 spreadsheetToken 和 range 读取表格单个范围的值，返回数据限制为10M。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/ugTMzUjL4EzM14COxMTN
func (r *DriveService) GetSheetValue(ctx context.Context, request *GetSheetValueReq, options ...MethodOptionFunc) (*GetSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveGetSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetValue mock enable")
		return r.cli.mock.mockDriveGetSheetValue(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#GetSheetValue call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetValue request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(getSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#GetSheetValue GET https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#GetSheetValue GET https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values/:range failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "GetSheetValue", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetSheetValue success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveGetSheetValue(f func(ctx context.Context, request *GetSheetValueReq, options ...MethodOptionFunc) (*GetSheetValueResp, *Response, error)) {
	r.mockDriveGetSheetValue = f
}

func (r *Mock) UnMockDriveGetSheetValue() {
	r.mockDriveGetSheetValue = nil
}

type GetSheetValueReq struct {
	ValueRenderOption    *string `query:"valueRenderOption" json:"-"`    // valueRenderOption=ToString 可返回 toString 后的值；valueRenderOption=FormattedValue可返回格式化后的字符串；
	DateTimeRenderOption *string `query:"dateTimeRenderOption" json:"-"` // dateTimeRenderOption=FormattedString 会将时间日期按照其格式进行格式化，但不会对数字进行格式化，返回格式化后的字符串。
	SpreadsheetToken     string  `path:"spreadsheetToken" json:"-"`      // spreadsheet 的 token，详见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	Range                string  `path:"range" json:"-"`                 // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见 ⁣[对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)  的第 5 项
}

type getSheetValueResp struct {
	Code int64              `json:"code,omitempty"`
	Msg  string             `json:"msg,omitempty"`
	Data *GetSheetValueResp `json:"data,omitempty"`
}

type GetSheetValueResp struct {
	Range    string        `json:"range,omitempty"`    // 查询范围
	Values   []interface{} `json:"values,omitempty"`   // 查询获得的值
	Revision int64         `json:"revision,omitempty"` // sheet 的版本号
}
