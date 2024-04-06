package model

import "time"

type Vote struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Title       string    `gorm:"column:title;default:NULL"`
	Type        int32     `gorm:"column:type;default:NULL;comment:'0 for single, 1 for multiple choice'"`
	Status      int32     `gorm:"column:status;default:NULL;comment:'0 for normal, 1 for expired'"`
	UserId      int64     `gorm:"column:user_id;default:NULL;comment:'who created'"`
	ExpiredIn   int64     `gorm:"column:expired_in;default:NULL"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *Vote) TableName() string {
	return "vote"
}

type VoteOpt struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name        string    `gorm:"column:name;default:NULL;comment:'选项名称'"`
	Count       int32     `gorm:"column:count;default:NULL;comment:'选项投票计数器'"`
	VoteId      int64     `gorm:"column:vote_id;default:NULL;comment:'绑定到的投票单'"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *VoteOpt) TableName() string {
	return "vote_opt"
}

type VoteOptUser struct {
	Id          int64     `gorm:"column:id;primary_key;NOT NULL"`
	VoteId      int64     `gorm:"column:vote_id;default:NULL;comment:'表单ID'"`
	UserId      int64     `gorm:"column:user_id;default:NULL;comment:'投票用户ID'"`
	VoteOptId   int64     `gorm:"column:vote_opt_id;default:NULL;comment:'选项ID'"`
	CreatedTime time.Time `gorm:"column:created_time;default:NULL"`
	UpdatedTime time.Time `gorm:"column:updated_time;default:NULL"`
}

func (v *VoteOptUser) TableName() string {
	return "vote_opt_user"
}
