# earnings-report

The following pools are Supported:

* [x] MiningPoolHub
* [x] Ethermine (ETC+ETH)
* [x] Ethpool
* [x] Flypool (ZCash)

## MiningPoolHub Usage

To generate a report of the earnings for the last 24 hour period, saving the information to your database, and sending an email:
1)  `export CONFIG_PATH=<path-config-file>`
2) `export RUN_MODE=report-save`
3) `./earnings-report mph`

To generate a report of the eranigns for the last 24 hour period, and diplsay them to STDOUT:
1) `export CONFIG_PATH=<path-cofig-file>`
2) `export RUN_MODE=report`
3) `./earnings-report mph`

## Ethermine, Ethpool, Flypool Usagee

To display all payouts in STDOUT, pretty in a friendly manner:
1) `export CONFIG_PATH=<path-config-file>`
2) `export RUN_MODE=payouts`
3) `./earnings-report ethermine payouts <miner-to-lookup>`

## MiningPoolHub Supported Features

* Get Recent Credits
* Get Recent Credits 24 Hours (local system time should be UTC)
* Save to database
* Send report email via sendgrid

## Ethermine, Ethpool, Flypool Supported Features

* Get all payouts (usually limited to the most recent 100)

## Config File

* MiningPoolHub URL should be `https://%s.miningpoolhub.com/index.php?page=api&action=%s&api_key=%s`
* Ethermine ETH URL should be `https://api.ethermine.org`
* Ethermine ETC URL should be `https://api-etc.ethermine.org`
* Flypool URL should be `https://api-zcash.flypool.org`
* Ethpool URL should be `http://api.ethpool.org`

Example config file
```JSON
{
        "coin": "ethereum",
        "url": "https://%s.miningpoolhub.com/index.php?page=api&action=%s&api_key=%s",
        "mph_api_key": "nil",
        "sendgrid_api_key": "nil",
        "db_url": "nil",
        "db_user": "nil",
        "db_pass": "nil"
}
```