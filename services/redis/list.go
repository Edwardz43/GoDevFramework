package redis

// Len returns the length of the list stored at key.
func (s *Service) Len(key string) (int, error) {
	return s.Client.Do("LLEN", key).Int()
}

// Pop removes and returns the last element of the list stored at key.
func (s *Service) Pop(key string) (string, error) {
	return s.Client.Do("LPOP", key).String()
}

// Push inserts all the specified values at the head of the list stored at key.
func (s *Service) Push(key, value string) (int, error) {
	return s.Client.Do("RPUSH", key, value).Int()
}

// PopPush atomically returns and removes the last element (tail) of the list stored at source, and pushes the element at the first element (head) of the list stored at destination.
func (s *Service) PopPush(orginKey, destKey string) (string, error) {
	result := s.Client.RPopLPush(orginKey, destKey)
	return result.Result()
}

// LGet returns the element at index index in the list stored at key.
func (s *Service) LGet(index int) (string, error) {
	return s.Client.Do("LINDEX", index).String()
}

// Retrieve returns the specified elements of the list stored at key.
func (s *Service) Retrieve(key string, start, stop int64) ([]string, error) {
	result := s.Client.LRange(key, start, stop)
	return result.Result()
}
