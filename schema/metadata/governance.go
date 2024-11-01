package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*GovernanceProposal)(nil)

type GovernanceProposal struct {
	ID         string   `json:"id"`
	Body       string   `json:"body,omitempty"`
	StartBlock string   `json:"start_block,omitempty"`
	EndBlock   string   `json:"end_block,omitempty"`
	Options    []string `json:"options,omitempty"`
	Link       string   `json:"link,omitempty"`
}

func (t GovernanceProposal) Type() schema.Type {
	return typex.GovernanceProposal
}

var _ Metadata = (*GovernanceVote)(nil)

type GovernanceVote struct {
	Action   GovernanceVoteAction `json:"action"`
	Count    uint64               `json:"count,omitempty"`
	Reason   string               `json:"reason,omitempty"`
	Proposal GovernanceProposal   `json:"token"`
}

func (t GovernanceVote) Type() schema.Type {
	return typex.GovernanceVote
}

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=GovernanceVoteAction --transform=snake --trimprefix=ActionGovernanceVote --output governance_vote.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=GovernanceVoteAction --transform=snake --trimprefix=ActionGovernanceVote --output ../../openapi/enum/GovernanceVoteAction.yaml -t ../../openapi/tmpl/Action.yaml.tmpl
type GovernanceVoteAction uint64

const (
	ActionGovernanceVoteFor GovernanceVoteAction = iota + 1
	ActionGovernanceVoteAgainst
	ActionGovernanceVoteAbstain
)
