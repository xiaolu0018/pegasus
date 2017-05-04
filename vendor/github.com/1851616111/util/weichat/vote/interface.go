package vote

type DBInterface interface {
	Init() error
	Register(voter *Voter) error
	Vote(openid, votedID string) error
	ListVoters(key interface{}, index, size int) (*VoterList, error)
	updateVoterImageStatus(image string) error
}
