/* Written by Dave Richards.
 *
 * This the interface for implementing all of the magic, the data store.  Knock yourself out, write a plugin.
 */
package discodove_interface_datastore

import (
	"log/syslog"
	"github.com/spf13/viper"
)


// The factory is used to allocate new sessions for a user, this must be implemented by a plugin writer.
type DiscoDoveDataStoreFactory interface { 

	/* This will be called once when we load this plugin, if you feel compelled to set something up, perhaps a 
	 * control/query/admin thread or something, then do it here in a controlled manner - similarly if 
	 * you want to pool connections, etc....  We assume that each plugin can scale itself, we do no magic
	 * to allow for scalability, so you might want some worker threads.
	 *
	 * Each plugin is responsible for creating it's own syslog connection as *syslog.Writer has a mutex, and 
	 * I don't want the auth threads to be blocking on writing to syslog - so you need to scale logging yourself.
	 * 
	 * We use Viper for config, and i will pass in the config directives for your module, but as it's viper you
	 * can access the entire discodove config too.  Feel free to specify your own config directives.
	 *
	 * name	 	: will be the name of the process, in 99.999% of cases it will just be "discodove" - please
	 *            prefix your log messages with this and perhaps your own identifier e.g. "local_maildir"
	 * syslogFacility : which facility to use in syslog.
	 * conf: a Viper subtree configuration for this service as specified in the discodove config.
	 */
	Initialize(name string, syslogFacility syslog.Priority, conf *viper.Viper) error

	/* Call this to create a new session for a user that is hosted on this data store.
	   options are specified per plugin, e.g. homedir, quota, etc....
	 */
	NewUserSession(user string, options string) (DiscoDoveDataStore, error)
}

/* The plugin represents a single user session.  Multiple of these sessions for a single user might exist 
 * at any one time, so plan accordingly.  
 */
type DiscoDoveDataStore interface {

	// The real work is done here, all the methods for each type of data should be implemented.

	// Name of the Store 
	StoreName() (storename string)

	// Current user name
	CurrentUser() (username string)

	// Name space strings, see RFC 2342
	Namespaces() (personal_namespace, personal_namespace_seperator, other_users_namespace, other_users_namespace_seperator, shared_namespace, shared_namespace_seperator string)

	// Return a list of personal mail folders, and associated attributes.
    PersonalMailFolders() (folders []DiscoDoveMailFolder, err error) 

    // Return a list of folders shared by other users, referred to as "others" in the IMAP specs
    SharedMailFolders() (folders []DiscoDoveMailFolder, err error) 

    // Return a list of public folders, also called "shared" in the IMAP spec - yes, confusing isn't it....
    PublicMailFolders() (folders []DiscoDoveMailFolder, err error) 
}
