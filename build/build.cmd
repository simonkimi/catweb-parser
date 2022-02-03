@echo off
set ANDROID_NDK_HOME=G:\SDK\AndroidSDK\ndk\24.0.7956693
set CATWEB_HOME=H:\Android\Project\catweb
set LIB_TARGET=%CATWEB_HOME%\android\app\libs
set CD=%cd%

go build -buildmode=c-shared -o ..\bin\windows\libgo.dll ..\main.go
echo Build native finish

set GOARCH=arm
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\armv7a-linux-androideabi21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o ..\bin\armeabi-v7a\libgo.so ..\main.go
echo Build armeabi-v7a finish

set GOARCH=arm64
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\aarch64-linux-android21-clang
go build -ldflags "-w -s" -buildmode=c-shared -o ..\bin\arm64-v8a\libgo.so ..\main.go
echo Build arm64-v8a finish

set GOARCH=amd64
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\x86_64-linux-android24-clang
go build -ldflags "-w -s" -buildmode=c-shared -o ..\bin\x86_64\libgo.so ..\main.go
echo Build x86_64 finish


set GOARCH=386
set GOOS=android
set CGO_ENABLED=1
set CC=%ANDROID_NDK_HOME%\toolchains\llvm\prebuilt\windows-x86_64\bin\i686-linux-android24-clang

go build -ldflags "-w -s" -buildmode=c-shared -o ..\bin\x86\libgo.so ..\main.go

echo Build x86 finish


md %LIB_TARGET%\armeabi-v7a
md %LIB_TARGET%\arm64-v8a
md %LIB_TARGET%\x86_64
md %LIB_TARGET%\x86
copy /y %CD%\..\bin\windows\libgo.dll %CATWEB_HOME%\lib\libs\libgo.dll
copy /y %CD%\..\bin\armeabi-v7a\libgo.so %LIB_TARGET%\armeabi-v7a\libgo.so
copy /y %CD%\..\bin\arm64-v8a\libgo.so %LIB_TARGET%\arm64-v8a\libgo.so
copy /y %CD%\..\bin\x86_64\libgo.so %LIB_TARGET%\x86_64\libgo.so
copy /y %CD%\..\bin\x86\libgo.so %LIB_TARGET%\x86\libgo.so
echo Copy files finish