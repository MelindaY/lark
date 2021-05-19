// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchSetSheetValue 该接口用于根据 spreadsheetToken 和 range 向多个范围写入数据，若范围内有数据，将被更新覆盖；单次写入不超过5000行，100列，每个格子大小为0.5M。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEjMzUjLxIzM14SMyMTN
func (r *DriveService) BatchSetSheetValue(ctx context.Context, request *BatchSetSheetValueReq, options ...MethodOptionFunc) (*BatchSetSheetValueResp, *Response, error) {
	if r.cli.mock.mockDriveBatchSetSheetValue != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetValue mock enable")
		return r.cli.mock.mockDriveBatchSetSheetValue(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#BatchSetSheetValue call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetValue request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(batchSetSheetValueResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#BatchSetSheetValue POST https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#BatchSetSheetValue POST https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/values_batch_update failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "BatchSetSheetValue", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetValue success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveBatchSetSheetValue(f func(ctx context.Context, request *BatchSetSheetValueReq, options ...MethodOptionFunc) (*BatchSetSheetValueResp, *Response, error)) {
	r.mockDriveBatchSetSheetValue = f
}

func (r *Mock) UnMockDriveBatchSetSheetValue() {
	r.mockDriveBatchSetSheetValue = nil
}

type BatchSetSheetValueReq struct {
	SpreadsheetToken string                            `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	ValueRanges      *BatchSetSheetValueReqValueRanges `json:"valueRanges,omitempty"`     // 需要更新的多个范围
}

type BatchSetSheetValueReqValueRanges struct {
	Range  string           `json:"range,omitempty"`  // 更新范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见 ⁣[对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)  的第 5 项
	Values [][]SheetContent `json:"values,omitempty"` // 需要写入的值，如要写入公式、超链接、emial、@人等，可详看附录[sheet 支持写入数据类型](/ssl:ttdoc/ukTMukTMukTM/ugjN1UjL4YTN14CO2UTN)
}

type batchSetSheetValueResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *BatchSetSheetValueResp `json:"data,omitempty"`
}

type BatchSetSheetValueResp struct {
	Responses        *BatchSetSheetValueRespResponses `json:"responses,omitempty"`        // 响应
	Revision         int64                            `json:"revision,omitempty"`         // sheet 的版本号
	SpreadsheetToken string                           `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
}

type BatchSetSheetValueRespResponses struct {
	SpreadsheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 写入的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 写入的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 写入的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 写入的单元格总数
}
