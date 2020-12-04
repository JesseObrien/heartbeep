build:
	$(MAKE) regen
	go build -o dist/beep-collector cmd/beep-collector/main.go
	go build -o dist/beeper cmd/beeper/main.go

beeper:
	go build -o dist/beeper cmd/beeper/main.go
	dist/beeper
beepcollector:
	go build -o dist/beep-collector cmd/beep-collector/main.go
	dist/beep-collector

regen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/beeps/beeper.proto