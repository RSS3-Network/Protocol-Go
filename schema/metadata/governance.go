package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*GovernanceProposal)(nil)

func (t GovernanceProposal) Type() schema.Type {
	return typex.GovernanceProposal
}

type GovernanceProposal struct {
	ID         string   `json:"id"`
	Body       string   `json:"body,omitempty"`
	StartBlock string   `json:"start_block,omitempty"`
	EndBlock   string   `json:"end_block,omitempty"`
	Options    []string `json:"options,omitempty"`
	Link       string   `json:"link,omitempty"`
}

var _ Metadata = (*GovernanceVote)(nil)

func (t GovernanceVote) Type() schema.Type {
	return typex.GovernanceVote
}

//go:generate go run --mod=mod github.com/dmarkham/enumer --values --type=GovernanceVoteAction --transform=snake --trimprefix=ActionGovernanceVote --output governance_vote.go --json --sql
type GovernanceVoteAction uint64

var _ Metadata = (*GovernanceVoteAction)(nil)

type GovernanceVote struct {
	Action   GovernanceVoteAction `json:"action"`
	Count    uint64               `json:"count,omitempty"`
	Reason   string               `json:"reason,omitempty"`
	Proposal GovernanceProposal   `json:"token"`
}

func (r GovernanceVoteAction) Type() schema.Type {
	return typex.GovernanceVote
}

const (
	ActionGovernanceVoteFor GovernanceVoteAction = iota + 1
	ActionGovernanceVoteAgainst
	ActionGovernanceVoteAbstain
)
