package adapters

type TokenAdapter interface {
	GenerateToken(username string, email string) (string, error)
}
