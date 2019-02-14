proto:
	protoc -I user/ user/user.proto --go_out=plugins=grpc:user

test:
	cd server && godog && cd -
