go:
	go build -mod vendor .
	bash -c 'FADECANDYCAL_DATE=`date  "+%B %d"` ./fadetree'

stop:
	ssh fadetree -- /etc/init.d/fadetree stop

start:
	ssh fadetree -- /etc/init.d/fadetree start

restart:
	ssh fadetree -- /etc/init.d/fadetree restart

deploy: fadetree.mips
#	scp fadetree.mips fadetree:/tmp/
	scp fadetree.mips archive:/var/www/
	make restart

watch: deploy
	ssh -t fadetree -- /tmp/fadetree.mips

fadetree.mips: *.go colors/colors.go
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -mod vendor -o fadetree.mips .

test:
	go test -v .
	cd colors && go test -v

fmt:
	go fmt ...

clean:
	rm fadetree.mips fadetree
