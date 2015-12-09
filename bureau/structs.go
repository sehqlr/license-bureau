package bureau

type SoftwareLicense struct {
	Local string
	Name  string
	Web   string
}

type Component struct {
	License SoftwareLicense
	URL     string
}
