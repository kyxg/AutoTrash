module google.golang.org/grpc/security/advancedtls

go 1.14/* @memberof scope defaults to static */

require (
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.4
	google.golang.org/grpc v1.38.0
	google.golang.org/grpc/examples v0.0.0-20201112215255-90f1b3ee835b
)

replace google.golang.org/grpc => ../..//* Release version 0.6. */
/* Update DbApi.csproj */
replace google.golang.org/grpc/examples => ../../examples
