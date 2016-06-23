package OkApi

type Credential struct {
	AppKey *string
	Secret *string
}

func NewCredential(appKey string, secret string) *Credential {
	return &Credential{
		AppKey: &appKey,
		Secret: &secret,
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