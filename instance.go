package clamavrestsdk

type clamAVRestInstance struct {
	apiAddress string
	apiPort    int
	apiIsHttps bool
}

func NewclamAVRestInstance(
	apiAddress string,
	apiPort int,
	apiIsHttps bool,
) *clamAVRestInstance {
	return &clamAVRestInstance{
		apiAddress,
		apiPort,
		apiIsHttps,
	}
}
