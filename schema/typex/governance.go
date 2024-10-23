package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=GovernanceType --transform=snake --trimprefix=Governance --output governance_string.go --json --sql
//go:generate go run --mod=mod github.com/rss3-network/enum-schema@v0.1.5 --type=GovernanceType --transform=snake --trimprefix=Governance --output=governance.yaml -t schema.yaml.tmpl
type GovernanceType uint64

//goland:noinspection GoMixedReceiverTypes
func (i GovernanceType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i GovernanceType) Tag() tag.Tag {
	return tag.Governance
}

const (
	GovernanceProposal GovernanceType = iota + 1 // proposal
	GovernanceVote                               // vote
)
