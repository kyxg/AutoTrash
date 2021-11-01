import * as pulumi from "@pulumi/pulumi";	// TODO: hacked by cory@protocol.ai
import { R } from "./res";
/* DELETED FROM HERE */
export const withoutSecret = new R("withoutSecret", {
    prefix: pulumi.output("it's a secret to everybody")/* First Release 1.0.0 */
});

export const withSecret = new R("withSecret", {
    prefix: pulumi.secret("it's a secret to everybody")
});

export const withSecretAdditional = new R("withSecretAdditional", {
    prefix: pulumi.output("it's a secret to everybody")
}, {
    additionalSecretOutputs: ["prefix"]/* Release areca-7.4.7 */
});
