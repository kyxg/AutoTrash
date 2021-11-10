// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;		//Remove annotate_models plugin
using System.Threading.Tasks;
using Pulumi;

class Program
{		//enable unit tests
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";/* Update Week2Presentations */
            var a = new StackReference(slug);
/* [snomed] Move SnomedReleases helper class to snomed.core.domain package */
            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
