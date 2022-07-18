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

// CreateAttendanceGroup 考勤组, 是对部门或者员工在某个特定场所及特定时间段内的出勤情况（包括上下班、迟到、早退、病假、婚假、丧假、公休、工作时间、加班情况等）的一种规则设定。
//
// 通过设置考勤组, 可以从部门、员工两个维度, 来设定考勤方式、考勤时间、考勤地点等考勤规则。
// 出于安全考虑, 目前通过该接口只允许修改自己创建的考勤组。
// ## 考勤组负责人
// 考勤组负责人可修改该考勤组的排班, 并查看该考勤组的考勤数据。
// 如果考勤组负责人同时被企业管理员赋予了考勤管理员的角色, 则该考勤组负责人还拥有考勤管理员的权限, 可以编辑及删除考勤规则。
// ## 考勤组人员
// 可按部门、员工两个维度, 设置需要参加考勤或无需参加考勤的人员。
// - 若是按部门维度添加的考勤人员, 当有新员工加入该部门时, 其会自动加入该考勤组。
// - 若是按员工维度添加的考勤人员, 当其上级部门被添加到其他考勤组时, 该员工不会更换考勤组。
// ## 考勤组类型
// 提供 3 种不同的考勤类型: 固定班制、排班制、自由班制。
// - 固定班制: 指考勤组内每位人员的上下班时间一致, 适用于上下班时间固定或无需安排多个班次的考勤组。
// - 排班制: 指考勤组人员的上下班时间不完全一致, 可自定义安排每位人员的上下班时间, 适用于存在多个班次如早晚班的考勤组。
// - 自由班制: 指没有具体的班次, 考勤组人员可以在打卡时段内自由打卡, 按照打卡时段统计上班时长。
// ## 考勤班次
// - 固定班制下, 需设置周一到周日每天安排哪个班次, 以及可针对特殊日期进行打卡设置。
// - 排班制下, 需对考勤组内每一位人员的每一天进行班次指定。
// - 自由班制下, 需设置一天中最早打卡时间和最晚打卡时间, 以及一周中哪几天需要打卡。
// ## 考勤方式
// 支持 3 种考勤方式: GPS 打卡、Wi-Fi 打卡、考勤机打卡。
// - GPS 打卡: 需设置经纬度信息及考勤地点名称。
// - Wi-Fi 打卡: 需设置 Wi-Fi 名称及 Wi-Fi 的 MAC 地址。
// - 考勤机打卡: 需设置考勤机名称及考勤机序号。
// ## 考勤其他设置
// - 规则设置: 支持设置是否允许外勤打卡, 是否允许补卡以及一个月补卡的次数, 是否允许 PC 端打卡。
// - 安全设置: 支持设置是否开启人脸识别打卡, 以及什么情况下开启人脸识别。
// - 统计设置: 支持设置考勤组人员是否可以查看到某些维度的统计数据。
// - 加班设置: 支持配置加班时间的计算规则。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create
func (r *AttendanceService) CreateAttendanceGroup(ctx context.Context, request *CreateAttendanceGroupReq, options ...MethodOptionFunc) (*CreateAttendanceGroupResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateAttendanceGroup != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateAttendanceGroup mock enable")
		return r.cli.mock.mockAttendanceCreateAttendanceGroup(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateAttendanceGroup",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/attendance/v1/groups",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createAttendanceGroupResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAttendanceCreateAttendanceGroup mock AttendanceCreateAttendanceGroup method
func (r *Mock) MockAttendanceCreateAttendanceGroup(f func(ctx context.Context, request *CreateAttendanceGroupReq, options ...MethodOptionFunc) (*CreateAttendanceGroupResp, *Response, error)) {
	r.mockAttendanceCreateAttendanceGroup = f
}

// UnMockAttendanceCreateAttendanceGroup un-mock AttendanceCreateAttendanceGroup method
func (r *Mock) UnMockAttendanceCreateAttendanceGroup() {
	r.mockAttendanceCreateAttendanceGroup = nil
}

// CreateAttendanceGroupReq ...
type CreateAttendanceGroupReq struct {
	EmployeeType EmployeeType                   `query:"employee_type" json:"-"` // 用户 ID 的类型, 示例值: "employee_id", 可选值有: employee_id: 员工 employee ID, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的用户 ID, employee_no: 员工工号, 即飞书管理后台 > 组织架构 > 成员与部门 > 成员详情中的工号
	DeptType     string                         `query:"dept_type" json:"-"`     // 部门 ID 的类型, 示例值: "od-fcb45c28a45311afd441b8869541ece8", 可选值有: open_id: 暂时只支持部门的 openid
	Group        *CreateAttendanceGroupReqGroup `json:"group,omitempty"`         // 6921319402260496386
	OperatorID   *string                        `json:"operator_id,omitempty"`   // 操作人uid, 如果您未操作[考勤管理后台“API 接入”流程](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/attendance-development-guidelines), 则此字段为必填字段, 示例值: "dd31248a"
}

// CreateAttendanceGroupReqGroup ...
type CreateAttendanceGroupReqGroup struct {
	GroupID                 *string                                               `json:"group_id,omitempty"`                    // 考勤组 ID（仅修改时提供）, 需要从“获取打卡结果”的接口中获取 groupId, 示例值: "6919358128597097404"
	GroupName               string                                                `json:"group_name,omitempty"`                  // 考勤组名称, 示例值: "开心考勤"
	TimeZone                string                                                `json:"time_zone,omitempty"`                   // 时区, 示例值: "Asia/Shanghai"
	BindDeptIDs             []string                                              `json:"bind_dept_ids,omitempty"`               // 绑定的部门 ID, 示例值: od-fcb45c28a45311afd440b7869541fce8
	ExceptDeptIDs           []string                                              `json:"except_dept_ids,omitempty"`             // 排除的部门 ID, 示例值: od-fcb45c28a45311afd440b7869541fce8
	BindUserIDs             []string                                              `json:"bind_user_ids,omitempty"`               // 绑定的用户 ID, 示例值: 52aa1fa1
	ExceptUserIDs           []string                                              `json:"except_user_ids,omitempty"`             // 排除的用户 ID, 示例值: 52aa1fa1
	GroupLeaderIDs          []string                                              `json:"group_leader_ids,omitempty"`            // 考勤主负责人 ID 列表, 必选字段, 示例值: 2bg4a9be
	SubGroupLeaderIDs       []string                                              `json:"sub_group_leader_ids,omitempty"`        // 考勤子负责人 ID 列表, 示例值: 52aa1fa1
	AllowOutPunch           *bool                                                 `json:"allow_out_punch,omitempty"`             // 是否允许外勤打卡, 示例值: true
	OutPunchNeedApproval    *bool                                                 `json:"out_punch_need_approval,omitempty"`     // 外勤打卡需审批（需要允许外勤打卡才能设置生效）, 示例值: true
	OutPunchNeedRemark      *bool                                                 `json:"out_punch_need_remark,omitempty"`       // 外勤打卡需填写备注（需要允许外勤打卡才能设置生效）, 示例值: true
	OutPunchNeedPhoto       *bool                                                 `json:"out_punch_need_photo,omitempty"`        // 外勤打卡需拍照（需要允许外勤打卡才能设置生效）, 示例值: true
	OutPunchAllowedHideAddr *bool                                                 `json:"out_punch_allowed_hide_addr,omitempty"` // 外勤打卡允许员工隐藏详细地址（需要允许外勤打卡才能设置生效）, 示例值: true
	AllowPcPunch            *bool                                                 `json:"allow_pc_punch,omitempty"`              // 是否允许 PC 端打卡, 示例值: true
	AllowRemedy             *bool                                                 `json:"allow_remedy,omitempty"`                // 是否限制补卡, 示例值: true
	RemedyLimit             *bool                                                 `json:"remedy_limit,omitempty"`                // 是否限制补卡次数, 示例值: true
	RemedyLimitCount        *int64                                                `json:"remedy_limit_count,omitempty"`          // 补卡次数, 示例值: 3
	RemedyDateLimit         *bool                                                 `json:"remedy_date_limit,omitempty"`           // 是否限制补卡时间, 示例值: true
	RemedyDateNum           *int64                                                `json:"remedy_date_num,omitempty"`             // 补卡时间, 几天内补卡, 示例值: 3
	AllowRemedyTypeLack     *bool                                                 `json:"allow_remedy_type_lack,omitempty"`      // 允许缺卡补卡（需要允许补卡才能设置生效）, 示例值: true
	AllowRemedyTypeLate     *bool                                                 `json:"allow_remedy_type_late,omitempty"`      // 允许迟到补卡（需要允许补卡才能设置生效）, 示例值: true
	AllowRemedyTypeEarly    *bool                                                 `json:"allow_remedy_type_early,omitempty"`     // 允许早退补卡（需要允许补卡才能设置生效）, 示例值: true
	AllowRemedyTypeNormal   *bool                                                 `json:"allow_remedy_type_normal,omitempty"`    // 允许正常补卡（需要允许补卡才能设置生效）, 示例值: true
	ShowCumulativeTime      *bool                                                 `json:"show_cumulative_time,omitempty"`        // 是否展示累计时长, 示例值: true
	ShowOverTime            *bool                                                 `json:"show_over_time,omitempty"`              // 是否展示加班时长, 示例值: true
	HideStaffPunchTime      *bool                                                 `json:"hide_staff_punch_time,omitempty"`       // 是否隐藏员工打卡详情, 示例值: true
	FacePunch               *bool                                                 `json:"face_punch,omitempty"`                  // 是否开启人脸识别打卡, 示例值: true
	FacePunchCfg            *int64                                                `json:"face_punch_cfg,omitempty"`              // 人脸识别打卡规则, 1: 每次打卡均需人脸识别, 2: 疑似作弊打卡时需要人脸识别, 示例值: 1
	FaceDowngrade           *bool                                                 `json:"face_downgrade,omitempty"`              // 人脸识别失败时是否允许普通拍照打卡, 示例值: true
	ReplaceBasicPic         *bool                                                 `json:"replace_basic_pic,omitempty"`           // 人脸识别失败时是否允许替换基准图片, 示例值: true
	Machines                []*CreateAttendanceGroupReqGroupMachine               `json:"machines,omitempty"`                    // 考勤机列表
	GpsRange                *int64                                                `json:"gps_range,omitempty"`                   // GPS 打卡的有效范围（不建议使用）, 示例值: 300
	Locations               []*CreateAttendanceGroupReqGroupLocation              `json:"locations,omitempty"`                   // 地址列表
	GroupType               int64                                                 `json:"group_type,omitempty"`                  // 考勤类型, 0: 固定班制, 2: 排班制, 3: 自由班制, 示例值: 0
	PunchDayShiftIDs        []string                                              `json:"punch_day_shift_ids,omitempty"`         // 固定班制必须填, 示例值: 6921319402260496386
	FreePunchCfg            *CreateAttendanceGroupReqGroupFreePunchCfg            `json:"free_punch_cfg,omitempty"`              // 配置自由班制
	CalendarID              int64                                                 `json:"calendar_id,omitempty"`                 // 国家日历  ID, 0: 不根据国家日历排休, 1: 中国大陆, 2: 美国, 3: 日本, 4: 印度, 5: 新加坡, 默认 1, 示例值: 1
	NeedPunchSpecialDays    []*CreateAttendanceGroupReqGroupNeedPunchSpecialDay   `json:"need_punch_special_days,omitempty"`     // 必须打卡的特殊日期
	NoNeedPunchSpecialDays  []*CreateAttendanceGroupReqGroupNoNeedPunchSpecialDay `json:"no_need_punch_special_days,omitempty"`  // 无需打卡的特殊日期
	WorkDayNoPunchAsLack    *bool                                                 `json:"work_day_no_punch_as_lack,omitempty"`   // 自由班制下工作日不打卡是否记为缺卡, 示例值: true
	EffectNow               *bool                                                 `json:"effect_now,omitempty"`                  // 是否立即生效, 默认 false, 示例值: true
	RemedyPeriodType        *int64                                                `json:"remedy_period_type,omitempty"`          // 补卡周期类型, 示例值: 0
	RemedyPeriodCustomDate  *int64                                                `json:"remedy_period_custom_date,omitempty"`   // 补卡自定义周期起始日期, 示例值: 1
	PunchType               *int64                                                `json:"punch_type,omitempty"`                  // 打卡类型, 位运算。1: GPS 打卡, 2: Wi-Fi 打卡, 4: 考勤机打卡, 8: IP 打卡, 示例值: 1
	RestClockInNeedApproval *bool                                                 `json:"rest_clockIn_need_approval,omitempty"`  // 休息日打卡需审批, 示例值: true
	ClockInNeedPhoto        *bool                                                 `json:"clockIn_need_photo,omitempty"`          // 每次打卡均需拍照, 示例值: true
}

// CreateAttendanceGroupReqGroupFreePunchCfg ...
type CreateAttendanceGroupReqGroupFreePunchCfg struct {
	FreeStartTime        string `json:"free_start_time,omitempty"`           // 自由班制打卡开始时间, 示例值: "7:00"
	FreeEndTime          string `json:"free_end_time,omitempty"`             // 自由班制打卡结束时间, 示例值: "18:00"
	PunchDay             int64  `json:"punch_day,omitempty"`                 // 打卡的时间, 为 7 位数字, 每一位依次代表周一到周日, 0 为不上班, 1 为上班, 示例值: 1111100
	WorkDayNoPunchAsLack *bool  `json:"work_day_no_punch_as_lack,omitempty"` // 工作日不打卡是否记为缺卡, 示例值: true
}

// CreateAttendanceGroupReqGroupLocation ...
type CreateAttendanceGroupReqGroupLocation struct {
	LocationName string   `json:"location_name,omitempty"` // 地址名称, 示例值: "浙江省杭州市余杭区五常街道木桥头西溪八方城"
	LocationType int64    `json:"location_type,omitempty"` // 地址类型, 1: GPS, 2: Wi-Fi, 8: IP, 示例值: 1
	Latitude     *float64 `json:"latitude,omitempty"`      // 地址纬度, 示例值: 30.28994
	Longitude    *float64 `json:"longitude,omitempty"`     // 地址经度, 示例值: 120.04509
	Ssid         *string  `json:"ssid,omitempty"`          // Wi-Fi 名称, 示例值: "TP-Link-af12ca"
	Bssid        *string  `json:"bssid,omitempty"`         // Wi-Fi 的 MAC 地址, 示例值: "08:00:20:0A:8C:6D"
	MapType      *int64   `json:"map_type,omitempty"`      // 地图类型, 1: 高德, 2: 谷歌, 示例值: 1
	Address      *string  `json:"address,omitempty"`       // 地址名称, 示例值: "北京市海淀区中航广场"
	Ip           *string  `json:"ip,omitempty"`            // IP 地址, 示例值: "122.224.123.146"
	Feature      *string  `json:"feature,omitempty"`       // 额外信息, 例如: 运营商信息, 示例值: "中国电信"
	GpsRange     *int64   `json:"gps_range,omitempty"`     // GPS 打卡的有效范围, 示例值: 300
}

// CreateAttendanceGroupReqGroupMachine ...
type CreateAttendanceGroupReqGroupMachine struct {
	MachineSn   string `json:"machine_sn,omitempty"`   // 考勤机序列号, 示例值: "FS0701"
	MachineName string `json:"machine_name,omitempty"` // 考勤机名称, 示例值: "创实 9 楼"
}

// CreateAttendanceGroupReqGroupNeedPunchSpecialDay ...
type CreateAttendanceGroupReqGroupNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期, 示例值: 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID, 示例值: "6919668827865513935"
}

// CreateAttendanceGroupReqGroupNoNeedPunchSpecialDay ...
type CreateAttendanceGroupReqGroupNoNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期, 示例值: 20190101
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID, 示例值: "6919668827865513935"
}

// CreateAttendanceGroupResp ...
type CreateAttendanceGroupResp struct {
	Group *CreateAttendanceGroupRespGroup `json:"group,omitempty"` // 6921319402260496386
}

// CreateAttendanceGroupRespGroup ...
type CreateAttendanceGroupRespGroup struct {
	GroupID                 string                                                 `json:"group_id,omitempty"`                    // 考勤组 ID（仅修改时提供）, 需要从“获取打卡结果”的接口中获取 groupId
	GroupName               string                                                 `json:"group_name,omitempty"`                  // 考勤组名称
	TimeZone                string                                                 `json:"time_zone,omitempty"`                   // 时区
	BindDeptIDs             []string                                               `json:"bind_dept_ids,omitempty"`               // 绑定的部门 ID
	ExceptDeptIDs           []string                                               `json:"except_dept_ids,omitempty"`             // 排除的部门 ID
	BindUserIDs             []string                                               `json:"bind_user_ids,omitempty"`               // 绑定的用户 ID
	ExceptUserIDs           []string                                               `json:"except_user_ids,omitempty"`             // 排除的用户 ID
	GroupLeaderIDs          []string                                               `json:"group_leader_ids,omitempty"`            // 考勤主负责人 ID 列表, 必选字段
	SubGroupLeaderIDs       []string                                               `json:"sub_group_leader_ids,omitempty"`        // 考勤子负责人 ID 列表
	AllowOutPunch           bool                                                   `json:"allow_out_punch,omitempty"`             // 是否允许外勤打卡
	OutPunchNeedApproval    bool                                                   `json:"out_punch_need_approval,omitempty"`     // 外勤打卡需审批（需要允许外勤打卡才能设置生效）
	OutPunchNeedRemark      bool                                                   `json:"out_punch_need_remark,omitempty"`       // 外勤打卡需填写备注（需要允许外勤打卡才能设置生效）
	OutPunchNeedPhoto       bool                                                   `json:"out_punch_need_photo,omitempty"`        // 外勤打卡需拍照（需要允许外勤打卡才能设置生效）
	OutPunchAllowedHideAddr bool                                                   `json:"out_punch_allowed_hide_addr,omitempty"` // 外勤打卡允许员工隐藏详细地址（需要允许外勤打卡才能设置生效）
	AllowPcPunch            bool                                                   `json:"allow_pc_punch,omitempty"`              // 是否允许 PC 端打卡
	AllowRemedy             bool                                                   `json:"allow_remedy,omitempty"`                // 是否限制补卡
	RemedyLimit             bool                                                   `json:"remedy_limit,omitempty"`                // 是否限制补卡次数
	RemedyLimitCount        int64                                                  `json:"remedy_limit_count,omitempty"`          // 补卡次数
	RemedyDateLimit         bool                                                   `json:"remedy_date_limit,omitempty"`           // 是否限制补卡时间
	RemedyDateNum           int64                                                  `json:"remedy_date_num,omitempty"`             // 补卡时间, 几天内补卡
	AllowRemedyTypeLack     bool                                                   `json:"allow_remedy_type_lack,omitempty"`      // 允许缺卡补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeLate     bool                                                   `json:"allow_remedy_type_late,omitempty"`      // 允许迟到补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeEarly    bool                                                   `json:"allow_remedy_type_early,omitempty"`     // 允许早退补卡（需要允许补卡才能设置生效）
	AllowRemedyTypeNormal   bool                                                   `json:"allow_remedy_type_normal,omitempty"`    // 允许正常补卡（需要允许补卡才能设置生效）
	ShowCumulativeTime      bool                                                   `json:"show_cumulative_time,omitempty"`        // 是否展示累计时长
	ShowOverTime            bool                                                   `json:"show_over_time,omitempty"`              // 是否展示加班时长
	HideStaffPunchTime      bool                                                   `json:"hide_staff_punch_time,omitempty"`       // 是否隐藏员工打卡详情
	FacePunch               bool                                                   `json:"face_punch,omitempty"`                  // 是否开启人脸识别打卡
	FacePunchCfg            int64                                                  `json:"face_punch_cfg,omitempty"`              // 人脸识别打卡规则, 1: 每次打卡均需人脸识别, 2: 疑似作弊打卡时需要人脸识别
	FaceDowngrade           bool                                                   `json:"face_downgrade,omitempty"`              // 人脸识别失败时是否允许普通拍照打卡
	ReplaceBasicPic         bool                                                   `json:"replace_basic_pic,omitempty"`           // 人脸识别失败时是否允许替换基准图片
	Machines                []*CreateAttendanceGroupRespGroupMachine               `json:"machines,omitempty"`                    // 考勤机列表
	GpsRange                int64                                                  `json:"gps_range,omitempty"`                   // GPS 打卡的有效范围（不建议使用）
	Locations               []*CreateAttendanceGroupRespGroupLocation              `json:"locations,omitempty"`                   // 地址列表
	GroupType               int64                                                  `json:"group_type,omitempty"`                  // 考勤类型, 0: 固定班制, 2: 排班制, 3: 自由班制
	PunchDayShiftIDs        []string                                               `json:"punch_day_shift_ids,omitempty"`         // 固定班制必须填
	FreePunchCfg            *CreateAttendanceGroupRespGroupFreePunchCfg            `json:"free_punch_cfg,omitempty"`              // 配置自由班制
	CalendarID              int64                                                  `json:"calendar_id,omitempty"`                 // 国家日历  ID, 0: 不根据国家日历排休, 1: 中国大陆, 2: 美国, 3: 日本, 4: 印度, 5: 新加坡, 默认 1
	NeedPunchSpecialDays    []*CreateAttendanceGroupRespGroupNeedPunchSpecialDay   `json:"need_punch_special_days,omitempty"`     // 必须打卡的特殊日期
	NoNeedPunchSpecialDays  []*CreateAttendanceGroupRespGroupNoNeedPunchSpecialDay `json:"no_need_punch_special_days,omitempty"`  // 无需打卡的特殊日期
	WorkDayNoPunchAsLack    bool                                                   `json:"work_day_no_punch_as_lack,omitempty"`   // 自由班制下工作日不打卡是否记为缺卡
	EffectNow               bool                                                   `json:"effect_now,omitempty"`                  // 是否立即生效, 默认 false
	RemedyPeriodType        int64                                                  `json:"remedy_period_type,omitempty"`          // 补卡周期类型
	RemedyPeriodCustomDate  int64                                                  `json:"remedy_period_custom_date,omitempty"`   // 补卡自定义周期起始日期
	PunchType               int64                                                  `json:"punch_type,omitempty"`                  // 打卡类型, 位运算。1: GPS 打卡, 2: Wi-Fi 打卡, 4: 考勤机打卡, 8: IP 打卡
	RestClockInNeedApproval bool                                                   `json:"rest_clockIn_need_approval,omitempty"`  // 休息日打卡需审批
	ClockInNeedPhoto        bool                                                   `json:"clockIn_need_photo,omitempty"`          // 每次打卡均需拍照
}

// CreateAttendanceGroupRespGroupFreePunchCfg ...
type CreateAttendanceGroupRespGroupFreePunchCfg struct {
	FreeStartTime        string `json:"free_start_time,omitempty"`           // 自由班制打卡开始时间
	FreeEndTime          string `json:"free_end_time,omitempty"`             // 自由班制打卡结束时间
	PunchDay             int64  `json:"punch_day,omitempty"`                 // 打卡的时间, 为 7 位数字, 每一位依次代表周一到周日, 0 为不上班, 1 为上班
	WorkDayNoPunchAsLack bool   `json:"work_day_no_punch_as_lack,omitempty"` // 工作日不打卡是否记为缺卡
}

// CreateAttendanceGroupRespGroupLocation ...
type CreateAttendanceGroupRespGroupLocation struct {
	LocationID   string  `json:"location_id,omitempty"`   // 地址 ID
	LocationName string  `json:"location_name,omitempty"` // 地址名称
	LocationType int64   `json:"location_type,omitempty"` // 地址类型, 1: GPS, 2: Wi-Fi, 8: IP
	Latitude     float64 `json:"latitude,omitempty"`      // 地址纬度
	Longitude    float64 `json:"longitude,omitempty"`     // 地址经度
	Ssid         string  `json:"ssid,omitempty"`          // Wi-Fi 名称
	Bssid        string  `json:"bssid,omitempty"`         // Wi-Fi 的 MAC 地址
	MapType      int64   `json:"map_type,omitempty"`      // 地图类型, 1: 高德, 2: 谷歌
	Address      string  `json:"address,omitempty"`       // 地址名称
	Ip           string  `json:"ip,omitempty"`            // IP 地址
	Feature      string  `json:"feature,omitempty"`       // 额外信息, 例如: 运营商信息
	GpsRange     int64   `json:"gps_range,omitempty"`     // GPS 打卡的有效范围
}

// CreateAttendanceGroupRespGroupMachine ...
type CreateAttendanceGroupRespGroupMachine struct {
	MachineSn   string `json:"machine_sn,omitempty"`   // 考勤机序列号
	MachineName string `json:"machine_name,omitempty"` // 考勤机名称
}

// CreateAttendanceGroupRespGroupNeedPunchSpecialDay ...
type CreateAttendanceGroupRespGroupNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

// CreateAttendanceGroupRespGroupNoNeedPunchSpecialDay ...
type CreateAttendanceGroupRespGroupNoNeedPunchSpecialDay struct {
	PunchDay int64  `json:"punch_day,omitempty"` // 打卡日期
	ShiftID  string `json:"shift_id,omitempty"`  // 班次 ID
}

// createAttendanceGroupResp ...
type createAttendanceGroupResp struct {
	Code int64                      `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 错误描述
	Data *CreateAttendanceGroupResp `json:"data,omitempty"`
}
