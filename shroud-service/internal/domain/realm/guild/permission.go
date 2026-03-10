package guild

type PermissionMask uint64

type Permission PermissionMask

const (
	// Server Management
	PermissionAdministrator Permission = 1 << iota
	PermissionManageChannels
	PermissionManageRoles
	PermissionCreateInteractions
	PermissionManageInteractions
	PermissionViewAuditLog
	PermissionManageApplications
	PermissionManageServer
	// Moderation
	PermissionManageProfiles
	PermissionManageUsers // Kick / Approve / Reject
	PermissionBanMembers
	PermissionTimeOutMembers
	PermissionManageThreads
	// User Interaction
	PermissionViewChannels
	PermissionCreateInvite
	PermissionChangeProfile
	PermissionSendMessages
	PermissionCreatePublicThreads
	PermissionCreatePrivateThreads
	PermissionEmbedLinks
	PermissionAttachFiles

	PermissionAddReactions
	PermissionUseExternalInteractions
	PermissionMentionEveryone
	PermissionPinMessages
	PermissionBypassSlowmode
	PermissionReadMessageHistory
	PermissionSendTTSMessages
	PermissionSendVoiceMessages
	PermissionCreatePolls
	// Voice Channels
	PermissionConnect
	PermissionSpeak
	PermissionVideo
	PermissionSoundboard
	PermissionUseVoiceActivity
	PermissionPrioritySpeaker
	PermissionMuteMembers
	PermissionDeafenMembers
	PermissionMoveMembers
	PermissionSetStatus
	// Application
	PermissionUseApplicationCommands
	PermissionActivities
	// Events
	PermissionCreateEvents
	PermissionManageEvents
)
