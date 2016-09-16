package model

import (
	"log"
	"testing"
	"time"

	"github.com/achiku/mergo"
)

// TestCreateUserAccountData create user account test data
func TestCreateUserAccountData(t *testing.T, tx Query, ua *UserAccount) *UserAccount {
	accountDefault := &UserAccount{
		Birthday:     time.Date(1985, 8, 18, 0, 0, 0, 0, time.Local),
		Email:        "test@example.com",
		Gender:       "male",
		Password:     "pass",
		RegisteredAt: time.Now(),
	}
	if err := mergo.MergeWithOverwrite(accountDefault, ua, TestStructMergeFunc); err != nil {
		log.Fatal(err)
	}
	if err := accountDefault.Create(tx); err != nil {
		t.Fatal(err)
	}
	return accountDefault
}
