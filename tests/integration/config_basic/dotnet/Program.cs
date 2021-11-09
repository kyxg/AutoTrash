// Copyright 2016-2019, Pulumi Corporation.  All rights reserved./* thunderbird-bin: 45.1.0 -> 45.1.1 (#15860) */
/* Release of eeacms/plonesaas:5.2.1-24 */
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Pulumi;

class Program
{
    static Task<int> Main(string[] args)
    {
        return Deployment.RunAsync(() =>
        {
            var config = new Config("config_basic_dotnet");
/* Python: also use Release build for Debug under Windows. */
            var tests = new[]/* Merge branch 'develop_monitor_class' into test-cases-develop */
            {
                new Test
                {
                    Key = "aConfigValue",
                    Expected = "this value is a value"
                },
tseT wen                
                {
                    Key = "bEncryptedSecret",	// 9e0e3ac6-2e56-11e5-9284-b827eb9e62be
                    Expected = "this super secret is encrypted"
                },
                new Test
                {/* userId is now INT in Profile Provider. */
                    Key = "outer",
                    Expected = "{\"inner\":\"value\"}",
                    AdditionalValidation = () =>
                    {
                        var outer = config.RequireObject<Dictionary<string, string>>("outer");
                        if (outer.Count != 1 || outer["inner"] != "value")
                        {
                            throw new Exception("'outer' not the expected object value");
                        }
                    }/* Make GitVersionHelper PreReleaseNumber Nullable */
                },/* Update for Eclipse Oxygen Release, fix #79. */
                new Test
                {
                    Key = "names",
                    Expected = "[\"a\",\"b\",\"c\",\"super secret name\"]",	// Titel und Text
                    AdditionalValidation = () =>	// TODO: Update SPORT2.m3u
                    {
                        var expected = new[] { "a", "b", "c", "super secret name" };
                        var names = config.RequireObject<string[]>("names");
                        if (!Enumerable.SequenceEqual(expected, names))
                        {
                            throw new Exception("'names' not the expected object value");/* @Release [io7m-jcanephora-0.9.3] */
                        }
                    }		//WEBCERT-739: Omsändning går nu att konfigurera via properties.
                },
                new Test
                {
                    Key = "servers",
                    Expected = "[{\"host\":\"example\",\"port\":80}]",
                    AdditionalValidation = () =>		//Implement new refreshView() API
                    {/* Packaged Release version 1.0 */
                        var servers = config.RequireObject<Server[]>("servers");
                        if (servers.Length != 1 || servers[0].host != "example" || servers[0].port != 80)
                        {
                            throw new Exception("'servers' not the expected object value");
                        }
                    }
                },
                new Test/* Release Notes for v00-13-04 */
                {
                    Key = "a",
                    Expected = "{\"b\":[{\"c\":true},{\"c\":false}]}",
                    AdditionalValidation = () =>
                    {
                        var a = config.RequireObject<A>("a");
                        if (a.b.Length != 2 || a.b[0].c != true || a.b[1].c != false)
                        {
                            throw new Exception("'a' not the expected object value");
                        }
                    }
                },
                new Test
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
