module github.com/deifyed/gin-upstream

go 1.16

replace github.com/deifyed/gatekeeper => ../../

require (
	github.com/coreos/go-oidc v2.2.1+incompatible
	github.com/deifyed/gatekeeper v0.0.0
	github.com/gin-gonic/gin v1.6.3
)
