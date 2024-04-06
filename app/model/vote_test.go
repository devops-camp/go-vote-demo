package model

import "testing"

// 初始化 Mysql 链接
func init() {
	NewMysql()
}

// TestGetVotes 测试获取投票列表
func TestGetVotes(t *testing.T) {
	votes, err := GetVotes()

	if err != nil {
		t.Errorf("GetVotes() failed: %v", err)
	}

	//t.Logf("votes: %v", votes)
	for _, v := range votes {
		t.Logf("vote: %v", v)
	}
}
