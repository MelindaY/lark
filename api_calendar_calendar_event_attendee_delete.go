// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteCalendarEventAttendee 批量删除日程的参与人。
//
// - 当前身份需要有日历的 writer 或 owner 权限，并且日历的类型只能为 primary 或 shared。
// - 当前身份需要是日程的组织者。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/calendar-v4/calendar-event-attendee/batch_delete
func (r *CalendarService) DeleteCalendarEventAttendee(ctx context.Context, request *DeleteCalendarEventAttendeeReq, options ...MethodOptionFunc) (*DeleteCalendarEventAttendeeResp, *Response, error) {
	if r.cli.mock.mockCalendarDeleteCalendarEventAttendee != nil {
		r.cli.logDebug(ctx, "[lark] Calendar#DeleteCalendarEventAttendee mock enable")
		return r.cli.mock.mockCalendarDeleteCalendarEventAttendee(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Calendar#DeleteCalendarEventAttendee call api")
	r.cli.logDebug(ctx, "[lark] Calendar#DeleteCalendarEventAttendee request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(deleteCalendarEventAttendeeResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Calendar#DeleteCalendarEventAttendee POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Calendar#DeleteCalendarEventAttendee POST https://open.feishu.cn/open-apis/calendar/v4/calendars/:calendar_id/events/:event_id/attendees/batch_delete failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Calendar", "DeleteCalendarEventAttendee", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Calendar#DeleteCalendarEventAttendee request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockCalendarDeleteCalendarEventAttendee(f func(ctx context.Context, request *DeleteCalendarEventAttendeeReq, options ...MethodOptionFunc) (*DeleteCalendarEventAttendeeResp, *Response, error)) {
	r.mockCalendarDeleteCalendarEventAttendee = f
}

func (r *Mock) UnMockCalendarDeleteCalendarEventAttendee() {
	r.mockCalendarDeleteCalendarEventAttendee = nil
}

type DeleteCalendarEventAttendeeReq struct {
	CalendarID       string   `path:"calendar_id" json:"-"`        // 日历 ID, 示例值："feishu.cn_xxxxxxxxxx@group.calendar.feishu.cn"
	EventID          string   `path:"event_id" json:"-"`           // 日程 ID, 示例值："xxxxxxxxx_0"
	AttendeeIDs      []string `json:"attendee_ids,omitempty"`      // 要移除的参与人 ID 列表
	NeedNotification *bool    `json:"need_notification,omitempty"` // 删除日程参与人时是否要给参与人发送bot通知，默认为true, 示例值：false
}

type deleteCalendarEventAttendeeResp struct {
	Code int                              `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *DeleteCalendarEventAttendeeResp `json:"data,omitempty"`
}

type DeleteCalendarEventAttendeeResp struct{}
