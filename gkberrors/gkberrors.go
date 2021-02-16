package gkberrors

import "errors"

// ErrFailToPost : 어떤 객체를 포스트하지 못했을 때 발생하는 오류입니다.
// 주로 서버 수 포스팅에 문제가 생겼을 때 나옵니다.
var ErrFailToPost = errors.New("Failed to post something")