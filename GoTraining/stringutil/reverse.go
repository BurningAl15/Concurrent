package stringutil

func Reverse(s string) string {
	return reversetwo(s)
}

/*
go build
	go build reverse.go reversetwo.go
	won't produce an output file

go install
	will place the package inside the pkg directory of the workspace
*/
