/* Written by Dave Richards.
 *
 * This the interface for implementing all of the magic, the data store.  Knock yourself out, write a plugin.
 */
package discodove_interface_datastore


// Representaiton of a mail folder
// A UidValidity of zero(0) means it's invalid
type DiscoDoveMailFolder interface { 
	Name() string
	URI() string
	UidValidity() (uint32, error)
	Messages() ([]DiscoDoveMailMessage, error)
}

const (
	FlagSeenMessage			= `\Seen`
	FlagAnsweredMessage		= `\Answered`
	FlagDeletedMessage		= `\Deleted`
	FlagDraftMessage		= `\Draft`
	FlagFlaggedMessage 		= `\Flagged`
	FlagRecentMessage 		= `\Recent`
)

// Representation of a message
type DiscoDoveMailMessage interface {
	Uid() uint32
	URI() string
	RawMsg() string
	IMAPSize() uint
	IsNew() bool
	SetFlag(flag string) error
	UnSetFlag(flag string) error
	GetFlags() (flags []string, err error)
}