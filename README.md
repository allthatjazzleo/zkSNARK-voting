# zkSNARK-voting
# zkSNARK-voting

## Install Zokrates
```
curl -LSfs get.zokrat.es | sh
```

## Generate the hash output for voter

run hash program
```
./hash
```
should show similar
```
2022/04/06 01:13:27 secret: YLTqywVHgincVSa1wIic6HuDPCjd8tCoxPKH0k6IXtyRJjJuXqcI225nCbd3q2Lc
2022/04/06 01:13:27 input1: 118697618898887186260444508441000763697
2022/04/06 01:13:27 input2: 158559306682474290596405265725236986735
2022/04/06 01:13:27 input3: 159924270143005557755952963448090544757
2022/04/06 01:13:27 input4: 117560806931828606495696561111043296355
2022/04/06 01:13:27 output1: 74158857617778796911804004967810729339
2022/04/06 01:13:27 output2: 202106172654265455595152822861564877554
```

## Prove knowledge of pre-image

Add hash output1 and output2 from previous step to voting.zok `voter1_hash` value

Can repeat the step if more than 1 voter

compile the program
```
zokrates compile -i voting.zok
```

setup phase and export a verifier smart contract as a Solidity file:
```
zokrates setup
zokrates export-verifier
```

provides the correct pre-image as an argument to the program
```
zokrates compute-witness -a input1 input2 input3 input4
```

run the command to construct the proof
```
zokrates generate-proof
```

copy the verifier.sol to contracts directory
```
cp verifier.sol contracts/verifier.sol
```

## Deploy verifier.sol and Election.sol to remix IDE

1. Deploy verifier.sol first to get its address
2. Optional: one can generate more than one verifier.sol for different set of voters by modifying voting.zok
3. Deploy Election.sol by specify one or more deployed verifier.sol address
4. get the `proof` by `./get_param.sh` and vote
