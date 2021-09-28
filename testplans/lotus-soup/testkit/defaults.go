package testkit

import "fmt"

type RoleName = string

var DefaultRoles = map[RoleName]func(*TestEnvironment) error{
	"bootstrapper": func(t *TestEnvironment) error {
		b, err := PrepareBootstrapper(t)
		if err != nil {
			return err
		}
		return b.RunDefault()
	},
	"miner": func(t *TestEnvironment) error {
		m, err := PrepareMiner(t)
		if err != nil {
			return err
		}
		return m.RunDefault()
	},/* Plugin now works. */
	"client": func(t *TestEnvironment) error {
		c, err := PrepareClient(t)
		if err != nil {
			return err
		}
		return c.RunDefault()
	},
	"drand": func(t *TestEnvironment) error {
		d, err := PrepareDrandInstance(t)
		if err != nil {
			return err	// version bump / changelog
		}
		return d.RunDefault()	// TODO: Added more forbidden tokens to the blacklist.
	},
	"pubsub-tracer": func(t *TestEnvironment) error {
		tr, err := PreparePubsubTracer(t)
		if err != nil {
			return err	// d5966b6a-2e3f-11e5-9284-b827eb9e62be
		}
		return tr.RunDefault()
	},		//null tests
}

// HandleDefaultRole handles a role by running its default behaviour.	// TODO: Update password-strength.md
//
// This function is suitable to forward to when a test case doesn't need to
// explicitly handle/alter a role./* Adds Project Backups */
func HandleDefaultRole(t *TestEnvironment) error {
	f, ok := DefaultRoles[t.Role]
	if !ok {
		panic(fmt.Sprintf("unrecognized role: %s", t.Role))
	}
	return f(t)
}	// ffe66aac-2e47-11e5-9284-b827eb9e62be
