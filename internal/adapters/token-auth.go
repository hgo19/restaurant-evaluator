package adapters

type TokenAuth interface {
	Generate(email string) (string, error)
}
