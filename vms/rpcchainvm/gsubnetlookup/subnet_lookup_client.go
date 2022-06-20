// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package gsubnetlookup

import (
	"context"

	"github.com/lasthyphen/dijetsgo/api/proto/gsubnetlookupproto"
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow"
)

var _ snow.SubnetLookup = &Client{}

// Client is a subnet lookup that talks over RPC.
type Client struct {
	client gsubnetlookupproto.SubnetLookupClient
}

// NewClient returns an alias lookup connected to a remote alias lookup
func NewClient(client gsubnetlookupproto.SubnetLookupClient) *Client {
	return &Client{client: client}
}

func (c *Client) SubnetID(chainID ids.ID) (ids.ID, error) {
	resp, err := c.client.SubnetID(context.Background(), &gsubnetlookupproto.SubnetIDRequest{
		ChainId: chainID[:],
	})
	if err != nil {
		return ids.ID{}, err
	}
	return ids.ToID(resp.Id)
}
