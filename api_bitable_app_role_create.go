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

// CreateBitableAppRole 新增自定义权限
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-role/create
func (r *BitableService) CreateBitableAppRole(ctx context.Context, request *CreateBitableAppRoleReq, options ...MethodOptionFunc) (*CreateBitableAppRoleResp, *Response, error) {
	if r.cli.mock.mockBitableCreateBitableAppRole != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#CreateBitableAppRole mock enable")
		return r.cli.mock.mockBitableCreateBitableAppRole(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "CreateBitableAppRole",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/roles",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createBitableAppRoleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableCreateBitableAppRole mock BitableCreateBitableAppRole method
func (r *Mock) MockBitableCreateBitableAppRole(f func(ctx context.Context, request *CreateBitableAppRoleReq, options ...MethodOptionFunc) (*CreateBitableAppRoleResp, *Response, error)) {
	r.mockBitableCreateBitableAppRole = f
}

// UnMockBitableCreateBitableAppRole un-mock BitableCreateBitableAppRole method
func (r *Mock) UnMockBitableCreateBitableAppRole() {
	r.mockBitableCreateBitableAppRole = nil
}

// CreateBitableAppRoleReq ...
type CreateBitableAppRoleReq struct {
	AppToken   string                              `path:"app_token" json:"-"`    // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	RoleName   string                              `json:"role_name,omitempty"`   // 自定义权限的名字, 示例值："自定义权限1"
	TableRoles []*CreateBitableAppRoleReqTableRole `json:"table_roles,omitempty"` // 数据表权限, 最大长度：`100`
}

// CreateBitableAppRoleReqTableRole ...
type CreateBitableAppRoleReqTableRole struct {
	TableName string                                   `json:"table_name,omitempty"` // 数据表名, 示例值："数据表1"
	TablePerm int64                                    `json:"table_perm,omitempty"` // 数据表权限，`协作者可编辑自己的记录`和`可编辑指定字段`是`可编辑记录`的特殊情况，可通过指定`rec_rule`或`field_perm`参数实现相同的效果, 示例值：0, 可选值有: `0`：无权限, `1`：可阅读, `2`：可编辑记录, `4`：可编辑字段和记录, 默认值: `0`
	RecRule   *CreateBitableAppRoleReqTableRoleRecRule `json:"rec_rule,omitempty"`   // 记录筛选条件，在table_perm为1或2时有意义，用于指定可编辑或可阅读某些记录
	FieldPerm map[string]int64                         `json:"field_perm,omitempty"` // 字段权限，仅在table_perm为2时有意义，设置字段可编辑或可阅读。类型为 map，key 是字段名，value 是字段权限。,**value 枚举值有：**, `1`：可阅读, `2`：可编辑
}

// CreateBitableAppRoleReqTableRoleRecRule ...
type CreateBitableAppRoleReqTableRoleRecRule struct {
	Conditions  []*CreateBitableAppRoleReqTableRoleRecRuleCondition `json:"conditions,omitempty"`  // 记录筛选条件, 最大长度：`100`
	Conjunction *string                                             `json:"conjunction,omitempty"` // 多个筛选条件的关系, 示例值："and", 可选值有: `and`：与, `or`：或, 默认值: `and`
	OtherPerm   *int64                                              `json:"other_perm,omitempty"`  // 其他记录权限，仅在table_perm为2时有意义, 示例值：0, 可选值有: `0`：禁止查看, `1`：仅可阅读, 默认值: `0`
}

// CreateBitableAppRoleReqTableRoleRecRuleCondition ...
type CreateBitableAppRoleReqTableRoleRecRuleCondition struct {
	FieldName string   `json:"field_name,omitempty"` // 字段名，记录筛选条件是`创建人包含访问者本人`时，此参数值为"", 示例值："单选"
	Operator  *string  `json:"operator,omitempty"`   // 运算符, 示例值："is", 可选值有: `is`：等于, `isNot`：不等于, `contains`：包含, `doesNotContain`：不包含, `isEmpty`：为空, `isNotEmpty`：不为空, 默认值: `is`
	Value     []string `json:"value,omitempty"`      // 单选或多选字段的选项id, 示例值：["optbdVHf4q", "optrpd3eIJ"]
}

// createBitableAppRoleResp ...
type createBitableAppRoleResp struct {
	Code int64                     `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                    `json:"msg,omitempty"`  // 错误描述
	Data *CreateBitableAppRoleResp `json:"data,omitempty"`
}

// CreateBitableAppRoleResp ...
type CreateBitableAppRoleResp struct {
	Role *CreateBitableAppRoleRespRole `json:"role,omitempty"` // 自定义权限
}

// CreateBitableAppRoleRespRole ...
type CreateBitableAppRoleRespRole struct {
	RoleName   string                                   `json:"role_name,omitempty"`   // 自定义权限的名字
	RoleID     string                                   `json:"role_id,omitempty"`     // 自定义权限的id
	TableRoles []*CreateBitableAppRoleRespRoleTableRole `json:"table_roles,omitempty"` // 数据表权限
}

// CreateBitableAppRoleRespRoleTableRole ...
type CreateBitableAppRoleRespRoleTableRole struct {
	TableName string                                        `json:"table_name,omitempty"` // 数据表名
	TablePerm int64                                         `json:"table_perm,omitempty"` // 数据表权限，`协作者可编辑自己的记录`和`可编辑指定字段`是`可编辑记录`的特殊情况，可通过指定`rec_rule`或`field_perm`参数实现相同的效果, 可选值有: `0`：无权限, `1`：可阅读, `2`：可编辑记录, `4`：可编辑字段和记录
	RecRule   *CreateBitableAppRoleRespRoleTableRoleRecRule `json:"rec_rule,omitempty"`   // 记录筛选条件，在table_perm为1或2时有意义，用于指定可编辑或可阅读某些记录
	FieldPerm map[string]int64                              `json:"field_perm,omitempty"` // 字段权限，仅在table_perm为2时有意义，设置字段可编辑或可阅读。类型为 map，key 是字段名，value 是字段权限。,**value 枚举值有：**, `1`：可阅读, `2`：可编辑
}

// CreateBitableAppRoleRespRoleTableRoleRecRule ...
type CreateBitableAppRoleRespRoleTableRoleRecRule struct {
	Conditions  []*CreateBitableAppRoleRespRoleTableRoleRecRuleCondition `json:"conditions,omitempty"`  // 记录筛选条件
	Conjunction string                                                   `json:"conjunction,omitempty"` // 多个筛选条件的关系, 可选值有: `and`：与, `or`：或
	OtherPerm   int64                                                    `json:"other_perm,omitempty"`  // 其他记录权限，仅在table_perm为2时有意义, 可选值有: `0`：禁止查看, `1`：仅可阅读
}

// CreateBitableAppRoleRespRoleTableRoleRecRuleCondition ...
type CreateBitableAppRoleRespRoleTableRoleRecRuleCondition struct {
	FieldName string   `json:"field_name,omitempty"` // 字段名，记录筛选条件是`创建人包含访问者本人`时，此参数值为""
	Operator  string   `json:"operator,omitempty"`   // 运算符, 可选值有: `is`：等于, `isNot`：不等于, `contains`：包含, `doesNotContain`：不包含, `isEmpty`：为空, `isNotEmpty`：不为空
	Value     []string `json:"value,omitempty"`      // 单选或多选字段的选项id
	FieldType int64    `json:"field_type,omitempty"` // 字段类型
}