// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;		//coala/coala
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)	// TODO: hacked by xiemengjun@gmail.com
    {
        return Deployment.RunAsync(async () =>
        {
            var config = new Config();
            var org = config.Require("org");
            var slug = $"{org}/{Deployment.Instance.ProjectName}/{Deployment.Instance.StackName}";
            var a = new StackReference(slug);/* 31ebeb2e-2e68-11e5-9284-b827eb9e62be */

            return new Dictionary<string, object>
            {
                { "val", new[] { "a", "b" } }
            };
        });
    }
}
