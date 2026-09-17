module github.com/yoshihio/go-memps

go 1.26.1

require (
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/valyala/fasthttp v1.72.0
	github.com/yoshihio/go-memps/pkg/publisher v0.0.0-00010101000000-000000000000
	github.com/yoshihio/go-memps/pkg/subscriber v0.0.0-00010101000000-000000000000
	go.uber.org/dig v1.19.0
	google.golang.org/grpc v1.82.0
)

require (
	github.com/andybalholm/brotli v1.2.2 // indirect
	github.com/klauspost/compress v1.19.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.39.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260706201446-f0a921348800 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace github.com/yoshihio/go-memps/pkg/publisher => /pkg/publisher

replace github.com/yoshihio/go-memps/pkg/subscriber => /pkg/subscriber
