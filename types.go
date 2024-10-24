package main

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type LoginRequest struct {
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TransferRequest struct {
	ToAccount string `json:"to_account"`
	Amount    string `json:"amount"`
}

type Account struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Number    int64     `json:"number"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		FirstName: firstName,
		LastName:  lastName,
		Number:    int64(rand.Intn(1000000)),
		CreatedAt: time.Now().UTC(),
	}
}

type UserWithJSON struct {
	LastStep string                    `json:"last_step"`
	CardInfo datatypes.JSONType[Steps] `json:"card_info"`
}

type CardInfo struct {
	Card     Card     `json:"1"`
	Personal Personal `json:"2"`
	Address  Address  `json:"3"`
	Nominee  Nominee  `json:"4"`
}

type Card struct {
	CardTheme string `json:"card_theme"`
}

type Personal struct {
	CardName      string `json:"card_name"`
	Occupation    string `json:"occupation"`
	MonthlyIncome int64  `json:"monthly_income"`
	MaritalStatus string `json:"marital_status"`
	Religion      string `json:"religion"`
	Signature     string `json:"signature"`
}

type Address struct {
	PresentAddress     string `json:"present_address"`
	PresentAdditional  string `json:"pa_additional_details"`
	PresentPostalCode  string `json:"pa_postcode"`
	ShippingAddress    string `json:"shipping_address"`
	ShippingAdditional string `json:"sa_additional_details"`
	ShippingPostalCode string `json:"sa_postcode"`
	IsSameAddress      bool   `json:"is_same_address"`
}

type Nominee struct {
	Relation        string `json:"relation"`
	PhoneNumber     string `json:"phone_number"`
	NomineePhoto    string `json:"nominee_photo"`
	NomineeNidFront string `json:"nominee_nid_front"`
	NomineeNidBack  string `json:"nominee_nid_back"`
}

type CardApplication struct {
	DefaultModel
	UserSlug      string   `json:"user_slug" gorm:"size:40"`
	ApplicationID string   `json:"application_id" gorm:"size:40"`
	UserPhone     string   `json:"user_phone" gorm:"size:20"`
	Status        string   `json:"status" gorm:"size:20"`
	LastStep      string   `json:"last_step" gorm:"size:20"`
	CardData      CardInfo `gorm:"type:jsonb"`
	OtpRef        string   `json:"otp_ref" gorm:"size:40"`
	Remarks       string   `json:"remark" gorm:"size:255"`
	TicketID      string   `json:"ticket_id" gorm:"size:40"`
	Retry         int      `json:"retry" gorm:"column:retry"`
	ExpiredUntil  int64    `json:"expired_until" gorm:"column:expired_until"`
	Toc           bool     `json:"toc" gorm:"toc"`
}

type DefaultModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primarykey;default:public.uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at" gorm:"index:,type:brin"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"index:,type:brin"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index:,type:brin"`
}

type TwoFa struct {
	DefaultModel
	ApplicationID    string `json:"application_id" gorm:"size:40"`
	Token            string `json:"token" gorm:"column:token;size:255;index:idx_two_fa_token"`
	ValidUntil       int64  `json:"valid_until" gorm:"column:valid_until;index"`
	RetryCount       int    `json:"retry_count" gorm:"column:retry_count"`
	ResendCount      int    `json:"resend_count" gorm:"column:resend_count"`
	ResendValidUntil int64  `json:"resend_valid_until" gorm:"column:resend_valid_until;index"`
}
