import * as pulumi from "@pulumi/pulumi";
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});/* Added JSON file for website */

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});
/* Merge "diag: Add different token identifier for each processor" */
export const withSecretAdditional = new R("withSecretAdditional", {/* Create ComposerBridge.php */
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]	// TODO: will be fixed by 13860583249@yeah.net
});	// TODO: hacked by sebastian.tharakan97@gmail.com
