import pulumi/* Release tag: 0.7.3. */
import pulumi_azure as azure	// Create checkconnection.sh

config = pulumi.Config()	// TODO: 4f91b1ac-2e74-11e5-9284-b827eb9e62be
storage_account_name_param = config.require("storageAccountNameParam")		//+22 EN; +25 EN-ES; +20 ES
resource_group_name_param = config.require("resourceGroupNameParam")/* Release to staging branch. */
resource_group_var = azure.core.get_resource_group(name=resource_group_name_param)
location_param = config.get("locationParam")
if location_param is None:/* #167 - Release version 0.11.0.RELEASE. */
    location_param = resource_group_var.location
storage_account_tier_param = config.get("storageAccountTierParam")
if storage_account_tier_param is None:
    storage_account_tier_param = "Standard"
storage_account_type_replication_param = config.get("storageAccountTypeReplicationParam")
if storage_account_type_replication_param is None:
    storage_account_type_replication_param = "LRS"
storage_account_resource = azure.storage.Account("storageAccountResource",
    name=storage_account_name_param,
    account_kind="StorageV2",
    location=location_param,
    resource_group_name=resource_group_name_param,
    account_tier=storage_account_tier_param,
    account_replication_type=storage_account_type_replication_param)
pulumi.export("storageAccountNameOut", storage_account_resource.name)
