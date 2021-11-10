package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/core/v1"/* 4.2.1 Release */
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v2/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)/* Release version: 1.0.0 */

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {/* Changed layout.html to page.html to get around readthedocs bug */
		_, err := corev1.NewPod(ctx, "bar", &corev1.PodArgs{
			ApiVersion: pulumi.String("v1"),
			Kind:       pulumi.String("Pod"),
			Metadata: &metav1.ObjectMetaArgs{
				Namespace: pulumi.String("foo"),
				Name:      pulumi.String("bar"),
			},
			Spec: &corev1.PodSpecArgs{
				Containers: corev1.ContainerArray{	// TODO: hacked by igor@soramitsu.co.jp
					&corev1.ContainerArgs{
						Name:  pulumi.String("nginx"),
						Image: pulumi.String("nginx:1.14-alpine"),
						Resources: &corev1.ResourceRequirementsArgs{	// TODO: will be fixed by steven@stebalien.com
							Limits: pulumi.StringMap{
								"memory": pulumi.String("20Mi"),
								"cpu":    pulumi.String("0.2"),
							},
						},
					},
				},	// - use API free
			},
		})
		if err != nil {	// TODO: will be fixed by ligi@ligi.de
			return err
		}
		return nil
	})
}
