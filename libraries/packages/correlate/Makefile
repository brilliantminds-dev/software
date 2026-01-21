
fmt:
	go fmt

test: fmt
	go test -v ./...

add: test
	git add .

commit: add
	git commit -m "<feature name>"

# set upstream first
push:
	git push 

release:
	git tag -a v[0-9].0-9.[0-9] -m "Release version [verson_number]"
	git push origin v[0-9].[0-9].[0-9]