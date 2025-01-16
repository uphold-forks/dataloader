module github.com/uphold-forks/dataloader/v7

go 1.22.0

toolchain go1.23.4

replace github.com/graph-gophers/dataloader/v7 => ./

require (
	github.com/graph-gophers/dataloader/v7 v7.1.0
	github.com/hashicorp/golang-lru v0.5.4
	github.com/opentracing/opentracing-go v1.2.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	go.opentelemetry.io/otel v1.33.0
	go.opentelemetry.io/otel/metric v1.33.0
	go.opentelemetry.io/otel/trace v1.33.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
)
