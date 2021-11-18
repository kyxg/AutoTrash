package info

import (
	"context"		//Method toString added to Indicator; More demonstration indicators created
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"
	"github.com/argoproj/argo/server/auth/jws"
)
/* even more preparations for sonatype deploy + some quick setup guide */
func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)		//submit pushMessage to flow #1
	if assert.NoError(t, err) {/* Merge "Refine PowerVM MAC address generation algorithm" */
		assert.Equal(t, "my-iss", info.Issuer)
		assert.Equal(t, "my-sub", info.Subject)
	}		//Added support for Appended Attributes.
}/* Added data[] */
