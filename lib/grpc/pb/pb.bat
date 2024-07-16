@echo off

@REM set PROJECT_ROOT=F:\workspace\src\yiriGo\grpc
set PROJECT_ROOT= .

set SRC_PATH=%PROJECT_ROOT%
set TAR_PATH=%PROJECT_ROOT%

@REM protoc --proto_path=%SRC_PATH% --go_out=%TAR_PATH% --go_opt=paths=source_relative %SRC_PATH%/*.proto
@REM protoc --proto_path=%SRC_PATH% --go_out=%TAR_PATH% --go_opt=paths=source_relative %SRC_PATH%/hello.proto

protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./hello.proto

pause