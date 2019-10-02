build:
	protoc -I. --go_out=plugins=micro:. proto/vessel/vessel.proto
	docker build -t vessel-service .

run:
	docker run --net="host" \
	-p 50052 \
	-e MICRO_SERVER_ADDRESS=:50052 \
	-e MICRO_REGISTRY=mdns \
	vessel-service