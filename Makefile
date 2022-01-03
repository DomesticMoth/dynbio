
dependencies:
	go get

build: dependencies  dynbio.go
	go build -o build/dynbio dynbio.go

clean:
	go clean
	rm -rf build
	rm -rf package
	rm -f ./package.deb

createdeb: build
	mkdir -p package/DEBIAN
	cp control package/DEBIAN/control
	mkdir -p package/usr/bin
	cp build/dynbio package/usr/bin/dynbio
	mkdir -p package/lib/systemd/system/
	cp dynbio.service package/lib/systemd/system/dynbio.service
	dpkg-deb --build ./package

installdeb: clean createdeb
	apt install ./package.deb
	systemctl daemon-reload
	systemctl enable dynbio
	systemctl start dynbio
