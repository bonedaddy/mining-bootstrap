# earnings-report

This utility is used to parse and collect information from miningpoolhub. Additional pools will be added overtime.

The following need sto be in the "url" field for the config.json file if using miningpool hub
`https://%s.miningpoolhub.com/index.php?page=api&action=%s&api_key=%s`


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