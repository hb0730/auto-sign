package config

type Support interface {
	Do()
	Support(t interface{}) error
}
