package eventchainregistry


var (
	Registry Registrar
)


func init() {
	Registry = new(internalRegistrar)
}
