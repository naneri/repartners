package pack

type PackSize struct {
	Size int
}

func NewPack(size int) PackSize {
	return PackSize{
		Size: size,
	}
}
