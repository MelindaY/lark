// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetUserTaskRemedy
//
// 获取授权内员工的补卡记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//GetUsersRemedyRecords
func (r *AttendanceService) GetUserTaskRemedy(ctx context.Context, request *GetUserTaskRemedyReq, options ...MethodOptionFunc) (*GetUserTaskRemedyResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetUserTaskRemedy != nil {
		r.cli.logDebug(ctx, "[lark] Attendance#GetUserTaskRemedy mock enable")
		return r.cli.mock.mockAttendanceGetUserTaskRemedy(ctx, request, options...)
	}

	r.cli.logInfo(ctx, "[lark] Attendance#GetUserTaskRemedy call api")
	r.cli.logDebug(ctx, "[lark] Attendance#GetUserTaskRemedy request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getUserTaskRemedyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		r.cli.logError(ctx, "[lark] Attendance#GetUserTaskRemedy POST https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query failed: %s", err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.logError(ctx, "[lark] Attendance#GetUserTaskRemedy POST https://open.feishu.cn/open-apis/attendance/v1/user_task_remedys/query failed, code: %d, msg: %s", resp.Code, resp.Msg)
		return nil, response, NewError("Attendance", "GetUserTaskRemedy", resp.Code, resp.Msg)
	}

	r.cli.logDebug(ctx, "[lark] Attendance#GetUserTaskRemedy request_id: %s, response: %s", response.RequestID, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAttendanceGetUserTaskRemedy(f func(ctx context.Context, request *GetUserTaskRemedyReq, options ...MethodOptionFunc) (*GetUserTaskRemedyResp, *Response, error)) {
	r.mockAttendanceGetUserTaskRemedy = f
}

func (r *Mock) UnMockAttendanceGetUserTaskRemedy() {
	r.mockAttendanceGetUserTaskRemedy = nil
}

type GetUserTaskRemedyReq struct {
	EmployeeType  EmployeeType `query:"employee_type" json:"-"`   // 请求体中的 user_id 的员工工号类型，可用值：【employee_id（员工的 employeeId），employee_no（员工工号）】，示例值：“employee_id”
	UserIDs       []string     `json:"user_ids,omitempty"`        // employee_no 或 employee_id 列表
	CheckTimeFrom string       `json:"check_time_from,omitempty"` // 查询的起始时间，精确到秒的时间戳
	CheckTimeTo   string       `json:"check_time_to,omitempty"`   // 查询的结束时间，精确到秒的时间戳
}

type getUserTaskRemedyResp struct {
	Code int                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                 `json:"msg,omitempty"`  // 错误描述
	Data *GetUserTaskRemedyResp `json:"data,omitempty"` // -
}

type GetUserTaskRemedyResp struct {
	UserRemedys []*GetUserTaskRemedyRespUserRemedy `json:"user_remedys,omitempty"` // 补卡记录列表
}

type GetUserTaskRemedyRespUserRemedy struct {
	UserID     string `json:"user_id,omitempty"`     // 员工工号
	Status     int    `json:"status,omitempty"`      // 补卡状态，可用值：【0（pending），2（已通过），3（已取消），4（通过后撤销）】
	Reason     string `json:"reason,omitempty"`      // 补卡原因
	Time       string `json:"time,omitempty"`        // 补卡时间，精确到秒的时间戳
	TimeZone   string `json:"time_zone,omitempty"`   // 补卡时的考勤组时区
	CreateTime string `json:"create_time,omitempty"` // 补卡发起时间，精确到秒的时间戳
	UpdateTime string `json:"update_time,omitempty"` // 补卡状态更新时间，精确到秒的时间戳
}
