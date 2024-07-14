package cloud

type Provider interface {
	Initialize() error

	PreCreate() error
	Create() error
	PostCreate() error

	PreDelete() error
	Delete() error
	PostDelete() error
}
