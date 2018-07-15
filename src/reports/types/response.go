package types

type MiningPoolHubAPIResponse struct {
	Version string                 `json:"version"`
	Runtime float64                `json:"runtime"`
	Data    map[string]interface{} `json:"data"`
}

type RecentCredits struct {
	Date   string  `json:"date"`
	Amount float64 `json:"amount"`
}

type USDResponse struct {
	ExchangeRate float64 `json:"val"`
}

/*
   "recent_credits": [
       {
           "date": "2018-07-15",
           "amount": 0.02531947
*/
/*
{
    "getdashboarddata": {
        "version": "1.0.0",
        "runtime": 753.03101539612,
        "data": {
            "raw": {
                "personal": {
                    "hashrate": 1789277.849
                },
                "pool": {
                    "hashrate": 26521504450.353
                },
                "network": {
                    "hashrate": 248059467743.06,
                    "esttimeperblock": 14
                }
            },
            "personal": {
                "hashrate": 1789.277849,
                "sharerate": 0,
                "sharedifficulty": 0,
                "shares": {
                    "valid": null,
                    "invalid": null,
                    "invalid_percent": 0,
                    "unpaid": 0
                },
                "estimates": {
                    "block": 0,
                    "fee": 0,
                    "donation": 0,
                    "payout": 0
                }
            },
            "balance": {
                "confirmed": 0.06200165,
                "unconfirmed": 0.00305212
            },
            "balance_for_auto_exchange": {
                "confirmed": 0,
                "unconfirmed": 0
            },
            "balance_on_exchange": 0,
            "recent_credits_24hours": {
                "amount": 0.13644299
            },
            "recent_credits": [
                {
                    "date": "2018-07-15",
                    "amount": 0.02079864
                },
                {
                    "date": "2018-07-14",
                    "amount": 0.13677851
                },
                {
                    "date": "2018-07-13",
                    "amount": 0.12787103
                },
                {
                    "date": "2018-07-12",
                    "amount": 0.13114245
                },
                {
                    "date": "2018-07-11",
                    "amount": 0.12561407
                },
                {
                    "date": "2018-07-10",
                    "amount": 0.11707432
                },
                {
                    "date": "2018-07-09",
                    "amount": 0.12143573
                },
                {
                    "date": "2018-07-08",
                    "amount": 0.12425866
                },
                {
                    "date": "2018-07-07",
                    "amount": 0.1386152
                },
                {
                    "date": "2018-07-06",
                    "amount": 0.14021762
                },
                {
                    "date": "2018-07-05",
                    "amount": 0.13512156
                },
                {
                    "date": "2018-07-04",
                    "amount": 0.13317691
                },
                {
                    "date": "2018-07-03",
                    "amount": 0.13336064
                },
                {
                    "date": "2018-07-02",
                    "amount": 0.14924752
                }
            ],
            "pool": {
                "info": {
                    "name": "Ethereum (ETH) Mining Pool Hub",
                    "currency": "ETH"
                },
                "workers": 187711,
                "hashrate": 26.521504450353,
                "shares": {
                    "valid": null,
                    "invalid": null,
                    "invalid_percent": 0,
                    "estimated": 206996950,
                    "progress": 0
                },
                "price": "",
                "difficulty": 16,
                "target_bits": 20
            },
            "system": {
                "load": [
                    17.7,
                    20.11,
                    22.93
                ]
            },
            "network": {
                "hashrate": 248.05946774306,
                "difficulty": "3472832548402771",
                "block": "5966614",
                "esttimeperblock": 14
            }
        }
    }
}



{
    "getuserworkers": {
        "version": "1.0.0",
        "runtime": 186.67006492615,
        "data": [
            {
                "id": 15807822,
                "username": "postables.rig01",
                "password": "x",
                "monitor": 1,
                "hashrate": 320111.129,
                "difficulty": 0
            },
            {
                "id": 15684531,
                "username": "postables.rig02",
                "password": "x",
                "monitor": 1,
                "hashrate": 391180.571,
                "difficulty": 0
            },
            {
                "id": 15689897,
                "username": "postables.rig03",
                "password": "x",
                "monitor": 1,
                "hashrate": 0,
                "difficulty": 0
            },
            {
                "id": 15697746,
                "username": "postables.rig04",
                "password": "x",
                "monitor": 1,
                "hashrate": 0,
                "difficulty": 0
            },
            {
                "id": 15684662,
                "username": "postables.rig05",
                "password": "x",
                "monitor": 1,
                "hashrate": 358333.349,
                "difficulty": 0
            },
            {
                "id": 15684672,
                "username": "postables.rig06",
                "password": "x",
                "monitor": 1,
                "hashrate": 438361.122,
                "difficulty": 0
            },
            {
                "id": 15697741,
                "username": "postables.rig07",
                "password": "x",
                "monitor": 1,
                "hashrate": 319513.904,
                "difficulty": 0
            },
            {
                "id": 15697754,
                "username": "postables.rig08",
                "password": "x",
                "monitor": 1,
                "hashrate": 0,
                "difficulty": 0
            }
        ]
    }
}


{
    "getuserstatus": {
        "version": "1.0.0",
        "runtime": 881.93893432617,
        "data": {
            "username": "postables",
            "shares": {
                "valid": 145065.79089355,
                "valid1": 1121.3124084473,
                "valid2": 0,
                "valid3": 0,
                "valid4": 0,
                "valid5": 0,
                "valid6": 0,
                "invalid": 0,
                "id": 142320,
                "donate_percent": 0,
                "is_anonymous": 1,
                "username": "postables"
            },
            "hashrate": 1698500.071,
            "sharerate": 0
        }
    }
}
*/
