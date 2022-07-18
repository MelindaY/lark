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

// DeleteUser 该接口向通讯录删除一个用户信息, 可以理解为员工离职。[常见问题答疑](https://open.feishu.cn/document/ugTN1YjL4UTN24CO1UjN/uQzN1YjL0cTN24CN3UjN)。
//
// 若用户归属部门A、部门B, 应用的通讯录权限范围必须包括部门A和部门B才可以删除用户。应用商店应用无权限调用接口。用户可以在删除员工时设置删除员工数据的接收者, 如果不设置则由其leader接收, 如果该员工没有leader, 则会将该员工的数据删除。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/user/delete
func (r *ContactService) DeleteUser(ctx context.Context, request *DeleteUserReq, options ...MethodOptionFunc) (*DeleteUserResp, *Response, error) {
	if r.cli.mock.mockContactDeleteUser != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#DeleteUser mock enable")
		return r.cli.mock.mockContactDeleteUser(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "DeleteUser",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/users/:user_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteUserResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactDeleteUser mock ContactDeleteUser method
func (r *Mock) MockContactDeleteUser(f func(ctx context.Context, request *DeleteUserReq, options ...MethodOptionFunc) (*DeleteUserResp, *Response, error)) {
	r.mockContactDeleteUser = f
}

// UnMockContactDeleteUser un-mock ContactDeleteUser method
func (r *Mock) UnMockContactDeleteUser() {
	r.mockContactDeleteUser = nil
}

// DeleteUserReq ...
type DeleteUserReq struct {
	UserID                       string                      `path:"user_id" json:"-"`                           // 用户ID, 需要与查询参数中的user_id_type类型保持一致, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	UserIDType                   *IDType                     `query:"user_id_type" json:"-"`                     // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 用户的 open id, union_id: 用户的 union id, user_id: 用户的 user id, 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	DepartmentChatAcceptorUserID *string                     `json:"department_chat_acceptor_user_id,omitempty"` // 部门群接收者。被删除用户为部门群群主时, 转让群主给指定接收者, 不指定接收者则默认转让给群内第一个入群的人, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	ExternalChatAcceptorUserID   *string                     `json:"external_chat_acceptor_user_id,omitempty"`   // 外部群接收者。被删除用户为外部群群主时, 转让群主给指定接收者, 不指定接收者则默认转让给群内与被删除用户在同一组织的第一个入群的人, 如果组织内只有该用户在群里, 则解散外部群, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	DocsAcceptorUserID           *string                     `json:"docs_acceptor_user_id,omitempty"`            // 文档接收者。用户被删除时, 其拥有的文档转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则直接删除文档资源, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	CalendarAcceptorUserID       *string                     `json:"calendar_acceptor_user_id,omitempty"`        // 日程接收者。用户被删除时, 其拥有的日程转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则直接删除日程资源, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	ApplicationAcceptorUserID    *string                     `json:"application_acceptor_user_id,omitempty"`     // 应用接收者。用户被删除时, 其创建的应用转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则不会转移应用, 会造成应用不可用, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	HelpdeskAcceptorUserID       *string                     `json:"helpdesk_acceptor_user_id,omitempty"`        // 服务台资源接收者。用户被删除时, 其拥有的服务台资源转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则直接删除服务台资源, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	MinutesAcceptorUserID        *string                     `json:"minutes_acceptor_user_id,omitempty"`         // 用户妙记接收者。。用户被删除时, 其拥有的妙记资源转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则直接删除妙记资源, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	SurveyAcceptorUserID         *string                     `json:"survey_acceptor_user_id,omitempty"`          // 飞书问卷接收者。用户被删除时, 其拥有的飞书问卷资源转让给接收者, 不指定接收者则默认转让给直接领导, 如果无直接领导则直接删除飞书问卷资源, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
	EmailAcceptor                *DeleteUserReqEmailAcceptor `json:"email_acceptor,omitempty"`                   // 用户邮件资源处理。用户被删除时, 根据传递的操作指令对其拥有的邮件资源做对应处理。未传递指令时默认转让给直接领导, 如果无直接领导则直接保留邮件资源
}

// DeleteUserReqEmailAcceptor ...
type DeleteUserReqEmailAcceptor struct {
	ProcessingType string  `json:"processing_type,omitempty"`  // 资源处理类型, 示例值: "1", 可选值有: 1: 转移资源, 2: 保留资源, 3: 删除资源
	AcceptorUserID *string `json:"acceptor_user_id,omitempty"` // 在 processing_type 为 1 （转移资源时）, 邮件资源接收者, 示例值: "ou_7dab8a3d3cdcc9da365777c7ad535d62"
}

// DeleteUserResp ...
type DeleteUserResp struct {
}

// deleteUserResp ...
type deleteUserResp struct {
	Code int64           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteUserResp `json:"data,omitempty"`
}
