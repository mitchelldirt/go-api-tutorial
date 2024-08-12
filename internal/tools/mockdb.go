package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"testuser": LoginDetails{
		AuthToken: "testtoken",
		Username:  "testuser",
	},
	"jason": LoginDetails{
		AuthToken: "jasonstoken",
		Username:  "jason",
	},
	"marie": LoginDetails{
		AuthToken: "mariestoken",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"testuser": CoinDetails{
		Coins:    100,
		Username: "testuser",
	},
	"jason": CoinDetails{
		Coins:    1000,
		Username: "jason",
	},
	"marie": CoinDetails{
		Coins:    10000,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
