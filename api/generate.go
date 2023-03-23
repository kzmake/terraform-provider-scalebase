//go:generate go run github.com/go-swagger/go-swagger/cmd/swagger@latest generate client --quiet -A api -m api/models -c api/client -t ../. -f https://api.scalebase.com/

package api
