package domain

import "errors"

var ErrNotFound = errors.New("not found")
var ErrLogout = errors.New("the user is logged out")
var ErrDeviceIncompatible = errors.New("the device uuid is invalid")
