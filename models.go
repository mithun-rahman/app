package main

import (
	"time"
)

var (
	Models []interface{}
)

func init() {
	Models = append(Models, &PreAuthSessions{})
}

type PreAuthSessions struct {
	SessionIDExpiry      *time.Time `json:"session_expiry" gorm:"column:session_expiry"`
	MerchantRequestId    string     `json:"merchant_req_id" gorm:"column:merchant_req_id;unique"`
	CallBackUrl          string     `json:"call_back_url" gorm:"column:call_back_url"`
	UserSlug             string     `json:"user_slug" gorm:"column:user_slug;size:40"`
	MaskedNumber         string     `json:"masked_number" gorm:"column:masked_number;size:40"`
	IsParentRequest      bool       `json:"is_parent_req" gorm:"column:is_parent_req"`
	ParentMerchantId     string     `json:"parent_merchant_id" gorm:"column:parent_merchant_id"`
	ChildIdList          string     `json:"child_merchant_id_list" gorm:"column:child_merchant_id_list"`
	TemporaryToken       *string    `json:"temporary_token" gorm:"column:temporary_token;uniqueIndex:idx_temporary_token"`
	TemporaryTokenExpiry *time.Time `json:"temporary_token_expiry" gorm:"column:temporary_token_expiry"`
	Status               string     `json:"status" gorm:"column:status"` //initiated,processing,authenticated
}
