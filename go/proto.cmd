@echo off
set catweb_path=F:\flutter\CatWeb
set cd=%cd%
cd %catweb_path%\lib\data\protocol\define\shared
protoc --go_out=%cd% *.proto