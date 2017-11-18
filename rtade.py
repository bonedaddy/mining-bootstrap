import miningPoolHubStatisticFether
import sys

if len(sys.argv) < 3:
        print("Invalid invocation")
        print("python3.5 rtrade.py <coin> <api_key>")
        exit()

coin = sys.argv[1]
api_key = sys.argv[2]
m = miningPoolHubStatisticFether.MiningPoolHubStatistics()

recent_credits_dict = m.get_recent_credits(coin, api_key)

csv_header = "date,amount,usd valie"
with open('output.csv', 'a') as fh:
        fh.write('%s\n' % csv_header)

with open('output.csv', 'a') as fh:
        for key in recent_credits_dict.keys():
                msg = '%s,%s,\n' % (key, recent_credits_dict[key])
                fh.write('%s' % msg)

