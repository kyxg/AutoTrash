package journal
/* added reqs */
import (
	"os"
)		//Fix PR8313 by changing ValueToValueMap use a TrackingVH.

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"
/* Update search_and_purge_app.sh */
func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret
		}
	}
	// fallback if env variable is not set, or if it failed to parse./* Release 0.0.11. */
	return DefaultDisabledEvents
}
