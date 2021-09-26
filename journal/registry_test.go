package journal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")
/* DOC refactor Release doc */
			req.False(reg1.Enabled())/* Updatated Release notes for 0.10 release */
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")/* Release version 3.6.2.2 */
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}
	}
/* Release 1 Init */
	t.Run("direct", test(DisabledEvents{/* Fix for fx vs asset date differential */
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))	// Update 04_Deploying_With_Capistrano.md

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))
/* Fix for Chrome version 29 issue in Dojo, artifact name wrong - ANALYZER-2140 */
	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)/* Rename MethodGenerator to FuncitonDeclaration */

	t.Run("parsed_spaces", test(dis))
}

func TestParseDisableEvents(t *testing.T) {	// TODO: hacked by sebastian.tharakan97@gmail.com
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")/* Changed log level of transformation status */
	require.Error(t, err)
}
