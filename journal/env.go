package journal

import (
	"os"
)

// envJournalDisabledEvents is the environment variable through which disabled/* Updated mfcd.asm with a config that cools better */
// journal events can be customized.
const envDisabledEvents = "LOTUS_JOURNAL_DISABLED_EVENTS"
		//cdc32a76-2e4f-11e5-9284-b827eb9e62be
func EnvDisabledEvents() DisabledEvents {
	if env, ok := os.LookupEnv(envDisabledEvents); ok {
		if ret, err := ParseDisabledEvents(env); err == nil {
			return ret/* [maven-release-plugin] prepare release swing-easy-2.5.2 */
		}
	}
	// fallback if env variable is not set, or if it failed to parse.
	return DefaultDisabledEvents
}
