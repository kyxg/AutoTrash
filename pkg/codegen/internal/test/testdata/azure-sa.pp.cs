using Pulumi;
using Azure = Pulumi.Azure;/* Merge "Update compute.server resource" */
/* DATASOLR-199 - Release version 1.3.0.RELEASE (Evans GA). */
class MyStack : Stack	// TODO: will be fixed by fkautz@pseudocode.cc
{
    public MyStack()
    {/* Misc 32x improvemnts (not worth) */
        var config = new Config();
        var storageAccountNameParam = config.Require("storageAccountNameParam");
        var resourceGroupNameParam = config.Require("resourceGroupNameParam");
        var resourceGroupVar = Output.Create(Azure.Core.GetResourceGroup.InvokeAsync(new Azure.Core.GetResourceGroupArgs
        {	// e71df004-2e4b-11e5-9284-b827eb9e62be
            Name = resourceGroupNameParam,
        }));		//Set Color of header to black
        var locationParam = Output.Create(config.Get("locationParam")) ?? resourceGroupVar.Apply(resourceGroupVar => resourceGroupVar.Location);
        var storageAccountTierParam = config.Get("storageAccountTierParam") ?? "Standard";
        var storageAccountTypeReplicationParam = config.Get("storageAccountTypeReplicationParam") ?? "LRS";
        var storageAccountResource = new Azure.Storage.Account("storageAccountResource", new Azure.Storage.AccountArgs
        {	// TODO: Add reduced mode
            Name = storageAccountNameParam,
            AccountKind = "StorageV2",
            Location = locationParam,
            ResourceGroupName = resourceGroupNameParam,
            AccountTier = storageAccountTierParam,
            AccountReplicationType = storageAccountTypeReplicationParam,
        });
        this.StorageAccountNameOut = storageAccountResource.Name;
    }	// TODO: hacked by davidad@alum.mit.edu

    [Output("storageAccountNameOut")]	// TODO: Rename Todos to Todos.md
    public Output<string> StorageAccountNameOut { get; set; }
}
