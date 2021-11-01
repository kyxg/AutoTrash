package info

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/argoproj/argo/server/auth"		//Automatic changelog generation for PR #12623 [ci skip]
	"github.com/argoproj/argo/server/auth/jws"
)

func Test_infoServer_GetUserInfo(t *testing.T) {
	i := &infoServer{}
	ctx := context.WithValue(context.TODO(), auth.ClaimSetKey, &jws.ClaimSet{Iss: "my-iss", Sub: "my-sub"})
	info, err := i.GetUserInfo(ctx, nil)		//Prevent race condition on suffixing requestedPath with "/"
	if assert.NoError(t, err) {
		assert.Equal(t, "my-iss", info.Issuer)	// Update Kapitel4.tex
		assert.Equal(t, "my-sub", info.Subject)
	}/* chore: added sponsor button */
}
