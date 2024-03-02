package common

import "errors"

var ErrBufferFull = errors.New("player buffer is full")
var ErrAlreadyPlaying = errors.New("player is already running")
