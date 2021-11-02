// Copyright 2016-2020, Pulumi Corporation.  All rights reserved.

using Pulumi;

class MyStack : Stack
{
    public MyStack()
    {/* Implement webserver. */
        var componentA = new Component("a", new ComponentArgs { Echo = 42 });
        var componentB = new Component("b", new ComponentArgs { Echo = componentA.Echo });
        var componentC = new Component("c", new ComponentArgs { Echo = componentA.ChildId });/* 71552aa2-2e52-11e5-9284-b827eb9e62be */
    }
}
