package nip29

import (
	"slices"

	"github.com/nbd-wtf/go-nostr"
)

type Role struct {
	Name        string
	Permissions map[Permission]struct{}
}

type Permission string

const (
	PermAddUser          Permission = "add-user"
	PermEditMetadata     Permission = "edit-metadata"
	PermDeleteEvent      Permission = "delete-event"
	PermRemoveUser       Permission = "remove-user"
	PermAddPermission    Permission = "add-permission"
	PermRemovePermission Permission = "remove-permission"
	PermEditGroupStatus  Permission = "edit-group-status"
)

type KindRange []int

var ModerationEventKinds = KindRange{
	nostr.KindSimpleGroupAddUser,
	nostr.KindSimpleGroupRemoveUser,
	nostr.KindSimpleGroupEditMetadata,
	nostr.KindSimpleGroupAddPermission,
	nostr.KindSimpleGroupRemovePermission,
	nostr.KindSimpleGroupDeleteEvent,
	nostr.KindSimpleGroupEditGroupStatus,
	nostr.KindSimpleGroupCreateGroup,
	nostr.KindSimpleGroupDeleteGroup,
}

var MetadataEventKinds = KindRange{
	nostr.KindSimpleGroupMetadata,
	nostr.KindSimpleGroupAdmins,
	nostr.KindSimpleGroupMembers,
}

const JoinRequestKind = nostr.KindSimpleGroupJoinRequest

func (kr KindRange) Includes(kind int) bool {
	_, ok := slices.BinarySearch(kr, kind)
	return ok
}

var (
	// used for normal members without admin powers
	EmptyRole *Role = nil

	PermissionsMap = map[Permission]struct{}{
		PermAddUser:          {},
		PermEditMetadata:     {},
		PermDeleteEvent:      {},
		PermRemoveUser:       {},
		PermAddPermission:    {},
		PermRemovePermission: {},
		PermEditGroupStatus:  {},
	}
)
