package util

import "context"

const KeyRequester = "requester"

type Requester interface {
	GetUID() string
}

type requesterData struct {
	UID string `json:"uid"`
}

func NewRequester(uid string) *requesterData {
	return &requesterData{
		UID: uid,
	}
}

func (r *requesterData) GetUID() string {
	return r.UID
}

func GetRequester(ctx context.Context) Requester {
	if requester, ok := ctx.Value(KeyRequester).(Requester); ok {
		return requester
	}

	return nil
}

func ContextWithRequester(ctx context.Context, requester Requester) context.Context {
	return context.WithValue(ctx, KeyRequester, requester)
}
