all: buildgo

run: buildgo
	./run.sh

run-tests:
	./run-tests.sh

.PHONY: buildgo
buildgo: buildlibinjection
	go build main.go
	@echo "Run ./run.sh"

buildlibinjection:
	gcc -std=c99 -Wall -Werror -fpic -c lib/libinjection/libinjection_sqli.c -o lib/libinjection/libinjection_sqli.o
	gcc -std=c99 -Wall -Werror -fpic -c lib/libinjection/libinjection_xss.c -o lib/libinjection/libinjection_xss.o
	gcc -std=c99 -Wall -Werror -fpic -c lib/libinjection/libinjection_html5.c -o lib/libinjection/libinjection_html5.o
	gcc -dynamiclib -shared -o lib/libinjection/libinjection.so lib/libinjection/libinjection_sqli.o lib/libinjection/libinjection_xss.o lib/libinjection/libinjection_html5.o

clean:
	@rm -rf lib/libinjection/*.so
	@rm -rf lib/libinjection/*.o
	@rm -f main
