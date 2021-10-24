import * as pulumi from "@pulumi/pulumi";

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");
		//Added null subtitles event to clear subtitles when media changed
// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");		//Merge "[FAB-7766] Document on CouchDB (fix links)"
const project = pulumi.getProject();/* Include LICENSE in gem. */
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
