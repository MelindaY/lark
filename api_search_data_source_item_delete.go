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

// DeleteSearchDataSourceItem 删除数据记录
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source-item/delete
func (r *SearchService) DeleteSearchDataSourceItem(ctx context.Context, request *DeleteSearchDataSourceItemReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceItemResp, *Response, error) {
	if r.cli.mock.mockSearchDeleteSearchDataSourceItem != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#DeleteSearchDataSourceItem mock enable")
		return r.cli.mock.mockSearchDeleteSearchDataSourceItem(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "DeleteSearchDataSourceItem",
		Method:                "DELETE",
		URL:                   r.cli.openBaseURL + "/open-apis/search/v2/data_sources/:data_source_id/items/:item_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(deleteSearchDataSourceItemResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockSearchDeleteSearchDataSourceItem mock SearchDeleteSearchDataSourceItem method
func (r *Mock) MockSearchDeleteSearchDataSourceItem(f func(ctx context.Context, request *DeleteSearchDataSourceItemReq, options ...MethodOptionFunc) (*DeleteSearchDataSourceItemResp, *Response, error)) {
	r.mockSearchDeleteSearchDataSourceItem = f
}

// UnMockSearchDeleteSearchDataSourceItem un-mock SearchDeleteSearchDataSourceItem method
func (r *Mock) UnMockSearchDeleteSearchDataSourceItem() {
	r.mockSearchDeleteSearchDataSourceItem = nil
}

// DeleteSearchDataSourceItemReq ...
type DeleteSearchDataSourceItemReq struct {
	DataSourceID string `path:"data_source_id" json:"-"` // 数据源的ID, 示例值："service_ticket"
	ItemID       string `path:"item_id" json:"-"`        // 数据记录的ID, 示例值："01010111"
}

// deleteSearchDataSourceItemResp ...
type deleteSearchDataSourceItemResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *DeleteSearchDataSourceItemResp `json:"data,omitempty"`
}

// DeleteSearchDataSourceItemResp ...
type DeleteSearchDataSourceItemResp struct {
}