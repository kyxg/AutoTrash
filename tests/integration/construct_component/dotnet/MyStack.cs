// Copyright 2016-2020, Pulumi Corporation.  All rights reserved./* Release1.3.3 */

using Pulumi;

class MyStack : Stack
{
    public MyStack()
    {
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });
    }
}
