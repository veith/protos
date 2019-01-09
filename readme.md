#Proto Collection

Some protos in this repo are copies from [https://github.com/googleapis](https://github.com/googleapis/googleapis/tree/master/google/type) to work with [gogoprotobuf](https://github.com/gogo/protobuf) and golang.

## Copied Google Common Types
- color.proto
- date.proto
- dayofweek.proto
- latlng.proto
- money.proto
- postal_address.proto 
- timeofday.proto

## about gogoprotobuf
gogoprotobuf is a fork of golang/protobuf with extra code generation features.

The code generation is used to achieve:

- fast marshalling and unmarshalling
- more canonical Go structures
- goprotobuf compatibility
- less typing by optionally generating extra helper code
- peace of mind by optionally generating test and benchmark code
- other serialization formats

