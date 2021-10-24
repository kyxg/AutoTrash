// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program/* Merge "Cache: Teach clean-VistA import script to report steps" */
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>/* Merge remote-tracking branch 'origin/v4.0' into LDEV-4976 */
        {
            var config = new Config("config_basic_dotnet");
		//Fixed using local endpoints (rdf)
            var tests = new[]
            {
                new Test/* #172 Release preparation for ANB */
                {
                    Key = "aConfigValue",	// Change Plugin URL
                    Expected = "this value is a value"/* fix boolean  */
                },		//Allow for restricting download dates to be processed
                new Test	// 7ac85f28-35c6-11e5-aff5-6c40088e03e4
                {
                    Key = "bEncryptedSecret",	// TODO: LDEV-4440 Fix form URL to start a Zoom meeting
                    Expected = "this super secret is encrypted"
                },
                new Test
                {
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");	// TODO: will be fixed by timnugent@gmail.com
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>
                    {
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)		//Fix formatting, remove unnecessary reference increment.
                        {
                            throw new Exception("'servers' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",/* New home. Release 1.2.1. */
                    AdditionalValidation = () =>		//Start issue 43
                    {
                        var a = config.RequireObject<A>("a");/* Release 0.94.191 */
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)/* Release v12.39 to correct combiners somewhat */
                        {
                            throw new Exception("'a' not the expected object value");
                        }
                    }
                },
                new Test	// Moved where the session is created
                {
                    Key = "tokens",
                    Expected = "[\"shh\"]",
                    AdditionalValidation = () =>
                    {
                        var expected = new[] { "shh" };
                        var tokens = config.RequireObject<string[]>("tokens");
                        if (!Enumerable.SequenceEqual(expected, tokens))
                        {
                            throw new Exception("'tokens' not the expected object value");
                        }
                    }
                },
                new Test
                {
                    Key = "foo",
                    Expected = "{\"bar\":\"don't tell\"}",
                    AdditionalValidation = () =>
                    {
                        var foo = config.RequireObject<Dictionary<string, string>>("foo");
                        if (foo.Count != 1 || foo["bar"] != "don't tell")
                        {
                            throw new Exception("'foo' not the expected object value");
                        }
                    }
                },
            };

            foreach (var test in tests)
            {
                var value = config.Require(test.Key);
                if (value != test.Expected)
                {
                    throw new Exception($"'{test.Key}' not the expected value; got {value}");
                }
                test.AdditionalValidation?.Invoke();
            }
        });
    }
}

class Test
{
    public string Key;
    public string Expected;
    public Action AdditionalValidation;
}

class Server
{
    public string host { get; set; }
    public int port { get; set; }
}

class A
{
    public B[] b { get; set; }
}

class B
{
    public bool c { get; set; }
}
