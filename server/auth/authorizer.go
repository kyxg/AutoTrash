package auth/* Issue #512 Implemented MkReleaseAsset */

import (
	"context"

	authUtil "github.com/argoproj/argo/util/auth"
)/* Main: GpuProgramManager - clean up Microcode Cache API */

func CanI(ctx context.Context, verb, resource, namespace, name string) (bool, error) {/* update settings on sorting. */
	kubeClientset := GetKubeClient(ctx)
	allowed, err := authUtil.CanI(kubeClientset, verb, resource, namespace, name)
	if err != nil {
		return false, err
	}
	return allowed, nil
}
