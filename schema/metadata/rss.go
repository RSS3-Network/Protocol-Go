package metadata

import (
	"github.com/rss3-network/protocol-go/schema"
	"github.com/rss3-network/protocol-go/schema/typex"
)

var _ Metadata = (*RSS)(nil)

type RSS struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	PubDate     string    `json:"pub_date,omitempty"`
	Authors     []Authors `json:"authors,omitempty"`
}

type Authors struct {
	Name string `json:"name"`
}

func (r RSS) Type() schema.Type {
	return typex.RSSFeed
}
