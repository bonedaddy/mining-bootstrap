from Modules import MiningPoolHubStats
import sys

if len(sys.argv) < 4:
        print("Invalid invocation")
        print("python3.5 rtrade.py <coin> <api_key> <mode>")
        exit()

coin = sys.argv[1]
api_key = sys.argv[2]
mode = sys.argv[3]
m = MiningPoolHubStats.MiningPoolHubStatistics()
supported_modes = ['--recent-credits']


if mode == '--recent-credits':
        recent_credits_dict = m.get_recent_credits(coin, api_key)
        # lets invert the dictionary
        recent_credits_dict_keys = list(recent_credits_dict.keys())
        csv_header = "date,amount,usd value"
        with open('output.csv', 'a') as fh:
                fh.write('%s\n' % csv_header)

        with open('output.csv', 'a') as fh:
                for key in reversed(recent_credits_dict_keys):
                        msg = '%s,%s,\n' % (key, recent_credits_dict[key])
                        fh.write('%s' % msg)
else:
        msg = ""
        for e in supported_modes:
                msg += "%s " % e
        print("Invalid mode, please use one of the following modes:\n%s\n" % msg)
