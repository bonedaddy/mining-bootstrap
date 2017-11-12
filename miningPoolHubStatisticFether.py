from urllib.request import Request, urlopen
import json
# req = Request(url_template, headers={'User-Agent': 'Mozzila/5.0})
# webpage = urlopen(req).read()
# webpage

class MiningPoolHubStatistics():

    def __init__(self):
        self.url_template = 'https://{coin}.miningpoolhub.com/index.php?page=api&action={action}&api_key={api_key}'
        self.user_agent = 'Mozzila/5.0'

    def construct_url(self, _coin, _action, _api_key):
        self.constructed_url = url_template.format(coin=_coin, action=_action, api_key=_api_key)
        return str(self.construced_url)

    def fetch_data(self, _coin, _action, _api_key):
        formatted_url = construct_url(_coin, _action, _api_key)
        request_data = Request(formatted_url, headers={'User-Agent': '%s' % self.user_agent})
        json_object = json.load(urlopen(request_data))
        return json_object[action]['data']

    def retrieve_recent_credits(self, _coin, _api_key):
        # covers the previous two weeks
        data = fetch_data(_coin, 'getdashboarddata', _api_key)
        recent_credits = data['recent_credits']
        for i in recent_credits:
            print(i['amount'])