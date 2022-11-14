package t38c

import (
	"context"
	"strconv"
)

type KeysGetQueryBuilder struct {
	client        tile38Client
	key, objectID string
	withFields    bool
}

func newKeysGetQueryBuilder(client tile38Client, key, objectID string) KeysGetQueryBuilder {
	return KeysGetQueryBuilder{
		client:   client,
		key:      key,
		objectID: objectID,
	}
}

func (q KeysGetQueryBuilder) WithFields() KeysGetQueryBuilder {
	q.withFields = true
	return q
}

type GetObjectResponse struct {
	Object Object             `json:"object"`
	Fields map[string]float64 `json:"fields"`
}

func (q KeysGetQueryBuilder) Object(ctx context.Context) (*GetObjectResponse, error) {
	resp := new(GetObjectResponse)
	args := []string{q.key, q.objectID}
	if q.withFields {
		args = append(args, "WITHFIELDS")
	}

	err := q.client.jExecute(ctx, resp, "GET", args...)
	return resp, err
}

type GetPointResponse struct {
	Point  Point              `json:"point"`
	Fields map[string]float64 `json:"fields"`
}

func (q KeysGetQueryBuilder) Point(ctx context.Context) (*GetPointResponse, error) {
	resp := new(GetPointResponse)
	args := []string{q.key, q.objectID}
	if q.withFields {
		args = append(args, "WITHFIELDS")
	}

	args = append(args, "POINT")
	err := q.client.jExecute(ctx, resp, "GET", args...)
	return resp, err
}

type GetBoundsResponse struct {
	Bounds Bounds             `json:"bounds"`
	Fields map[string]float64 `json:"fields"`
}

func (q KeysGetQueryBuilder) Bounds(ctx context.Context) (*GetBoundsResponse, error) {
	resp := new(GetBoundsResponse)
	args := []string{q.key, q.objectID}
	if q.withFields {
		args = append(args, "WITHFIELDS")
	}

	args = append(args, "BOUNDS")
	err := q.client.jExecute(ctx, resp, "GET", args...)
	return resp, err
}

type GetHashResponse struct {
	Hash   string             `json:"hash"`
	Fields map[string]float64 `json:"fields"`
}

func (q KeysGetQueryBuilder) Hash(ctx context.Context, precision int) (*GetHashResponse, error) {
	resp := new(GetHashResponse)
	args := []string{q.key, q.objectID}
	if q.withFields {
		args = append(args, "WITHFIELDS")
	}

	args = append(args, "HASH", strconv.Itoa(precision))
	err := q.client.jExecute(ctx, resp, "GET", args...)
	return resp, err
}
