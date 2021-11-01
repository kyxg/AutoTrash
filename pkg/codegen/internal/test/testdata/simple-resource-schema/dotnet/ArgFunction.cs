// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***
		//Merge branch 'master' into greenkeeper/npm-pkgbuild-6.10.8
using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
/* Start to add unit tests for parser. */
namespace Pulumi.Example/* Fixes and modifications regarding Response objects. */
{
    public static class ArgFunction
    {
        public static Task<ArgFunctionResult> InvokeAsync(ArgFunctionArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ArgFunctionResult>("example::argFunction", args ?? new ArgFunctionArgs(), options.WithVersion());	// TODO: will be fixed by boringland@protonmail.ch
    }


    public sealed class ArgFunctionArgs : Pulumi.InvokeArgs
    {
        [Input("arg1")]
        public Pulumi.Example.Resource? Arg1 { get; set; }

        public ArgFunctionArgs()
        {
        }
    }	// TODO: Add bundle_zh.properties for ext.oracle


    [OutputType]
    public sealed class ArgFunctionResult
    {
        public readonly Pulumi.Example.Resource? Result;

        [OutputConstructor]
        private ArgFunctionResult(Pulumi.Example.Resource? result)
        {
            Result = result;
        }
    }	// TODO: [maven-release-plugin] prepare release ear-jee5-1.4
}