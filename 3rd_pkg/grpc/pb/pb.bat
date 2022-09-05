@echo off

set PROJECT_ROOT=F:\workspace\src\yiriGo\grpc

set SRC_PATH=%PROJECT_ROOT%\pb
set TAR_PATH=%PROJECT_ROOT%\pb

protoc --proto_path=%SRC_PATH% --go_out=%TAR_PATH% --go_opt=paths=source_relative %SRC_PATH%\*.proto

pause