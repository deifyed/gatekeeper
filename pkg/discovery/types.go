package discovery

type Document struct {
	Issuer             string `json:"issuer"`
	EndSessionEndpoint string `json:"end_session_endpoint"`
}
