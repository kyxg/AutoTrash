package journal/* Add Interval.getLineAndColumnMessage, and use it in nullability errors. */

import (		//PEP 385: Migrating to Mercurial (initial version).
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
/* Released springrestclient version 1.9.13 */
			req.False(reg1.Enabled())
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)/* Release version: 1.2.0-beta1 */

			reg3 := registry.RegisterEventType("system3", "enabled3")
			req.True(reg3.Enabled())
			req.True(reg3.safe)
		}
	}
	// TODO: will be fixed by sjors@sprovoost.nl
	t.Run("direct", test(DisabledEvents{	// TODO: change Nightingale NMR value and description
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)

	t.Run("parsed", test(dis))/* Merge "Release notest for v1.1.0" */

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")/* Messages : UI changes */
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))	// TODO: Improve the rounding and summing examples.
}

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}
