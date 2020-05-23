#!/usr/bin/env bash


#!/bin/bash
qacli tx qac createQuestion 50qatoken "2+2?" --from jack -y  | jq ".txhash" |  xargs $(sleep 5) qacli q tx
qacli tx qac createAnswer 30ad205245cd25231f697fa8934fdf8d65d0cb24040d9f0d4106fa4cfc68e4c7 "4" --from alice -y | jq ".txhash" |  xargs $(sleep 5) qacli q tx
qacli tx qac AnswerVote 4b227777d4dd1fc61c6f884f48641d02b4d121d3fd328cb08b5531fcacdabf8a --from jack -y | jq ".txhash" |  xargs $(sleep 5) qacli q tx

qacli query qac get 30ad205245cd25231f697fa8934fdf8d65d0cb24040d9f0d4106fa4cfc68e4c7
#qachaincli tx qachain revealSolution "A stick" --from you -y | jq ".txhash" |  xargs $(sleep 5) scavengeCLI q tx
qacli query qac list