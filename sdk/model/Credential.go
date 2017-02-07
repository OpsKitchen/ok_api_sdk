package model

type Credential struct {
	AppKey    string
	DeviceId  string
	Secret    string
	SessionId string
}

func (credential *Credential) SetAppKey(appKey string) *Credential {
	credential.AppKey = appKey
	return credential
}

func (credential *Credential) SetDeviceId(deviceId string) *Credential {
	credential.DeviceId = deviceId
	return credential
}

func (credential *Credential) SetSecret(secret string) *Credential {
	credential.Secret = secret
	return credential
}

func (credential *Credential) SetSessionId(sessionId string) *Credential {
	credential.SessionId = sessionId
	return credential
}
