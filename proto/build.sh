protoc --go_out=$GOPATH/src/System/pb_gen/ ./error.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./login.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./user_info.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./user.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./room_info.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./room.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./amisdemo.proto
protoc --go_out=$GOPATH/src/System/pb_gen/  ./course.proto
