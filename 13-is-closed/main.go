func IsClosed(c chan T) bool {
	select {
	case <-c:
		return true
	default:
	}
	return false
}
