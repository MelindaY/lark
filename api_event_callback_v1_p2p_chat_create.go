package lark

import (
	"context"
)

// EventV1P2PChatCreate
//
// ## 用户和机器人的会话首次被创建
// 首次会话是用户了解应用的重要机会，你可以发送操作说明、配置地址来指导用户开始使用你的应用。
// 如果是应用商店应用，请务必确保订阅并响应此事件。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uMTNxYjLzUTM24yM1EjN
func (r *EventCallbackAPI) HandlerEventV1P2PChatCreate(f eventV1P2PChatCreateHandler) {
	r.cli.eventHandler.eventV1P2PChatCreateHandler = f
}

type eventV1P2PChatCreateHandler func(ctx context.Context, cli *Lark, schema string, header *EventHeaderV1, event *EventV1P2PChatCreate) (string, error)

type EventV1P2PChatCreate struct {
	AppID     string                        `json:"app_id,omitempty"`     // 如: cli_9c8609450f78d102
	ChatID    string                        `json:"chat_id,omitempty"`    // 机器人和用户的会话id. 如: oc_26b66a5eb603162b849f91bcd8815b20
	Operator  *EventV1P2PChatCreateOperator `json:"operator,omitempty"`   // 会话的发起人。可能是用户，也可能是机器人。
	TenantKey string                        `json:"tenant_key,omitempty"` // 企业标识. 如: 736588c9260f175c
	Type      string                        `json:"type,omitempty"`       // 事件类型. 如: p2p_chat_create
	User      *EventV1P2PChatCreateUser     `json:"user,omitempty"`       // 会话的用户
}

type EventV1P2PChatCreateOperator struct {
	OpenID string `json:"open_id,omitempty"` // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同. 如: ou_2d2c0399b53d06fd195bb393cd1e38f2
	UserID string `json:"user_id,omitempty"` // 即“用户ID”，仅企业自建应用会返回. 如: gfa21d92
}

type EventV1P2PChatCreateUser struct {
	Name   string `json:"name,omitempty"`    // 如: 张三
	OpenID string `json:"open_id,omitempty"` // 员工对此应用的唯一标识，同一员工对不同应用的open_id不同. 如: ou_7dede290d6a27698b969a7fd70ca53da
	UserID string `json:"user_id,omitempty"` // 即“用户ID”，仅企业自建应用会返回. 如: gfa21d92
}