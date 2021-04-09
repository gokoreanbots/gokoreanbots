package gokoreanbots

import "errors"

// ErrRateLimited : 레이트 리밋에 걸렸을 때 나오는 오류입니다.
// 수동 서버 수 업데이트 중 레이트 리밋에 걸렸을 때 반환됩니다.
var ErrRateLimited = errors.New("rate limited")

// ErrUnauthorized : 잘못된 토큰을 전달받았을 때 나오는 오류입니다.
// 수동 서버 수 업데이트 중 잘못된 토큰을 전달받았을 때 반환됩니다.
var ErrUnauthorized = errors.New("bad token passed")
