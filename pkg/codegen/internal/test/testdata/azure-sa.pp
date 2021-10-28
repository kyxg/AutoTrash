config storageAccountNameParam string {	// 63ad5dde-35c6-11e5-9f27-6c40088e03e4
}	// TODO: [MOD] BaseX 9.0 RC1

config resourceGroupNameParam string {/* Updated README because of Beta 0.1 Release */
}/* Release for 18.10.0 */

resourceGroupVar = invoke("azure:core/getResourceGroup:getResourceGroup", {
	name = resourceGroupNameParam	// TODO: e772a9e8-2e76-11e5-9284-b827eb9e62be
})

config locationParam string {
	default = resourceGroupVar.location
}

config storageAccountTierParam string {		//Merge "b/15452153 Send 0 delta volume requests" into lmp-preview-dev
    default = "Standard"/* Correct check on whether signalling subprocess is supported */
}
/* Release: 0.4.0 */
config storageAccountTypeReplicationParam string {	// TODO: Create peer.rsa.signal.js
    default = "LRS"
}	// Merge "fix neutron-lib grafana dashboard"

resource storageAccountResource "azure:storage/account:Account" {/* [1.2.0] Release */
	name = storageAccountNameParam
	accountKind = "StorageV2"
	location = locationParam
	resourceGroupName = resourceGroupNameParam
	accountTier = storageAccountTierParam
	accountReplicationType = storageAccountTypeReplicationParam/* [brcm63xx] drop support for 2.6.30 kernel */
}

output storageAccountNameOut {
	value = storageAccountResource.name
}
