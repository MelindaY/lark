// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// GetAttendanceUserTaskRemedy 获取授权内员工的补卡记录。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query
func (r *AttendanceService) GetAttendanceUserTaskRemedy(ctx context.Context, request *GetAttendanceUserTaskRemedyReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskRemedyResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceUserTaskRemedy != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceUserTaskRemedy mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceUserTaskRemedy(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceUserTaskRemedy",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/user_task_remedys/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceUserTaskRemedyResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceGetAttendanceUserTaskRemedy mock AttendanceGetAttendanceUserTaskRemedy method
func (r *Mock) MockAttendanceGetAttendanceUserTaskRemedy(f func(ctx context.Context, request *GetAttendanceUserTaskRemedyReq, options ...MethodOptionFunc) (*GetAttendanceUserTaskRemedyResp, *Response, error)) {
	r.mockAttendanceGetAttendanceUserTaskRemedy = f
}

// UnMockAttendanceGetAttendanceUserTaskRemedy un-mock AttendanceGetAttendanceUserTaskRemedy method
func (r *Mock) UnMockAttendanceGetAttendanceUserTaskRemedy() {
	r.mockAttendanceGetAttendanceUserTaskRemedy = nil
}

// GetAttendanceUserTaskRemedyReq ...
type GetAttendanceUserTaskRemedyReq struct {
	EmployeeType  EmployeeType `query:"employee_type" json:"-"`   // 请求体中的 user_ids 和响应体中的 user_id 的员工工号类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	UserIDs       []string     `json:"user_ids,omitempty"`        // employee_no 或 employee_id 列表, 示例值: ["abd754f7"]
	CheckTimeFrom string       `json:"check_time_from,omitempty"` // 查询的起始时间, 精确到秒的时间戳, 示例值: "1566641088"
	CheckTimeTo   string       `json:"check_time_to,omitempty"`   // 查询的结束时间, 精确到秒的时间戳, 示例值: "1592561088"
}

// GetAttendanceUserTaskRemedyResp ...
type GetAttendanceUserTaskRemedyResp struct {
	UserRemedys []*GetAttendanceUserTaskRemedyRespUserRemedy `json:"user_remedys,omitempty"` // 补卡记录列表
}

// GetAttendanceUserTaskRemedyRespUserRemedy ...
type GetAttendanceUserTaskRemedyRespUserRemedy struct {
	UserID     string `json:"user_id,omitempty"`     // 用户 ID
	RemedyDate int64  `json:"remedy_date,omitempty"` // 补卡日期
	PunchNo    int64  `json:"punch_no,omitempty"`    // 第几次上下班, 0: 第 1 次上下班, 1: 第 2 次上下班, 2: 第 3 次上下班, 自由班制填 0
	WorkType   int64  `json:"work_type,omitempty"`   // 上班 / 下班, 1: 上班, 2: 下班, 自由班制填 0
	ApprovalID string `json:"approval_id,omitempty"` // 审批 ID
	RemedyTime string `json:"remedy_time,omitempty"` // 补卡时间, 时间格式为 yyyy-MM-dd HH:mm
	Status     int64  `json:"status,omitempty"`      // 补卡状态（默认为审批中）, 可选值有: 0: 审批中, 2: 已通过, 3: 已取消, 4: 通过后撤回
	Reason     string `json:"reason,omitempty"`      // 补卡原因
	Time       string `json:"time,omitempty"`        // 补卡时间, 精确到秒的时间戳
	TimeZone   string `json:"time_zone,omitempty"`   // 补卡时考勤组时区
	CreateTime string `json:"create_time,omitempty"` // 补卡发起时间, 精确到秒的时间戳
	UpdateTime string `json:"update_time,omitempty"` // 补卡状态更新时间, 精确到秒的时间戳
}

// getAttendanceUserTaskRemedyResp ...
type getAttendanceUserTaskRemedyResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceUserTaskRemedyResp `json:"data,omitempty"`
}
