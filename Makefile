compile: 
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./core/user.proto 

run-test:
	cd server && go test && cd -

run-server:
	go run main.go

build-server:
	GOOS=linux go build -tags prod

build-client:
	cd client && GOOS=linux go build -tags prod -o clientBin && cd -
	mv client/clientBin clientBin

dock-build:
	sudo docker build --build-arg GIT_ACCESS_TOKEN=[e5c798426f2ec39c989886c1b3c70f606831b9cd] -t usermocker .
dock-run:
	sudo docker run -p 9090:9090 -t usermocker