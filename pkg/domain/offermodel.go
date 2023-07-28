package domain

import "time"

type OfferCompany struct {
	OfferID            uint32    `json:"offer_id" gorm:"primaryKey;autoIncrement"`
	ClientID           uint32    `json:"client_id"`
	Country            string    `json:"country" gorm:"not null"`
	Image              string    `json:"image" gorm:"not null"`
	ImageWidth         int32     `json:"image_width"`
	ImageHeight        int32     `json:"image_height"`
	TextLocale         string    `json:"text_locale" gorm:"not null"`
	ValidityTextLocale string    `json:"validity_text_locale" gorm:"not null"`
	Position           int32     `json:"position" gorm:"not null"`
	ValidFrom          time.Time `json:"valid_from" gorm:"not null"`
	ShowFrom           time.Time `json:"show_from"`
	ValidTo            time.Time `json:"valid_to" gorm:"not null"`
	Flag               uint32    `json:"flag" gorm:"not null"`
	PageCount          uint32    `json:"page_count" gorm:"not null"`
	StoreURL           string    `json:"store_url"`
	StoreURLTitle      string    `json:"store_url_title" gorm:"not null"`
	OfferHome          int32     `json:"offer_home" gorm:"not null"`
}
