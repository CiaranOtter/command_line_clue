module message_service

go 1.22.6

replace clc_services => ../clc_services

replace game_database => ../game_database

require (
	clc_services v0.0.0-00010101000000-000000000000
	game_database v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.65.0
)

require (
	github.com/lib/pq v1.10.9 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
