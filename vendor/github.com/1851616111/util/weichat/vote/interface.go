package vote

type DBInterface interface {
	Init() error
	Register(voter *Voter) error
	Vote(openid, votedID string) error
	ListVoters(index, size int) ([]Voter, error)
}
