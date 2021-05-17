// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetCountryList 新建建筑时需要标明所处国家/地区，该接口用于获得系统预先提供的可供选择的国家 /地区列表。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQTNwYjL0UDM24CN1AjN
func (r *MeetingRoomService) GetCountryList(ctx context.Context, request *GetCountryListReq, options ...MethodOptionFunc) (*GetCountryListResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomGetCountryList != nil {
		r.cli.logDebug(ctx, "[lark] MeetingRoom#GetCountryList mock enable")
		return r.cli.mock.mockMeetingRoomGetCountryList(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] MeetingRoom#GetCountryList call api")
	r.cli.logDebug(ctx, "[lark] MeetingRoom#GetCountryList request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/country/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCountryListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] MeetingRoom#GetCountryList GET https://open.feishu.cn/open-apis/meeting_room/country/list failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] MeetingRoom#GetCountryList GET https://open.feishu.cn/open-apis/meeting_room/country/list failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("MeetingRoom", "GetCountryList", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] MeetingRoom#GetCountryList request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockMeetingRoomGetCountryList(f func(ctx context.Context, request *GetCountryListReq, options ...MethodOptionFunc) (*GetCountryListResp, *Response, error)) {
	r.mockMeetingRoomGetCountryList = f
}

func (r *Mock) UnMockMeetingRoomGetCountryList() {
	r.mockMeetingRoomGetCountryList = nil
}

type GetCountryListReq struct{}

type getCountryListResp struct {
	Code int                 `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string              `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *GetCountryListResp `json:"data,omitempty"` // 返回业务信息
}

type GetCountryListResp struct {
	Countries *GetCountryListRespCountries `json:"countries,omitempty"` // 国家地区列表
}

type GetCountryListRespCountries struct {
	CountryID string `json:"country_id,omitempty"` // 国家地区ID
	Name      string `json:"name,omitempty"`       // 国家地区名称
}