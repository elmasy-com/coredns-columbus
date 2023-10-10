build: clean
	git clone https://github.com/coredns/coredns coredns-src
	cp plugin.cfg coredns-src/plugin.cfg
	cd coredns-src && make
	mv coredns-src/coredns build/coredns
	rm -rf coredns-src/

build-dev: clean
	git clone https://github.com/coredns/coredns coredns-src
	cp plugin.cfg coredns-src/plugin.cfg
	cd coredns-src && go mod edit -replace github.com/elmasy-com/coredns-columbus=../.
	cd coredns-src && make
	mv coredns-src/coredns test/coredns
	rm -rf coredns-src

clean:
	@if [ -e "coredns-src" ]; then rm -r coredns-src/ ; fi