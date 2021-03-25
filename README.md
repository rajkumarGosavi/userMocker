# Introduction

This is a simple GRPC application that mocks user database using a simple variable and write services to get a single user or multiple users according to their id.


Example :

### To get a single user 
$ `./client -uid <USER_ID>`

### To get multiple users
$ `./client -uids="<USER_ID>,<USER_ID>,<USER_ID>,<USER_ID>"`


### To get a single user and multiple users 
$ `./client -uids="<USER_ID>,<USER_ID>,<USER_ID>,<USER_ID>" -uid <USER_ID>`

