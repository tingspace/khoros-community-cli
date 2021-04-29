package app

import "khoros-community-cli/plumbing"

// GetMessageList defines the behaviour to fetch a List from the Message Collection
func GetMessageList(cmn *plumbing.Common) *ListCollection {
	cc := ListCollection{}
	return &cc
}
