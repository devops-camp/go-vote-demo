package model

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

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

// GetVotes 获取投票列表
func GetVotes() ([]Vote, error) {
	votes := make([]Vote, 0)
	tx := Conn.Table("vote").Find(&votes)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return votes, nil
}

// GetVote 根据 ID 查询数据
func GetVote(id int64) (Vote, error) {
	vote := Vote{}
	tx := Conn.Table("vote").Where("id = ?", id).First(&vote)
	if tx.Error != nil {
		return Vote{}, tx.Error
	}

	if tx.RowsAffected == 0 {
		return Vote{}, fmt.Errorf("Record not found")
	}

	return vote, nil
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

// GetVoteOptsByVoteId 根据投票 ID 查询选项
func GetVoteOptsByVoteId(voteId int64) ([]VoteOpt, error) {
	voteOpts := make([]VoteOpt, 0)
	tx := Conn.Table("vote_opt").Where("vote_id = ?", voteId).Find(&voteOpts)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return voteOpts, nil
}

// UpdateVoteCount 更新 VoteOpt 表计数器
func UpdateVoteCount(id int64, voteId int64) error {
	tx := Conn.Table("vote_opt").
		Where("id = ? AND vote_id = ?", id, voteId).
		Update("count", gorm.Expr("count + ?", 1))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
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
