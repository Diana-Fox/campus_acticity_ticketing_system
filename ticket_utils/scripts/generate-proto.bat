protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\user\user.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\base\area.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\base\image.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\base\item_type.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\base\school.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\base\scope.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\item\item.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\item\item_comment.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\item\place.proto

protoc --go_out=./pkg --go_opt=paths=source_relative ^
        --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
        .\api\item\place_seat.proto

@REM protoc --go_out=./pkg --go_opt=paths=source_relative ^
@REM         --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
@REM         .\api\order\order.proto
@REM
@REM protoc --go_out=./pkg --go_opt=paths=source_relative ^
@REM         --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
@REM         .\api\order\order_link_user.proto
@REM
@REM protoc --go_out=./pkg --go_opt=paths=source_relative ^
@REM         --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
@REM         .\api\scheduler\scheduler.proto
@REM
@REM protoc --go_out=./pkg --go_opt=paths=source_relative ^
@REM         --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
@REM         .\api\scheduler\scheduler_seat.proto
@REM
@REM
@REM
@REM protoc --go_out=./pkg --go_opt=paths=source_relative ^
@REM         --go-grpc_out=./pkg --go-grpc_opt=paths=source_relative ^
@REM         .\api\user\campus_authent.proto

go mod tidy

