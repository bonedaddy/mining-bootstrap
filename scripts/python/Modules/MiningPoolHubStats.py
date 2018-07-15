from urllib.request import Request, urlopen
import json

# Author:       Postables
# Version:      0.0.1alpha
# Description:  Python module to fetch statistics from miningpoolhub designed to be called by PostablesMiningBot (Telegram bot)

class MiningPoolHubStatistics():

    def __init__(self):
        self.url_template = 'https://{coin}.miningpoolhub.com/index.php?page=api&action={action}&api_key={api_key}'
        self.user_agent = 'Mozzila/5.0'

    def construct_url(self, _coin, _action, _api_key):
        self.constructed_url = self.url_template.format(coin=_coin, action=_action, api_key=_api_key)
        return self.constructed_url

    def fetch_data(self, _coin, _action, _api_key):
        formatted_url = self.construct_url(_coin, _action, _api_key)
        request_data = Request(formatted_url, headers={'User-Agent': '%s' % self.user_agent})
        json_object = json.load(urlopen(request_data))
        return json_object[_action]['data']

    def get_recent_credits(self, _coin, _api_key):
        # covers the previous two weeks
        data = self.fetch_data(_coin, 'getdashboarddata', _api_key)
        recent_credits = data['recent_credits']
        credits_dict = {}
        for i in recent_credits:
            amount = i['amount']
            date = i['date']
            credits_dict[date] = amount
        return credits_dict
    
    def get_hashrate(self, _coin, _api_key):
        data = self.fetch_data(_coin, 'getdashboarddata', _api_key)
        hash_rate = data['personal']['hashrate']
        return hash_rate

    def get_user_workers(self, _coin, _api_key):
        data = self.fetch_data(coin, 'getuserworkers', _api_key)
        return data

    def get_offline_workers(self, _coin, _api_key):
        data = self.get_user_workers(_coin, _api_key)
        offline_workers = []
        for i in range(0, len(data)):
            if data[i]['hashrate'] == 0:
                first_half, second_half = data[i]['username'].split('.')
                offline_workers.append(second_half)
        if len(offline_workers) > 0:
            msg = ''
            for i in offline_workers:
                msg += '%s ' % i
        final_msg = "The following rigs are offline:\n" % msg
        return final_msg
        
    def get_user_status(self, _coin, _api_key):
        data = self.fetch_data(_coin, 'getusersatus', _api_key)
        return data

    def get_valid_shares(self, _coin, _api_key):
        data = self.get_user_status(_coin, _api_key)
        return data['valid']

    def get_invalid_shares(self, _coin, _api_key):
        data = self.get_user_status(_coin, _api_key)
        return data['invalid']
