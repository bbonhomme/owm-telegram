package owm

import "errors"

// setKey will initiate the api key for the request
func setKey(key string) (string, error) {
	if err := ValidAPIKey(key); err != nil {
		return "", err
	}
	return key, nil
}

// ValidAPIKey makes sure that the key given is a valid one
func ValidAPIKey(key string) error {
	if len(key) != 32 {
		return errors.New("invalid key")
	}
	return nil
}
