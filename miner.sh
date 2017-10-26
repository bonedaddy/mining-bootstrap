#! /bin/bash

# Author: Postables
# Version: 0.0.1alpha
# Description: Script is used to launch ccminer, and provide configurable variabels for easy management

rigName="$(hostname)"
poolUrl="stratum+tcp://hub.miningpoolhub.com:20507"
poolUser="postables.$rigName"
poolPassword="password"
minerAlgo="lyra2v2"

ccminer --url="$poolUrl" --user "$poolUser" --pass "$poolPassword" --algo="$minerAlgo"