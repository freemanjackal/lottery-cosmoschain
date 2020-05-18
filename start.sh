rm -rf ~/.lot*

lotd init validator --chain-id lottery
lotcli keys add validator --keyring-backend test
echo "smile stem oven genius cave resource better lunar nasty moon company ridge brass rather supply used horn three panic put venue analyst leader comic" | lotcli keys add requester --recover --keyring-backend test
echo "clutch amazing good produce frequent release super evidence jungle voyage design clip title involve offer brain tobacco brown glide wire soft depend stand practice" | lotcli keys add relayer --recover --keyring-backend test


lotd add-genesis-account validator 10000000000000loto --keyring-backend test
lotd add-genesis-account requester 10000000000000loto --keyring-backend test
lotd add-genesis-account relayer 10000000000000loto --keyring-backend test

lotcli config chain-id lottery
lotcli config output json
lotcli config indent true
lotcli config trust-node true

lotd gentx --name validator --keyring-backend test
lotd collect-gentxs

# Run chain
lotd start --rpc.laddr=tcp://0.0.0.0:26657 --pruning=nothing
