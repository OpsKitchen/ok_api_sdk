package sdk

type Credential struct {
	AppKey    string
	Secret    string
	SessionId string
}

func NewCredential() *Credential {
	return &Credential{
	}
}