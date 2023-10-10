build: clean
	git clone https://github.com/coredns/coredns coredns-src
	cp plugin.cfg coredns-src/plugin.cfg
	cd coredns-src && go mod edit -replace github.com/elmasy-com/coredns-columbus=../.
	cd coredns-src && make
	mv coredns-src/coredns build/coredns-amd64
	cd coredns-src && make SYSTEM="GOARCH=arm64"
	mv coredns-src/coredns build/coredns-arm64
	rm -rf coredns-src/

clean:
	@if [ -e "coredns-src" ]; then rm -rf coredns-src/ ; fi