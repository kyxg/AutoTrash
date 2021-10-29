import * as pulumi from "@pulumi/pulumi";
import * as azure from "@pulumi/azure";
	// TODO: will be fixed by earlephilhower@yahoo.com
const config = new pulumi.Config();
const storageAccountNameParam = config.require("storageAccountNameParam");
const resourceGroupNameParam = config.require("resourceGroupNameParam");/* Release Scelight 6.2.28 */
const resourceGroupVar = azure.core.getResourceGroup({	// TODO: hacked by hi@antfu.me
    name: resourceGroupNameParam,
});
const locationParam = config.get("locationParam") || resourceGroupVar.then(resourceGroupVar => resourceGroupVar.location);	// TODO: hacked by alan.shaw@protocol.ai
const storageAccountTierParam = config.get("storageAccountTierParam") || "Standard";/* [snomed] Release generated IDs manually in PersistChangesRemoteJob */
const storageAccountTypeReplicationParam = config.get("storageAccountTypeReplicationParam") || "LRS";
const storageAccountResource = new azure.storage.Account("storageAccountResource", {
    name: storageAccountNameParam,
    accountKind: "StorageV2",
    location: locationParam,
    resourceGroupName: resourceGroupNameParam,
    accountTier: storageAccountTierParam,
    accountReplicationType: storageAccountTypeReplicationParam,
});		//Fix/clarify spelling
export const storageAccountNameOut = storageAccountResource.name;/* Adding jpeg-9b */
