@echo off
set ANDROID_NDK_HOME=G:\SDK\AndroidSDK\ndk\24.0.7956693

set GOARCH=arm
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi21-clang

go build -ldflags "-w -s" -buildmode=c-shared -o ../bin/armeabi-v7a/libgo.so ../main.go

set GOARCH=arm64
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang

go build -ldflags "-w -s" -buildmode=c-shared -o ../bin/armeabi-v8a/libgo.so ../main.go