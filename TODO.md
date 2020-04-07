# Group 'search'
```
- [ ] SEARCH
- [ ] (arg) key
- [ ] (arg) [CURSOR start]
- [ ] (arg) [LIMIT count]
- [ ] (arg) [MATCH pattern]
- [ ] (arg) [ASC|DESC]
- [ ] (arg) [WHERE field min max ...]
- [ ] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [ ] (arg) [NOFIELDS]
- [ ] (arg) [COUNT|IDS]

- [ ] SCAN
- [ ] (arg) key
- [ ] (arg) [CURSOR start]
- [ ] (arg) [LIMIT count]
- [ ] (arg) [MATCH pattern]
- [ ] (arg) [ASC|DESC]
- [ ] (arg) [WHERE field min max ...]
- [ ] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [ ] (arg) [NOFIELDS]
- [ ] (arg) [COUNT|IDS|OBJECTS|POINTS|BOUNDS|(HASHES precision)]

- [x] WITHIN
- [x] (arg) key
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [SPARSE spread]
- [x] (arg) [MATCH pattern]
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [NOFIELDS]
- [ ] (arg) [FENCE]
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [x] (arg) [COUNT|IDS|OBJECTS|POINTS|BOUNDS|(HASHES precision)]
- [ ] (arg) (GET key id)|(BOUNDS minlat minlon maxlat maxlon)|(OBJECT geojson)|(CIRCLE lat lon meters)|(TILE x y z)|(QUADKEY quadkey)|(HASH geohash)

- [x] INTERSECTS
- [x] (arg) key
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [SPARSE spread]
- [x] (arg) [MATCH pattern]
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [CLIP]
- [x] (arg) [NOFIELDS]
- [ ] (arg) [FENCE]
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [x] (arg) [COUNT|IDS|OBJECTS|POINTS|BOUNDS|(HASHES precision)]
- [ ] (arg) (GET key id)|(BOUNDS minlat minlon maxlat maxlon)|(OBJECT geojson)|(CIRCLE lat lon meters)|(TILE x y z)|(QUADKEY quadkey)|(HASH geohash)

- [x] NEARBY
- [x] (arg) key
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [SPARSE spread]
- [x] (arg) [MATCH pattern]
- [ ] (arg) [DISTANCE]
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [NOFIELDS]
- [ ] (arg) [FENCE]
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [x] (arg) [COUNT|IDS|OBJECTS|POINTS|BOUNDS|(HASHES precision)]
- [ ] (arg) (POINT lat lon meters)|(ROAM key pattern meters)


```
# Group 'pubsub'
```
- [ ] DELCHAN
- [ ] (arg) name

- [ ] CHANS
- [ ] (arg) pattern

- [ ] PSUBSCRIBE
- [ ] (arg) pattern [pattern ...]

- [ ] SETCHAN
- [ ] (arg) name
- [ ] (arg) [META name value ...]
- [ ] (arg) [EX seconds]
- [ ] (arg) NEARBY|WITHIN|INTERSECTS
- [ ] (arg) key
- [ ] (arg) FENCE
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [ ] (arg) param [param ...]

- [ ] PDELCHAN
- [ ] (arg) pattern

- [ ] SUBSCRIBE
- [ ] (arg) channel [channel ...]


```
# Group 'server'
```
- [ ] CONFIG REWRITE
- [ ] CONFIG SET
- [ ] (arg) parameter
- [ ] (arg) [value]

- [ ] GC
- [ ] FLUSHDB
- [ ] SERVER
- [ ] CONFIG GET
- [ ] (arg) parameter

- [ ] READONLY
- [ ] (arg) yes|no


```
# Group 'replication'
```
- [ ] FOLLOW
- [ ] (arg) host
- [ ] (arg) port

- [ ] AOFMD5
- [ ] (arg) pos
- [ ] (arg) size

- [ ] AOFSHRINK
- [ ] AOF
- [ ] (arg) pos


```
# Group 'tests'
```
- [ ] TEST
- [ ] (arg) (POINT lat lon)|(GET key id)|(BOUNDS minlat minlon maxlat maxlon)|(OBJECT geojson)|(CIRCLE lat lon meters)|(TILE x y z)|(QUADKEY quadkey)|(HASH geohash)
- [ ] (arg) INTERSECTS|WITHIN
- [ ] (arg) [CLIP]
- [ ] (arg) (POINT lat lon)|(GET key id)|(BOUNDS minlat minlon maxlat maxlon)|(OBJECT geojson)|(CIRCLE lat lon meters)|(TILE x y z)|(QUADKEY quadkey)|(HASH geohash)


```
# Group 'keys'
```
- [ ] JDEL
- [ ] (arg) key
- [ ] (arg) id
- [ ] (arg) path

- [ ] RENAMENX
- [ ] (arg) key
- [ ] (arg) newkey

- [x] BOUNDS
- [x] (arg) key

- [ ] RENAME
- [ ] (arg) key
- [ ] (arg) newkey

- [ ] EXPIRE
- [ ] (arg) key
- [ ] (arg) id
- [ ] (arg) seconds

- [ ] PDEL
- [ ] (arg) key
- [ ] (arg) pattern

- [ ] DEL
- [ ] (arg) key
- [ ] (arg) id

- [x] SET
- [x] (arg) key
- [x] (arg) id
- [x] (arg) [FIELD name value ...]
- [ ] (arg) [EX seconds]
- [x] (arg) [NX|XX]
- [ ] (arg) (OBJECT geojson)|(POINT lat lon [z])|(BOUNDS minlat minlon maxlat maxlon)|(HASH geohash)|(STRING value)

- [ ] PERSIST
- [ ] (arg) key
- [ ] (arg) id

- [ ] FSET
- [ ] (arg) key
- [ ] (arg) id
- [ ] (arg) [XX]
- [ ] (arg) field value
- [ ] (arg) [field value ...]

- [ ] JSET
- [ ] (arg) key
- [ ] (arg) id
- [ ] (arg) path
- [ ] (arg) value
- [ ] (arg) [RAW|STR]

- [x] KEYS
- [x] (arg) pattern

- [x] DROP
- [x] (arg) key

- [ ] STATS
- [ ] (arg) key [key ...]

- [ ] JGET
- [ ] (arg) key
- [ ] (arg) id
- [ ] (arg) path
- [ ] (arg) [RAW]

- [x] GET
- [x] (arg) key
- [x] (arg) id
- [x] (arg) [WITHFIELDS]
- [x] (arg) [OBJECT|POINT|BOUNDS|(HASH geohash)]

- [ ] TTL
- [ ] (arg) key
- [ ] (arg) id


```
# Group 'connection'
```
- [ ] TIMEOUT
- [ ] (arg) seconds
- [ ] (arg) COMMAND
- [ ] (arg) [arg  ...]

- [x] PING
- [ ] QUIT
- [ ] OUTPUT
- [ ] (arg) [json|resp]

- [ ] AUTH
- [ ] (arg) password


```
# Group 'scripting'
```
- [ ] EVALSHA
- [ ] (arg) sha1
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] SCRIPT EXISTS
- [ ] (arg) sha1 ...

- [ ] EVALRO
- [ ] (arg) script
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] EVALNA
- [ ] (arg) script
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] EVALROSHA
- [ ] (arg) script
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] SCRIPT LOAD
- [ ] (arg) script

- [ ] EVALNASHA
- [ ] (arg) sha1
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] EVAL
- [ ] (arg) script
- [ ] (arg) numkeys
- [ ] (arg) [key ...]
- [ ] (arg) [arg ...]

- [ ] SCRIPT FLUSH

```
# Group 'webhook'
```
- [ ] HOOKS
- [ ] (arg) pattern

- [ ] PDELHOOK
- [ ] (arg) pattern

- [ ] DELHOOK
- [ ] (arg) name

- [ ] SETHOOK
- [ ] (arg) name
- [ ] (arg) endpoint
- [ ] (arg) [META name value ...]
- [ ] (arg) [EX seconds]
- [ ] (arg) NEARBY|WITHIN|INTERSECTS
- [ ] (arg) key
- [ ] (arg) FENCE
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [ ] (arg) param [param ...]


```
