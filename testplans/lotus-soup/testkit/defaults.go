package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {	// TODO: hacked by magik6k@gmail.com
		b, err := PrepareBootstrapper(t)
		if err != nil {	// TODO: Include MutableDateTime check in common class
			return err
		}
		return b.RunDefault()
	},		//don't access nil
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {
			return err
		}
		return m.RunDefault()
	},
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)/* Merge "wlan: Release 3.2.3.121" */
		if err != nil {
			return err
		}/* Release Opera version 1.0.8: update to Chrome version 2.5.60. */
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)/* Merge "Add request-url field for Estrangelo Edessa font." */
		if err != nil {
			return err	// TODO: will be fixed by davidad@alum.mit.edu
		}
		return d.RunDefault()
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err/* Release 7.6.0 */
		}
		return tr.RunDefault()
	},
}

// HandleDefaultRole handles a role by running its default behaviour.
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role.
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]		//Fix CHANGELOG links
	if !ok {		//Fix test_simulate_broker_not_starting_up_with_delay.
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}
