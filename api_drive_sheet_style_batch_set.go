// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// BatchSetSheetStyle
//
// 该接口用于根据 spreadsheetToken 、range和样式信息 批量更新单元格样式；单次写入不超过5000行，100列。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAzMzUjLwMzM14CMzMTN
func (r *DriveService) BatchSetSheetStyle(ctx context.Context, request *BatchSetSheetStyleReq, options ...MethodOptionFunc) (*BatchSetSheetStyleResp, *Response, error) {
	if r.cli.mock.mockDriveBatchSetSheetStyle != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetStyle mock enable")
		return r.cli.mock.mockDriveBatchSetSheetStyle(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Drive#BatchSetSheetStyle call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetStyle request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "PUT",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,

		NeedUserAccessToken: true,
	}
	resp := new(batchSetSheetStyleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#BatchSetSheetStyle PUT https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Drive#BatchSetSheetStyle PUT https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/styles_batch_update failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Drive", "BatchSetSheetStyle", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Drive#BatchSetSheetStyle success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockDriveBatchSetSheetStyle(f func(ctx context.Context, request *BatchSetSheetStyleReq, options ...MethodOptionFunc) (*BatchSetSheetStyleResp, *Response, error)) {
	r.mockDriveBatchSetSheetStyle = f
}

func (r *Mock) UnMockDriveBatchSetSheetStyle() {
	r.mockDriveBatchSetSheetStyle = nil
}

type BatchSetSheetStyleReq struct {
	SpreadsheetToken string                     `path:"spreadsheetToken" json:"-"` // spreadsheet 的 token，获取方式见 [对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 4 项
	Data             *BatchSetSheetStyleReqData `json:"data,omitempty"`            // 请求数据
}

type BatchSetSheetStyleReqData struct {
	Ranges []string                        `json:"ranges,omitempty"` // 查询范围，包含 sheetId 与单元格范围两部分，目前支持四种索引方式，详见 ⁣[对接前说明](/ssl:ttdoc/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)  的第 5 项
	Style  *BatchSetSheetStyleReqDataStyle `json:"style,omitempty"`  // 需要更新的样式
}

type BatchSetSheetStyleReqDataStyle struct {
	Font           *BatchSetSheetStyleReqDataStyleFont `json:"font,omitempty"`           // 字体相关样式
	TextDecoration *int64                              `json:"textDecoration,omitempty"` // 文本装饰 ，0 默认，1 下划线，2 删除线 ，3 下划线和删除线
	Formatter      *string                             `json:"formatter,omitempty"`      // 数字格式，详见附录 [sheet支持数字格式类型](/ssl:ttdoc/ukTMukTMukTM/uMjM2UjLzIjN14yMyYTN)
	HAlign         *int64                              `json:"hAlign,omitempty"`         // 水平对齐，0 左对齐，1 中对齐，2 右对齐
	VAlign         *int64                              `json:"vAlign,omitempty"`         // 垂直对齐, 0 上对齐，1 中对齐, 2 下对齐
	ForeColor      *string                             `json:"foreColor,omitempty"`      // 字体颜色
	BackColor      *string                             `json:"backColor,omitempty"`      // 背景颜色
	BorderType     *string                             `json:"borderType,omitempty"`     // 边框类型，可选 "FULL_BORDER"，"OUTER_BORDER"，"INNER_BORDER"，"NO_BORDER"，"LEFT_BORDER"，"RIGHT_BORDER"，"TOP_BORDER"，"BOTTOM_BORDER"
	BorderColor    *string                             `json:"borderColor,omitempty"`    // 边框颜色
	Clean          *bool                               `json:"clean,omitempty"`          // 是否清除所有格式,默认 false
}

type BatchSetSheetStyleReqDataStyleFont struct {
	Bold     *bool   `json:"bold,omitempty"`     // 是否加粗
	Italic   *bool   `json:"italic,omitempty"`   // 是否斜体
	FontSize *string `json:"fontSize,omitempty"` // 字体大小 字号大小为9~36 行距固定为1.5，如:10pt/1.5
}

type batchSetSheetStyleResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *BatchSetSheetStyleResp `json:"data,omitempty"`
}

type BatchSetSheetStyleResp struct {
	SpreadsheetToken    string                           `json:"spreadsheetToken,omitempty"`    // spreadsheet 的 token
	TotalUpdatedRows    int64                            `json:"totalUpdatedRows,omitempty"`    // 设置样式的总行数
	TotalUpdatedColumns int64                            `json:"totalUpdatedColumns,omitempty"` // 设置样式的总列数
	TotalUpdatedCells   int64                            `json:"totalUpdatedCells,omitempty"`   // 设置样式的单元格总数
	Revision            int64                            `json:"revision,omitempty"`            // sheet 的版本号
	Responses           *BatchSetSheetStyleRespResponses `json:"responses,omitempty"`           // 各个范围的设置单元格样式的范围、行列数等
}

type BatchSetSheetStyleRespResponses struct {
	SpreadsheetToken string `json:"spreadsheetToken,omitempty"` // spreadsheet 的 token
	UpdatedRange     string `json:"updatedRange,omitempty"`     // 设置样式的范围
	UpdatedRows      int64  `json:"updatedRows,omitempty"`      // 设置样式的行数
	UpdatedColumns   int64  `json:"updatedColumns,omitempty"`   // 设置样式的列数
	UpdatedCells     int64  `json:"updatedCells,omitempty"`     // 设置样式的单元格数
}
