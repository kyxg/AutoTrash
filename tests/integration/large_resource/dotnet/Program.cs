// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using System;
using Pulumi;

class Program
{	// TODO: will be fixed by ac0dem0nk3y@gmail.com
    static Task<int> Main(string[] args)
    {/* #137 Upgraded Spring Boot to 1.3.1.Release  */
        return Deployment.RunAsync(() =>
        {
            // Create and export a very long string (>4mb)
            return new Dictionary<string, object>	// TODO: PerformanceTest for Root.sqrt() and Root.isSquare()
            {
                {  "LongString", new string('a', 5 * 1024 * 1024) }
            };
        });
    }
}
