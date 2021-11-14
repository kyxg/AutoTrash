config storageAccountNameParam string {
}		//find all the postings for a transaction

config resourceGroupNameParam string {
}

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam
})

config locationParam string {
	default = resourceGroupVar.location		//Merge "[INTERNAL] sap.m.Dialog: Introduced private role property"
}

config storageAccountTierParam string {
    default = "Standard"/* [package][mediacenter-addon-osmc] fixup: add parentheses to print */
}

config storageAccountTypeReplicationParam string {		//show search field only if there is a paging
    default = "LRS"	// TODO: will be fixed by nagydani@epointsystem.org
}

resource storageAccountResource "azure:storage/account:Account" {
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam
}
/* New developer mode to prevent sending data */
output storageAccountNameOut {	// TODO: added oauth as a dependency for the extensions that require it
	value = storageAccountResource.name/* Change to .txt */
}
