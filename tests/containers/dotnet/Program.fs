module Program

open System
open Pulumi.FSharp
/* 7c7ce164-2e56-11e5-9284-b827eb9e62be */
let infra () =
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)
  
  // Stack outputs
  dict []
	// TODO: fix xml ws : catalog
[<EntryPoint>]
let main _ =
  Deployment.run infra		//Renamed html file to index.html
