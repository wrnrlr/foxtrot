module github.com/wrnrlr/foxtrot

go 1.13

require (
	gioui.org v0.0.0-20191124190138-726270ab2d84
	github.com/RoaringBitmap/roaring v0.4.21 // indirect
	github.com/blevesearch/bleve v0.8.1
	github.com/blevesearch/blevex v0.0.0-20190916190636-152f0fe5c040 // indirect
	github.com/blevesearch/go-porterstemmer v1.0.2 // indirect
	github.com/blevesearch/segment v0.0.0-20160915185041-762005e7a34f // indirect
	github.com/corywalker/expreduce v0.0.0-20190902204200-0a346d0d4ef1
	github.com/couchbase/vellum v0.0.0-20190829182332-ef2e028c01fd // indirect
	github.com/cznic/b v0.0.0-20181122101859-a26611c4d92d // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/etcd-io/bbolt v1.3.3 // indirect
	github.com/facebookgo/ensure v0.0.0-20160127193407-b4ab57deab51 // indirect
	github.com/facebookgo/stack v0.0.0-20160209184415-751773369052 // indirect
	github.com/facebookgo/subset v0.0.0-20150612182917-8dac2c3c4870 // indirect
	github.com/golang/protobuf v1.3.2 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/steveyen/gtreap v0.0.0-20150807155958-0abe01ef9be2 // indirect
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/tecbot/gorocksdb v0.0.0-20191122205208-eb0a0d0d32b3 // indirect
	go.etcd.io/bbolt v1.3.3 // indirect
	golang.org/x/exp v0.0.0-20191030013958-a1ab85dbe136 // indirect
	golang.org/x/image v0.0.0-20191009234506-e7c1f5e7dbb8
	golang.org/x/sys v0.0.0-20191110163157-d32e6e3b99c4 // indirect
	golang.org/x/text v0.3.2 // indirect
	gopkg.in/yaml.v2 v2.2.5 // indirect
	modernc.org/wl v1.0.0
)

replace (
	gioui.org => ./gio
	github.com/corywalker/expreduce => ./expreduce
)
