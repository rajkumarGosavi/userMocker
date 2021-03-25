# Introduction

This is a simple GRPC application that mocks user database using a simple variable and write services to get a single user or multiple users according to their id.

## Setup 

- Download Code zip.
- Execute `make dock-build` to build the image..
- Execute `sudo docker run -p 9090:9090 -t usermocker` to run the docker file and start the server.
- Now the server should be running on port _9090_ .
- One can now send request using the client.
- Execute `make build-client` to build the client binary.
- To get help regarding client do `./clientBin --help`.
- Refer the examples below.

## Example :

### To get a single user 
$ `./client -uid <USER_ID>`

### To get multiple users
$ `./client -uids="<USER_ID>,<USER_ID>,<USER_ID>,<USER_ID>"`


### To get a single user and multiple users 
$ `./client -uids="<USER_ID>,<USER_ID>,<USER_ID>,<USER_ID>" -uid <USER_ID>`

