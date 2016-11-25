package eventchain


type NotFoundComplainer interface {
	error
	NotFoundComplainer()
}


type internalNotFoundComplainer struct{}


func (internalNotFoundComplainer) Error() string {
	return "Not Found"
}


func (internalNotFoundComplainer) NotFoundComplainer() {
	// Nothing here.
}