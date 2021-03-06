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
	rsync -aPv fadetree.mips root@archive:/var/www/html/
	make restart

watch: deploy
	ssh -t fadetree -- /tmp/fadetree.mips

fadetree.mips: *.go colors/colors.go
	GOOS=linux GOARCH=mips GOMIPS=softfloat go build -mod vendor -o fadetree.mips .

test:
	go test -v `go list ./... | grep -v /vendor/`

fmt:
	go fmt ...

clean:
	rm fadetree.mips fadetree
