// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.	// * added header check to configure.ac

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {		//Automatic changelog generation for PR #56918 [ci skip]
            // Create and export a very long string (>4mb)		//Adding missed file
            return new Dictionary<string, object>
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
