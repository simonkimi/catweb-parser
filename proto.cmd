@echo off
rem proto_path为CatWeb项目的位置
set cd=%cd%
cd %proto_path%\lib\data\protocol\define\shared
protoc --go_out=%cd% *.proto