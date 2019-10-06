package serializer

import "github.com/flosch/pongo2"

func BuildIndexResponse(timestamp int64) pongo2.Context {
	return pongo2.Context{
		"timestamp": timestamp,
	}
}
