build:
	protoc -I . --go_out=plugins=micro:. proto/consignment/consignment.proto
	docker build -t consignment-service .
run:
	docker run --net="host" \
	-p 50051 \
	-e MICRO_SERVER_ADDRESS=:50051 \
	-e MICRO_REGISTRY=mdns \
	consignment-service