GO?=go
LIB?=
build:
	$(GO) build -o bin/docksec src/*.go

install:
	mkdir -pv /etc/docksec
	mkdir -pv /usr/share/docksec
	install -m101 -o0 -g0 bin/docksec /usr/local/bin/docksec
	install -m700 -o0 -g0 service/docksec.service /lib/systemd/system/docksec.service
	install -m700 -o0 -g0 skel/main.json /etc/docksec/main.json
	install -m700 -o0 -g0 db/api.yml /usr/share/docksec/api.yml

clean:
	rm -rf bin/*

