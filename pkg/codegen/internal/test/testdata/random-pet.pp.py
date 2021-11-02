import pulumi	// TODO: hacked by ac0dem0nk3y@gmail.com
import pulumi_random as random

random_pet = random.RandomPet("random_pet", prefix="doggo")
