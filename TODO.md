# Group 'search'
```
- [x] SCAN
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [MATCH pattern]
- [x] (arg) ASC
- [x] (arg) DESC
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [NOFIELDS]
- [x] (arg) COUNT
- [x] (arg) IDS
- [x] (arg) OBJECTS
- [x] (arg) POINTS
- [x] (arg) BOUNDS
- [x] (arg) HASHES precision

- [x] SEARCH
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [MATCH pattern]
- [x] (arg) ASC
- [x] (arg) DESC
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [NOFIELDS]
- [x] (arg) COUNT
- [x] (arg) IDS

- [x] NEARBY
- [x] (arg) [CURSOR start]
- [x] (arg) [LIMIT count]
- [x] (arg) [SPARSE spread]
- [x] (arg) [MATCH pattern]
- [x] (arg) [DISTANCE]
- [x] (arg) [WHERE field min max ...]
- [x] (arg) [WHEREIN field count value [value ...] ...]
- [ ] (arg) [WHEREEVAL script numargs arg [arg ...] ...]
- [ ] (arg) [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
- [x] (arg) [NOFIELDS]
- [ ] (arg) [FENCE]
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]
- [x] (arg) COUNT
- [x] (arg) IDS
- [x] (arg) OBJECTS
- [x] (arg) POINTS
- [x] (arg) BOUNDS
- [x] (arg) HASHES precision
- [x] (arg) POINT lat lon meters
- [ ] (arg) ROAM key pattern meters

- [x] WITHIN
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
- [x] (arg) COUNT
- [x] (arg) IDS
- [x] (arg) OBJECTS
- [x] (arg) POINTS
- [x] (arg) BOUNDS
- [x] (arg) HASHES precision
- [x] (arg) GET key id
- [x] (arg) BOUNDS minlat minlon maxlat maxlon
- [x] (arg) OBJECT geojson
- [x] (arg) CIRCLE lat lon meters
- [x] (arg) TILE x y z
- [x] (arg) QUADKEY quadkey
- [x] (arg) HASH geohash

- [x] INTERSECTS
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
- [x] (arg) COUNT
- [x] (arg) IDS
- [x] (arg) OBJECTS
- [x] (arg) POINTS
- [x] (arg) BOUNDS
- [x] (arg) HASHES precision
- [x] (arg) GET key id
- [x] (arg) BOUNDS minlat minlon maxlat maxlon
- [x] (arg) OBJECT geojson
- [x] (arg) CIRCLE lat lon meters
- [x] (arg) TILE x y z
- [x] (arg) QUADKEY quadkey
- [x] (arg) HASH geohash


```

# Group 'keys'
```
- [ ] STATS
- [ ] JSET
- [ ] (arg) RAW
- [ ] (arg) STR

- [ ] PERSIST
- [ ] JDEL
- [x] SET
- [x] (arg) [FIELD name value ...]
- [x] (arg) [EX seconds]
- [x] (arg) NX
- [x] (arg) XX
- [x] (arg) OBJECT geojson
- [x] (arg) POINT lat lon [z]
- [x] (arg) BOUNDS minlat minlon maxlat maxlon
- [x] (arg) HASH geohash
- [x] (arg) STRING value

- [ ] TTL
- [ ] JGET
- [ ] (arg) [RAW]

- [x] DROP
- [ ] EXPIRE
- [x] DEL
- [ ] RENAME
- [ ] PDEL
- [x] GET
- [x] (arg) [WITHFIELDS]
- [x] (arg) OBJECT
- [x] (arg) POINT
- [x] (arg) BOUNDS
- [x] (arg) HASH geohash

- [ ] FSET
- [ ] (arg) [XX]

- [x] BOUNDS
- [ ] RENAMENX
- [x] KEYS

```

# Group 'connection'
```
- [x] PING
- [ ] OUTPUT
- [ ] AUTH
- [ ] TIMEOUT
- [ ] (arg) COMMAND

- [ ] QUIT

```


# Group 'webhook'
```
- [ ] DELHOOK
- [ ] PDELHOOK
- [ ] SETHOOK
- [ ] (arg) [META name value ...]
- [ ] (arg) [EX seconds]
- [ ] (arg) NEARBY|WITHIN|INTERSECTS
- [ ] (arg) FENCE
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]

- [ ] HOOKS

```

# Group 'pubsub'
```
- [ ] PSUBSCRIBE
- [ ] CHANS
- [ ] SETCHAN
- [ ] (arg) [META name value ...]
- [ ] (arg) [EX seconds]
- [ ] (arg) NEARBY|WITHIN|INTERSECTS
- [ ] (arg) FENCE
- [ ] (arg) [DETECT what]
- [ ] (arg) [COMMANDS which]

- [ ] PDELCHAN
- [ ] SUBSCRIBE
- [ ] DELCHAN

```


# Group 'server'
```
- [ ] CONFIG GET
- [ ] FLUSHDB
- [ ] SERVER
- [ ] READONLY
- [ ] CONFIG REWRITE
- [ ] CONFIG SET
- [ ] GC

```


# Group 'scripting'
```
- [ ] EVALNASHA
- [ ] EVAL
- [ ] EVALRO
- [ ] SCRIPT FLUSH
- [ ] EVALNA
- [ ] EVALROSHA
- [ ] SCRIPT EXISTS
- [ ] EVALSHA
- [ ] SCRIPT LOAD

```


# Group 'replication'
```
- [ ] AOFSHRINK
- [ ] AOF
- [ ] AOFMD5
- [ ] FOLLOW

```


# Group 'tests'
```
- [ ] TEST
- [ ] (arg) POINT lat lon
- [ ] (arg) GET key id
- [ ] (arg) BOUNDS minlat minlon maxlat maxlon
- [ ] (arg) OBJECT geojson
- [ ] (arg) CIRCLE lat lon meters
- [ ] (arg) TILE x y z
- [ ] (arg) QUADKEY quadkey
- [ ] (arg) HASH geohash
- [ ] (arg) INTERSECTS
- [ ] (arg) WITHIN
- [ ] (arg) [CLIP]
- [ ] (arg) POINT lat lon
- [ ] (arg) GET key id
- [ ] (arg) BOUNDS minlat minlon maxlat maxlon
- [ ] (arg) OBJECT geojson
- [ ] (arg) CIRCLE lat lon meters
- [ ] (arg) TILE x y z
- [ ] (arg) QUADKEY quadkey
- [ ] (arg) HASH geohash


```


