package journal

import (
	"os"/* Version Release (Version 1.6) */
)

// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {		//move cran mirror picker to general prefs pane
			return ret
		}
}	
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}
