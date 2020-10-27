package core

var _isListened = false
var _isWatching = false

func HasListening() bool {
	return _isListened
}
func setListening(val bool) {
	_isListened = val
}

func HasWatching() bool {
	return _isWatching
}
func setWatching(val bool) {
	_isWatching = val
}
