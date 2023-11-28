package campaign

import (
	"time"

	"gonasiku.com/user"
)

type Campaign struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	Title       string    `json:"title" gorm:"type:varchar(255)"`
	Description string    `json:"description" gorm:"type:text"`
	StartDate   time.Time `json:"startDate" gorm:"type:date"`
	EndDate     time.Time `json:"endDate" gorm:"type:date"`
	Goal        uint      `json:"goal" gorm:"type:int"`

	CreatedAt time.Time `json:"createdAt" gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"not null;type:datetime;default:CURRENT_TIMESTAMP"`

	CreatorID   uint             `json:"creatorID"`
	CampaignImg []CampaignImages `json:"campaignImg"`
	Creator     user.User        `json:"creator" gorm:"foreignKey:CreatorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type CampaignImages struct {
	ID         uint   `json:"id" gorm:"primary_key"`
	CampaignID uint   `json:"campaignID"`
	Image      string `json:"image" gorm:"type:varchar(255)"`
	SetActive  bool   `json:"setActive" gorm:"type:bool"`
}
