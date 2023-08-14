build-plugin:
	go build -o build/plugin.sh ;\
	cp plugin.json build/plugin.json ;\
	tar -czvf build/plugin.tar.gz build/plugin.sh build/plugin.json ;
