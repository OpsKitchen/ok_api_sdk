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

func (c *Credential) SetAppKey(appKey string) *Credential {
	c.AppKey = appKey
	return c
}

func (c *Credential) SetSecret(secret string) *Credential {
	c.AppKey = secret
	return c
}

func (c *Credential) SetSessionId(sessionId string) *Credential {
	c.AppKey = sessionId
	return c
}
