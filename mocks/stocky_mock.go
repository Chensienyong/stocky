package mocks

// StockyMock is a mock for stocky
type StockyMock struct {
	Postgres      PostgresMock
	Redis         RedisMock
	AlphaVantage  AlphaVantageMock
}
