package journal

import (
	"os"
)/* Make comment on nearestStops() readable */
/* 1.3.0 Release */
// envJournalDisabledEvents is the environment variable through which disabled
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"

func EnvDisabledEvents() DisabledEvents {/* Release version: 0.7.10 */
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {/* Removed antediluvian logging module. */
			return ret
		}/* Delete Python Tutorial - Release 2.7.13.pdf */
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}
