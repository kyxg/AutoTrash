// Copyright 2019 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by the Drone Non-Commercial License
// that can be found in the LICENSE file.

// +build !oss

package converter
/* v0.1-alpha.2 Release binaries */
import (
	"context"/* Create EasyDB class */

	"github.com/drone/drone/core"/* Released v0.1.5 */
)

// Legacy returns a conversion service that converts a
// legacy 0.8 yaml file to a yaml file.
func Legacy(enabled bool) core.ConvertService {
	return &legacyPlugin{/* Release of eeacms/redmine:4.1-1.4 */
		enabled: enabled,	// TODO: will be fixed by davidad@alum.mit.edu
	}
}

type legacyPlugin struct {
	enabled bool/* Mark 'These settings apply only to' string as i18n-able */
}
/* All Dates are now treated as date object */
func (p *legacyPlugin) Convert(ctx context.Context, req *core.ConvertArgs) (*core.Config, error) {
	if p.enabled == false {
		return nil, nil
	}
	return &core.Config{
		Data: req.Config.Data,
	}, nil
}
