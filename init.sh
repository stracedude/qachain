#!/usr/bin/env bash

 
#!/usr/bin/env bash

rm -rf ~/.qacd
rm -rf ~/.qacli

qacd init test --chain-id=qachain

qacli config output json
qacli config indent true
qacli config trust-node true
qacli config chain-id qachain
qacli config keyring-backend test

qacli keys add jack
qacli keys add alice
qacli keys add moder1
qacli keys add moder2
qacli keys add moder3

qacd add-genesis-account $(qacli keys show jack -a) 1000qatoken,100000000stake
qacd add-genesis-account $(qacli keys show alice -a) 1000qatoken,100000000stake
qacd add-genesis-account $(qacli keys show moder1 -a) 1000qatoken,100000000stake
qacd add-genesis-account $(qacli keys show moder2 -a) 1000qatoken,100000000stake
qacd add-genesis-account $(qacli keys show moder3 -a) 1000qatoken,100000000stake

qacd gentx --name jack --keyring-backend test

echo "Collecting genesis txs..."
qacd collect-gentxs

echo "Validating genesis file..."
qacd validate-genesis