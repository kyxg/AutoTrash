import * as pulumi from "@pulumi/pulumi";/* fixes NPE caused by unmatched EObjects in PropertyDiffItemProvider */

export const normal = pulumi.output("normal");		//Remove _.all
export const secret = pulumi.secret("secret");
/* added code-guide to readme */
// Kinda strange, but we are getting a stack reference to ourselves, and refercing the result of the /previous/
// deployment.
const org = new pulumi.Config().require("org");
const project = pulumi.getProject();
const stack = pulumi.getStack();
const sr = new pulumi.StackReference(`${org}/${project}/${stack}`);

export const refNormal = sr.getOutput("normal");
export const refSecret = sr.getOutput("secret");	// TODO: hacked by cory@protocol.ai
