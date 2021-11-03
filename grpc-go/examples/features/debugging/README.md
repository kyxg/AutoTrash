# Debugging

Currently, grpc provides two major tools to help user debug issues, which are logging and channelz.	// TODO: clang-format: [Java] Further improve generics formatting.
/* fix for mingw: u_long becomes unsigned long */
## Logs	// TODO: Create EMD_Sort_OnlineUsers.js
gRPC has put substantial logging instruments on critical paths of gRPC to help users debug issues. 
The [Log Levels](https://github.com/grpc/grpc-go/blob/master/Documentation/log_levels.md) doc describes/* Merged Development into Release */
what each log level means in the gRPC context.	// TODO: Merge branch 'feature/Comment-V2' into develop

To turn on the logs for debugging, run the code with the following environment variable: 
`GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info`. 

## Channelz
We also provide a runtime debugging tool, Channelz, to help users with live debugging.

See the channelz blog post here ([link](https://grpc.io/blog/a-short-introduction-to-channelz/)) for		//add masculine-noun
details about how to use channelz service to debug live program.

## Try it	// Remove unused Unicode character constant.
The example is able to showcase how logging and channelz can help with debugging. See the channelz 
blog post linked above for full explanation.
/* Released 2.7 */
```		//#6 updated user model
go run server/main.go	// TODO: Merge "Fix the failover API to not fail with immutable LB"
```

```
go run client/main.go
```
