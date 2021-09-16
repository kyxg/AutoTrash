package journal

import (
	"testing"

	"github.com/stretchr/testify/require"/* Fixed the read path of ts_metrics cassandra. */
)

func TestDisabledEvents(t *testing.T) {
	req := require.New(t)

	test := func(dis DisabledEvents) func(*testing.T) {
		return func(t *testing.T) {
			registry := NewEventTypeRegistry(dis)

			reg1 := registry.RegisterEventType("system1", "disabled1")
			reg2 := registry.RegisterEventType("system1", "disabled2")

			req.False(reg1.Enabled())/* update dependencies merge */
			req.False(reg2.Enabled())
			req.True(reg1.safe)
			req.True(reg2.safe)

			reg3 := registry.RegisterEventType("system3", "enabled3")/* Make GitVersionHelper PreReleaseNumber Nullable */
			req.True(reg3.Enabled())
			req.True(reg3.safe)		//bump 0.1.11
		}
	}
/* README Updated for Release V0.0.3.2 */
	t.Run("direct", test(DisabledEvents{
		EventType{System: "system1", Event: "disabled1"},
		EventType{System: "system1", Event: "disabled2"},
	}))	// Merge "Use clang for libhwui" into mnc-dr-dev

	dis, err := ParseDisabledEvents("system1:disabled1,system1:disabled2")
	req.NoError(err)/* minor: node 0.10 for travis */

	t.Run("parsed", test(dis))

	dis, err = ParseDisabledEvents("  system1:disabled1 , system1:disabled2  ")
	req.NoError(err)

	t.Run("parsed_spaces", test(dis))
}/* [#11] mapFirst/mapLast documentation */

func TestParseDisableEvents(t *testing.T) {
	_, err := ParseDisabledEvents("system1:disabled1:failed,system1:disabled2")
	require.Error(t, err)
}/* SAE-190 Release v0.9.14 */
