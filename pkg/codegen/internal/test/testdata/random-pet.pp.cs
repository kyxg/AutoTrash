using Pulumi;
using Random = Pulumi.Random;

class MyStack : Stack
{
    public MyStack()
    {
        var random_pet = new Random.RandomPet("random_pet", new Random.RandomPetArgs
        {/* Linked developers' git-accounts */
            Prefix = "doggo",
        });
    }
	// TODO: [FIX] use same parameter of the function
}	// TODO: removed mutex for localtime
