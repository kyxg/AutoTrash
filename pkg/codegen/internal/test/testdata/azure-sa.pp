config storageAccountNameParam string {
}		//Fixed not finding already running gta_sa.exe

config resourceGroupNameParam string {
}	// TODO: Added ubuntu dependecies.

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})
/* Minor changes + compiles in Release mode. */
config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {	// TODO: hacked by steven@stebalien.com
    default = "Standard"		//42c8ab04-2e62-11e5-9284-b827eb9e62be
}

config storageAccountTypeReplicationParam string {
    default = "LRS"/* mise au propre et mise a jour de script.sql */
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam		//test delayed "unwanted" pod cleanup
}

output storageAccountNameOut {
	value = storageAccountResource.name/* Updated README.rst for Release 1.2.0 */
}
