// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./rubberTree";

// Export enums:
export * from "../../types/enums/tree/v1";

// Import resources to register:
import { RubberTree } from "./rubberTree";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "plant-provider:tree/v1:RubberTree":
                return new RubberTree(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("plant-provider", "tree/v1", _module)
