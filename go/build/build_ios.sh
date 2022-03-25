export CGO_ENABLED=1
export CGO_CFLAGS="-fembed-bitcode"

export GOARCH=arm64
export GOOS=ios
export CC=$GOROOT/misc/ios/clangwrap.sh
go build -ldflags "-w -s" -buildmode=c-archive -o ../../ios/Classes/libgo.a ../main.go


#export GOARCH=amd64
#export GOOS=ios
#export CGO_ENABLED=1
#
#go build -ldflags "-w -s" -buildmode=c-archive -o ./build/ios_amd64/libgo.a ./main.go