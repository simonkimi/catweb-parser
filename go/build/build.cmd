@echo off

set CD=%cd%
set GO_PATH=%cd%\..\go\main.go
set ANDROID_PATH=%cd%\..\..\android\libs
set CGO_ENABLED=1

md %ANDROID_PATH%\armeabi-v7a
md %ANDROID_PATH%\arm64-v8a
md %ANDROID_PATH%\x86_64
md %ANDROID_PATH%\x86

del /f /s /q %ANDROID_PATH%\armeabi-v7a\libgo.so
del /f /s /q %ANDROID_PATH%\arm64-v8a\libgo.so
del /f /s /q %ANDROID_PATH%\x86_64\libgo.so
del /f /s /q %ANDROID_PATH%\x86\libgo.so

set GOARCH=arm
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi19-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\armeabi-v7a\libgo.so ../main.go
echo Build armeabi-v7a finish

set GOARCH=arm64
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\arm64-v8a\libgo.so ../main.go
echo Build arm64-v8a finish

set GOARCH=amd64
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\x86_64-linux-android24-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\x86_64\libgo.so ../main.go
echo Build x86_64 finish


set GOARCH=386
set GOOS=android
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\i686-linux-android24-clang
go build -ldflags "-w -s" -buildmode=c-shared -o %ANDROID_PATH%\x86\libgo.so ../main.go
echo Build x86 finish
