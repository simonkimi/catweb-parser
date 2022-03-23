export GOARCH=arm64
export GOOS=ios
export CGO_ENABLED=1

cd .. && go build -ldflags "-w -s" -buildmode=c-archive -o ./build/ios/libgo.a main.go


export GOARCH=amd64
export GOOS=ios
export CGO_ENABLED=1

cd .. && go build -ldflags "-w -s" -buildmode=c-archive -o ./build/ios_amd64/libgo.a main.go