@echo off
rem proto_path为CatWeb项目的位置
set cd=%cd%
cd %proto_path%\lib\data\protocol\define
protoc --go_out=%cd% selector.proto
protoc --go_out=%cd% parser.proto
protoc --go_out=%cd% model.proto
protoc --go_out=%cd% rpc.proto