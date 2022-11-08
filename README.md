# Protobuf types for `go-pg`

## Example

```go

// Connect to the PostgreSQL
db := pg.Connect(&pg.Options{
	User: "postgres",
})

// Register protobuf appender and scanner
protobuf_types.RegisterAppenderAndScanner()

ctx := context.Background()

if err := db.Ping(ctx); err != nil {
    panic(err)
}

// Example struct
type MySchema struct {
    tableName struct{} `pg:"your_table"`

    IsLocked *wrapperspb.BoolValue `protobuf:"bytes,1,opt,name=is_locked,json=isLocked,proto3" json:"is_locked,omitempty"`
    CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty" pg:",type:timestamptz"`
}

// Select query
var schema []*MySchema
err := db.Model(&schema).Select()
if err != nil {
	panic(err)
}

// Result
fmt.Printf("%v", schema)

// outputs: 
// [is_locked:{value:true} create_time:{seconds:1667842124 nanos:909462000}]

```

# Useful links and libs

- [Golang ORM with focus on PostgreSQL features and performance](https://github.com/go-pg/pg)
- [Build protobufs in Go, easily](https://github.com/containerd/protobuild)
- [Inject custom tags to protobuf golang struct](https://github.com/favadi/protoc-go-inject-tag)
- [gRPC to JSON proxy generator following the gRPC HTTP spec](https://github.com/grpc-ecosystem/grpc-gateway)
- [Golang gRPC Middlewares: interceptor chaining, auth, logging, retries and more.](https://github.com/grpc-ecosystem/go-grpc-middleware)
- [Generic protocol generator based on golang's text/template (grpc/protobuf)](https://github.com/moul/protoc-gen-gotemplate)