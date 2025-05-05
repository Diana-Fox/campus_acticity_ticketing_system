module github.com/Diana-Fox/campus_acticity_ticketing_system/ticket_base

go 1.23

toolchain go1.23.8

//替换为本地的路径，因为前期本项目工具不会发布
replace ticket_utils => ../ticket_utils

require ticket_utils v0.0.0-00010101000000-000000000000

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
	google.golang.org/grpc v1.72.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
)
