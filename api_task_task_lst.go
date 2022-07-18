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

// GetTaskList 以分页的方式获取任务列表。当使用user_access_token时, 获取与该用户身份相关的所有任务。当使用tenant_access_token时, 获取以该应用身份通过“创建任务“接口创建的所有任务（并非获取该应用所在租户下所有用户创建的任务）。
//
// 本接口支持通过任务创建时间以及任务的完成状态对任务进行过滤。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/task-v1/task/list
func (r *TaskService) GetTaskList(ctx context.Context, request *GetTaskListReq, options ...MethodOptionFunc) (*GetTaskListResp, *Response, error) {
	if r.cli.mock.mockTaskGetTaskList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Task#GetTaskList mock enable")
		return r.cli.mock.mockTaskGetTaskList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "GetTaskList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v1/tasks",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getTaskListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskGetTaskList mock TaskGetTaskList method
func (r *Mock) MockTaskGetTaskList(f func(ctx context.Context, request *GetTaskListReq, options ...MethodOptionFunc) (*GetTaskListResp, *Response, error)) {
	r.mockTaskGetTaskList = f
}

// UnMockTaskGetTaskList un-mock TaskGetTaskList method
func (r *Mock) UnMockTaskGetTaskList() {
	r.mockTaskGetTaskList = nil
}

// GetTaskListReq ...
type GetTaskListReq struct {
	PageSize        *int64  `query:"page_size" json:"-"`         // 分页大小, 示例值: 10, 最大值: `100`
	PageToken       *string `query:"page_token" json:"-"`        // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "MTYzMTg3ODUxNQ["
	StartCreateTime *string `query:"start_create_time" json:"-"` // 范围查询任务时, 查询的起始时间。不填时默认起始时间为第一个任务的创建时间, 示例值: "1652323331"
	EndCreateTime   *string `query:"end_create_time" json:"-"`   // 范围查询任务时, 查询的结束时间。不填时默认结束时间为最后一个任务的创建时间, 示例值: "1652323335"
	TaskCompleted   *bool   `query:"task_completed" json:"-"`    // 可用于查询时过滤任务完成状态。true表示只返回已完成的任务, false表示只返回未完成的任务。不填时表示同时返回两种完成状态的任务, 示例值: false
	UserIDType      *IDType `query:"user_id_type" json:"-"`      // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
}

// GetTaskListResp ...
type GetTaskListResp struct {
	Items     []*GetTaskListRespItem `json:"items,omitempty"`      // 任务列表
	PageToken string                 `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                   `json:"has_more,omitempty"`   // 是否还有更多项
}

// GetTaskListRespItem ...
type GetTaskListRespItem struct {
	ID              string                             `json:"id,omitempty"`               // 任务 ID, 由飞书任务服务器发号
	Summary         string                             `json:"summary,omitempty"`          // 任务标题。创建任务时, 如果没有标题填充, 飞书服务器会将其视为无主题的任务
	Description     string                             `json:"description,omitempty"`      // 任务备注
	CompleteTime    string                             `json:"complete_time,omitempty"`    // 任务的完成时间戳（单位为秒）, 如果完成时间为 0, 则表示任务尚未完成
	CreatorID       string                             `json:"creator_id,omitempty"`       // 任务的创建者 ID。在创建任务时无需填充该字段
	Extra           string                             `json:"extra,omitempty"`            // 接入方可以自定义的附属信息二进制格式, 采用 base64 编码, 解析方式由接入方自己决定
	CreateTime      string                             `json:"create_time,omitempty"`      // 任务的创建时间戳（单位为秒）
	UpdateTime      string                             `json:"update_time,omitempty"`      // 任务的更新时间戳（单位为秒）
	Due             *GetTaskListRespItemDue            `json:"due,omitempty"`              // 任务的截止时间设置
	Origin          *GetTaskListRespItemOrigin         `json:"origin,omitempty"`           // 任务关联的第三方平台来源信息
	CanEdit         bool                               `json:"can_edit,omitempty"`         // 此字段用于控制该任务在飞书任务中心是否可编辑, 默认为false, 若为true则第三方需考虑是否需要接入事件来接收任务在任务中心的变更信息, （即将废弃）
	Custom          string                             `json:"custom,omitempty"`           // 此字段用于存储第三方需透传到端上的自定义数据, Json格式。取值举例中custom_complete字段存储「完成」按钮的跳转链接（href）或提示信息（tip）, pc、ios、android三端均可自定义, 其中tip字段的key为语言类型, value为提示信息, 可自行增加或减少语言类型, 支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn。href的优先级高于tip, href和tip同时不为空时只跳转不提示。链接和提示信息可自定义, 其余的key需按举例中的结构传递
	Source          int64                              `json:"source,omitempty"`           // 任务创建的来源, 可选值有: 0: 未知类型, 1: 来源任务中心创建, 2: 来源消息转任务, 3: 来源云文档, 4: 来源文档单品, 5: 来源PANO, 6: 来源tenant_access_token创建的任务, 7: 来源user_access_token创建的任务, 8: 来源新版云文档
	Followers       []*GetTaskListRespItemFollower     `json:"followers,omitempty"`        // 任务的关注者
	Collaborators   []*GetTaskListRespItemCollaborator `json:"collaborators,omitempty"`    // 任务的执行者
	CollaboratorIDs []string                           `json:"collaborator_ids,omitempty"` // 创建任务时添加的执行者用户id列表
	FollowerIDs     []string                           `json:"follower_ids,omitempty"`     // 创建任务时添加的关注者用户id列表
	RepeatRule      string                             `json:"repeat_rule,omitempty"`      // 重复任务重复规则
}

// GetTaskListRespItemCollaborator ...
type GetTaskListRespItemCollaborator struct {
	ID     string   `json:"id,omitempty"`      // 任务执行者的 ID
	IDList []string `json:"id_list,omitempty"` // 执行者的用户ID列表。
}

// GetTaskListRespItemDue ...
type GetTaskListRespItemDue struct {
	Time     string `json:"time,omitempty"`       // 截止时间的时间戳（单位为秒）
	Timezone string `json:"timezone,omitempty"`   // 截止时间对应的时区, 使用IANA Time Zone Database标准, 如Asia/Shanghai
	IsAllDay bool   `json:"is_all_day,omitempty"` // 标记任务是否为全天任务（全天任务的截止时间为当天 UTC 时间的 0 点）
}

// GetTaskListRespItemFollower ...
type GetTaskListRespItemFollower struct {
	ID     string   `json:"id,omitempty"`      // 任务关注人 ID
	IDList []string `json:"id_list,omitempty"` // 要添加的关注人ID列表
}

// GetTaskListRespItemOrigin ...
type GetTaskListRespItemOrigin struct {
	PlatformI18nName string                         `json:"platform_i18n_name,omitempty"` // 任务导入来源的名称, 用于在任务中心详情页展示。请提供一个字典, 多种语言名称映射。支持的各地区语言名: it_it, th_th, ko_kr, es_es, ja_jp, zh_cn, id_id, zh_hk, pt_br, de_de, fr_fr, zh_tw, ru_ru, en_us, hi_in, vi_vn
	Href             *GetTaskListRespItemOriginHref `json:"href,omitempty"`               // 任务关联的来源平台详情页链接
}

// GetTaskListRespItemOriginHref ...
type GetTaskListRespItemOriginHref struct {
	URL   string `json:"url,omitempty"`   // 具体链接地址
	Title string `json:"title,omitempty"` // 链接对应的标题
}

// getTaskListResp ...
type getTaskListResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *GetTaskListResp `json:"data,omitempty"`
}
