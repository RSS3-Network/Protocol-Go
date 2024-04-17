package typex

import "github.com/rss3-network/protocol-go/schema/tag"

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=TransactionType --transform=snake --trimprefix=Transaction --output transaction_string.go --json --sql
type TransactionType uint64

//goland:noinspection GoMixedReceiverTypes
func (i TransactionType) Name() string {
	return i.String()
}

//goland:noinspection GoMixedReceiverTypes
func (i TransactionType) Tag() tag.Tag {
	return tag.Transaction
}

const (
	TransactionApproval TransactionType = iota + 1
	TransactionBridge
	TransactionBurn
	TransactionMint
	TransactionTransfer
)
