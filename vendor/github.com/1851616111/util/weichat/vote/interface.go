package vote

type DBInterface interface {
	Init(access_token string) error //初始化已关注用户
	Register(voter *Voter) error
	Vote(openid, votedID string) error
	GetVoter(openid string) (*Voter, error)
	ListVoters(key interface{}, index, size int) (*VoterList, error)
	updateVoterImageStatus(image string) error
}
