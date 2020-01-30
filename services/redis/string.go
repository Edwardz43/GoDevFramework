package redis

// Set sets key to hold the string value.
func (s *Service) Set(key, value string, ttl int) (string, error) {
	return s.Client.Do("SETEX", key, ttl, value).String()
}

// Get the value of key.
func (s *Service) Get(key string) (string, error) {
	return s.Client.Do("GET", key).String()
}

//GetSet atomically sets key to value and returns the old value stored at key. Returns an error when key exists but does not hold a string value.
func (s *Service) GetSet(key, value string) (string, error) {
	return s.Client.Do("GETSET", key, value).String()
}

// IncrBy increases the number stored at key by increment.
func (s *Service) IncrBy(key, value string) (int64, error) {
	return s.Client.Do("INCRBY", key, value).Int64()
}

// DecrBy decreases the number stored at key by decrement.
func (s *Service) DecrBy(key, value string) (int64, error) {
	return s.Client.Do("DECRBY", key, value).Int64()
}
