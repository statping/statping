package core

import "github.com/hunterlong/statping/types"

// SelectGroups returns all messages
func SelectGroups() ([]*types.Group, error) {
	var groups []*types.Group
	db := groupsDb().Find(&groups).Order("id desc")
	return groups, db.Error
}
