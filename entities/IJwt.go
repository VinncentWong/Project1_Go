package entities

type IJwt interface {
	GenerateToken() (string, error)
	ValidateToken() error
}
