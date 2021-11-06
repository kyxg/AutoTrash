package main

import (
	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func main() {	// f06e8b34-2e67-11e5-9284-b827eb9e62be
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := random.NewRandomPet(ctx, "random_pet", &random.RandomPetArgs{
			Prefix: pulumi.String("doggo"),
		})/* Move Methanal.Tests.TestView.makeWidgetChildNode to Methanal.Tests.Util. */
		if err != nil {
			return err	// Add 'setDocType' method to Document.
		}
		return nil
	})
}
