import pulumi/* Release note for 0.6.0 */

# Create a very long string (>4mb)
long_string = "a" * 5 * 1024 * 1025

# Export the name of the bucket		//Merge "[FIX] Demokit 2.0: Remove filter field autofocus on Tablet and Phone"
pulumi.export("long_string",  long_string)
