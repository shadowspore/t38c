# Group 'search'
- [x] SCAN
    - [x] [CURSOR start]
    - [x] [LIMIT count]
    - [x] [MATCH pattern]
    - [x] ASC
    - [x] DESC
    - [x] [WHERE field min max ...]
    - [x] [WHEREIN field count value [value ...] ...]
    - [x] [WHEREEVAL script numargs arg [arg ...] ...]
    - [x] [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
    - [x] [NOFIELDS]
    - [x] COUNT
    - [x] IDS
    - [x] OBJECTS
    - [x] POINTS
    - [x] BOUNDS
    - [x] HASHES precision
- [x] SEARCH
    - [x] [CURSOR start]
    - [x] [LIMIT count]
    - [x] [MATCH pattern]
    - [x] ASC
    - [x] DESC
    - [x] [WHERE field min max ...]
    - [x] [WHEREIN field count value [value ...] ...]
    - [x] [WHEREEVAL script numargs arg [arg ...] ...]
    - [x] [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
    - [x] [NOFIELDS]
    - [x] COUNT
    - [x] IDS
- [x] NEARBY
    - [x] [CURSOR start]
    - [x] [LIMIT count]
    - [x] [SPARSE spread]
    - [x] [MATCH pattern]
    - [x] [DISTANCE]
    - [x] [WHERE field min max ...]
    - [x] [WHEREIN field count value [value ...] ...]
    - [x] [WHEREEVAL script numargs arg [arg ...] ...]
    - [x] [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
    - [x] [NOFIELDS]
    - [x] [FENCE]
    - [x] [DETECT what]
    - [x] [COMMANDS which]
    - [x] COUNT
    - [x] IDS
    - [x] OBJECTS
    - [x] POINTS
    - [x] BOUNDS
    - [x] HASHES precision
    - [x] POINT lat lon meters
    - [x] ROAM key pattern meters
- [x] WITHIN
    - [x] [CURSOR start]
    - [x] [LIMIT count]
    - [x] [SPARSE spread]
    - [x] [MATCH pattern]
    - [x] [WHERE field min max ...]
    - [x] [WHEREIN field count value [value ...] ...]
    - [x] [WHEREEVAL script numargs arg [arg ...] ...]
    - [x] [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
    - [x] [NOFIELDS]
    - [x] [FENCE]
    - [x] [DETECT what]
    - [x] [COMMANDS which]
    - [x] COUNT
    - [x] IDS
    - [x] OBJECTS
    - [x] POINTS
    - [x] BOUNDS
    - [x] HASHES precision
    - [x] GET key id
    - [x] BOUNDS minlat minlon maxlat maxlon
    - [x] OBJECT geojson
    - [x] CIRCLE lat lon meters
    - [x] TILE x y z
    - [x] QUADKEY quadkey
    - [x] HASH geohash
- [x] INTERSECTS
    - [x] [CURSOR start]
    - [x] [LIMIT count]
    - [x] [SPARSE spread]
    - [x] [MATCH pattern]
    - [x] [WHERE field min max ...]
    - [x] [WHEREIN field count value [value ...] ...]
    - [x] [WHEREEVAL script numargs arg [arg ...] ...]
    - [x] [WHEREEVALSHA sha1 numargs arg [arg ...] ...]
    - [x] [CLIP]
    - [x] [NOFIELDS]
    - [x] [FENCE]
    - [x] [DETECT what]
    - [x] [COMMANDS which]
    - [x] COUNT
    - [x] IDS
    - [x] OBJECTS
    - [x] POINTS
    - [x] BOUNDS
    - [x] HASHES precision
    - [x] GET key id
    - [x] BOUNDS minlat minlon maxlat maxlon
    - [x] OBJECT geojson
    - [x] CIRCLE lat lon meters
    - [x] TILE x y z
    - [x] QUADKEY quadkey
    - [x] HASH geohash

# Group 'keys'
- [x] STATS
- [x] JSET
    - [X] RAW
    - [X] STR
- [x] PERSIST
- [x] JDEL
- [x] SET
    - [x] [FIELD name value ...]
    - [x] [EX seconds]
    - [x] NX
    - [x] XX
    - [x] OBJECT geojson
    - [x] POINT lat lon [z]
    - [x] BOUNDS minlat minlon maxlat maxlon
    - [x] HASH geohash
    - [x] STRING value
- [x] TTL
- [x] JGET
    - [ ] [RAW]
- [x] DROP
- [x] EXPIRE
- [x] DEL
- [x] RENAME
- [x] PDEL
- [x] GET
    - [x] [WITHFIELDS]
    - [x] OBJECT
    - [x] POINT
    - [x] BOUNDS
    - [x] HASH geohash
- [x] FSET
    - [x] [XX]
- [x] BOUNDS
- [x] RENAMENX
- [x] KEYS

# Group 'connection'
- [x] PING
- [ ] OUTPUT
- [x] AUTH
- [ ] TIMEOUT
    - [ ] COMMAND
- [ ] QUIT

# Group 'webhook'
- [x] DELHOOK
- [x] PDELHOOK
- [x] SETHOOK
    - [x] [META name value ...]
    - [x] [EX seconds]
    - [x] NEARBY|WITHIN|INTERSECTS
    - [x] FENCE
    - [x] [DETECT what]
    - [x] [COMMANDS which]
- [x] HOOKS

# Group 'pubsub'
- [x] PSUBSCRIBE
- [x] CHANS
- [x] SETCHAN
    - [x] [META name value ...]
    - [x] [EX seconds]
    - [x] NEARBY|WITHIN|INTERSECTS
    - [x] FENCE
    - [x] [DETECT what]
    - [x] [COMMANDS which]
- [x] PDELCHAN
- [x] SUBSCRIBE
- [x] DELCHAN

# Group 'server'
- [ ] CONFIG GET
- [x] FLUSHDB
- [ ] SERVER
- [ ] READONLY
- [ ] CONFIG REWRITE
- [ ] CONFIG SET
- [ ] GC

# Group 'scripting'
- [x] EVALNASHA
- [x] EVAL
- [x] EVALRO
- [x] SCRIPT FLUSH
- [x] EVALNA
- [x] EVALROSHA
- [x] SCRIPT EXISTS
- [x] EVALSHA
- [x] SCRIPT LOAD

# Group 'replication'
- [ ] AOFSHRINK
- [ ] AOF
- [ ] AOFMD5
- [ ] FOLLOW
