package apiclients

type ApiClients struct {
	BookClient *BookClient
	UserClient *UserClient
}

func NewApiClients(bc *BookClient, uc *UserClient) *ApiClients {
	return &ApiClients{
		BookClient: bc,
		UserClient: uc,
	}
}
