package adapters

type Encrypter interface {
	HashPassword(password string) (string, error)
}
