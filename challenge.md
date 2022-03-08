Location History Server
=======================

## Task Description

Your task is to implement a toy in-memory location history server.

Clients should be able to speak JSON over HTTP to the server. The three endpoints it should support are:
* `POST /location/{order_id}/now`
* `GET /location/{order_id}?max=<N>`
* `DELETE /location/{order_id}`

Details about the endpoints:

`POST /location/{order_id}/now` - append a location to the history for the specified order.
Example interaction:
```
POST /location/def456/now
{
	"lat": 12.34,
	"lng": 56.78
}

200 OK
```
`GET /location/{order_id}?max=<N>` - Retrieve at most N items of history for the specified order. The most recent locations (in chronological order of insertion) should be returned first, if history is truncated by the `max` parameter.
Example interaction:
```
GET /location/abc123?max=2

200 OK
{
	"order_id": "abc123",
	"history": [
		{"lat": 12.34, "lng": 56.78},
		{"lat": 12.34, "lng": 56.79}
	]
}
```
The `max` query parameter may or may not be present. If it is not present, the entire history should be returned.

`DELETE /location/{order_id}` - delete history for the specified order. Example interaction:
```
DELETE /location/xyz987

200 OK
```

## Bonus points

Make location history items expire after a configurable number of seconds, e.g. `LOCATION_HISTORY_TTL_SECONDS=60`. Document and discuss your implementation strategy.

## Submission guidelines and notes

Create a Go module which implements the server described above. Running `go build` in the root directory of this module should produce a working binary. The server should serve HTTP at the address specified in the environment variable `HISTORY_SERVER_LISTEN_ADDR`. If this environment variable is not set, the listen address should default to `:8080`.

If the task statement is unclear about something, please feel free to reach out. That being said, decisions about how to treat certain edge cases that may come up are at your discretion.

Good luck, and enjoy! :)