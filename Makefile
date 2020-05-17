GO?=go
LIB?=
build:
	$(GO) get ./src/... 
	$(GO) build -o bin/docksec src/*.go

install:
	install -m700 -o0 -g0 bin/docksec /usr/local/bin/docksec
	install -m700 -o0 -g0 service/docksec.service /lib/systemd/system/docksec.service
