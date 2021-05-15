package interfaces

type IHasher interface {
	Hahser(text string) (string, error)
}
