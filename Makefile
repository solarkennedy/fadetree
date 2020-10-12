go:
	packr
	go build .
	packr clean
	bash -c 'FADECANDYCAL_DATE=`date  "+%B %d"` ./fadetree'

stop:
	ssh root@fadetree -- /etc/init.d/fadetree stop

restart:
	ssh root@fadetree -- /etc/init.d/fadetree restart

deploy: fadetree.mips
	scp fadetree.mips root@fadetree:/tmp/

fadetree.mips: fadetree.go colors/colors.go
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -o fadetree.mips .

test:
	go test -v .
	cd colors && go test -v

fmt:
	go fmt ...

clean:
	rm fadetree.mips fadetree
