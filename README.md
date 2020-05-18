# Lottery
Lottery works as a simple lottery system where a lottery is created. Once created users can bet on any number in the open lottery.
When lottery is closed is consulted random number from band oracles, from that is determined the winning number.  All funds are distributed to the winners. In the case there are no winners funds are accumulated and a new lottery is started.


## No front end implementation yet.

## Instruction

1. `Make install` to get `lotd` and `lotcli`

    1.1. `lotd init validator --chain-id lottery` ...
    * Configure the node, take a look at this article form more [details](https://blog.cosmos.network/guide-to-building-defi-using-band-protocol-oracle-and-cosmos-ibc-fa5348832f84)

2. Run single validator by `lotd start --rpc.laddr=tcp://0.0.0.0:26657 --pruning=nothing`

## Setup relayer

Information about how to setup a [relayer](https://blog.cosmos.network/guide-to-building-defi-using-band-protocol-oracle-and-cosmos-ibc-fa5348832f84)

### 

0. Set up channel in lottery chain by lotcli after set up the relayer


1. Create lottery transaction
When created the price per bets/plays is setted automatically(a fixed value)

```
lotcli tx lottery create --from [user]
```
2. Playing lottery
```
lotcli tx lottery bet [number] [price] --from [user]
```
3. Closing lottery
There will be just 1 open lottery at any time. When the open lottery is closed funds are distributed and a new lottery is created
```
lotcli tx lottery close --from [user]
```

TODO:
- Set lottery price by user input
- Improve structure and model of the stored information
- Front end



