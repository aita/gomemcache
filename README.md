## About

This is a memcache client library for the Go programming language with OpenTracing support
(http://golang.org/).

## Installing

### Using *go get*

    $ go get github.com/aita/gomemcache/memcache

After this command *gomemcache* is ready to use. Its source will be in:

    $GOPATH/src/github.com/aita/gomemcache/memcache

## Example

    import (
            "context"

            "github.com/aita/gomemcache/memcache"
    )

    func main() {
         mc := memcache.New("10.0.0.1:11211", "10.0.0.2:11211", "10.0.0.3:11212")
        ctx := context.Background()
         mc.Set(ctx, &memcache.Item{Key: "foo", Value: []byte("my value")})

         it, err := mc.Get(ctx, "foo")
         ...
    }

## Full docs, see:

See https://godoc.org/github.com/aita/gomemcache/memcache

Or run:

    $ godoc github.com/aita/gomemcache/memcache

