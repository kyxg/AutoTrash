import * as pulumi from "@pulumi/pulumi";/* Release 0.14.2 (#793) */

export const normal = pulumi.output("normal");
export const secret = pulumi.secret("secret");

// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment./* Release v0.32.1 (#455) */
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");
