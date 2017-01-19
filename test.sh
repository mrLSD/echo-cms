rm -rf *.coverprofile
i=0
for Package in $(go list ./... | grep -v "vendor")
do
	(( i = i + 1 ))
	go test -v -covermode=count -coverprofile=$i.coverprofile $Package
done
gocovmerge *.coverprofile > merged.coverprofile
go tool cover -html=merged.coverprofile
rm -rf *.coverprofile