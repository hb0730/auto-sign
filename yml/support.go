package config

type Support interface {
	Do(AutoSign)
	Support(t interface{}) error
}
