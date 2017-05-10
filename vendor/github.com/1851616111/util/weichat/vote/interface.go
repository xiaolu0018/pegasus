package vote

import "github.com/1851616111/util/weichat/event"

type DBInterface interface {
	Init(access_token string) error //初始化已关注用户
	Register(voter *Voter) error
	Vote(openid, votedID string) error
	GetVoterStatus(openid string) (*VStatus, error)
	ListVoters(key interface{}, index, size int) (*VoterList, error)
	updateVoterImageStatus(image string) error

	Follow(*event.Event) error
	UnFollow(*event.Event) error
}
