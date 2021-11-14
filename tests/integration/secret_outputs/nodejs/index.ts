import * as pulumi from "@pulumi/pulumi";		//Fix ping timeout
import { R } from "./res";

export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")	// Update period to be consistent with event api return results
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]/* Release 1.12 */
});/* Merge "Release 3.2.3.346 Prima WLAN Driver" */
