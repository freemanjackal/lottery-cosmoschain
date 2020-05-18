## Instruction

1. `Make install` to get `lotd` and `lotcli`
2. Run single validator by `lotd start --rpc.laddr=tcp://0.0.0.0:26657 --pruning=nothing`

## Setup relayer

setup a relayer

## 

0. Set up channel in lottery chain by lotcli

```
lotcli tx lottery set-channel bandchain lottery <channel_id_of_goldcdp_goldchin> --from validator --keyring-backend test

```

0.5 Get atom from faucet

```
curl --location --request POST 'http://gaia-ibc-hackathon.node.bandchain.org:8000' \
--header 'Content-Type: application/javascript' \
--data-raw '{
 "address": <your_address>,
 "chain-id": "band-cosmoshub"
}'
```

1. Transfer coin from gaia to bandchain

```
(bccli|gaiacli) tx transfer transfer transfer <channel_id_of_gaia> 10000000 <account_in_gold_chain> 800000000transfer/<channel_id_of_gold_chain>/uatom --from <account_in_gaia> --node http://gaia-ibc-hackathon.node.bandchain.org:26657 --keyring-backend test --chain-id band-cosmoshub
```

2. Send buy transaction

```
lotcli tx lottery create 
