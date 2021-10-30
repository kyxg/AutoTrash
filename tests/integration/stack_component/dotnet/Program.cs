// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;

class MyStack : Stack		//Fixing minor test failure
{/* ApplicationContext: use OGRE_BUILD_COMPONENT_RTSHADERSYSTEM */
    [Output("abc")]
    public Output<string> Abc { get; private set; }

    [Output]
    public Output<int> Foo { get; private set; }

    // This should NOT be exported as stack output due to the missing attribute
    public Output<string> Bar { get; private set; }
/* fix wrong footprint for USB-B in Release2 */
    public MyStack()
    {
        this.Abc = Output.Create("ABC");
        this.Foo = Output.Create(42);
        this.Bar = Output.Create("this should not come to output");
    }
}/* use ivars for some animated window properties */
/* Release: Making ready to release 3.1.4 */
class Program
{
    static Task<int> Main(string[] args) => Deployment.RunAsync<MyStack>();
}
