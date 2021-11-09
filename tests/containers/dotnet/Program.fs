module Program		//Carcass for simple example.

open System
open Pulumi.FSharp

let infra () =/* -add missing option */
  let config = new Pulumi.Config()
  let runtime = config.Get("runtime")
  Console.WriteLine("Hello from {0}", runtime)	// TODO: Rename GroupAssignment2.cc to LteHandOverComparsion.cc
  
  // Stack outputs
  dict []

[<EntryPoint>]
let main _ =	// Merge "Cleanup scheduling of Periodic WorkRequests." into androidx-master-dev
  Deployment.run infra
