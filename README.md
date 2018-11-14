hardfork
---
The repository aims to seperate the bitcoin-abc and bitcoin-sv after the timestamp `1542300000`(compared to mediantime). 

#### Notice:

This repository is under experiment state, so you should not test on your real coins unless you know what you do. **Important, I will not undertake the responsibility if you lost money.**

#### Usage:

1. create a transaction with particular opcode. Only this transacion will not work.

   for example:

   ```
   $ ./create-sv-p2sh --privkey=cTQMnqAB2BRmekZqhXxg9PgxqwN6Dz6pefYrAUrX6HDzX2c9nKXJ --to=bchtest:qqfzckruhgshztq8s3twl3q7rhut9ycvjg4xwgatgq --hash=d985008b856db4566ebbb0fc45df200cd79630e49df8b21f6847766f58fbd740 --idx=0 --value=10000000
   
   // will get:
   010000000140d7fb586f7647681fb2f89de43096d70c20df45fcb0bb6e56b46d858b0085d9000000006a473044022031fa8df85b0f1aa67e53d3fe09d4a25894da5e62a7692acb87da273c96b20728022060c633c3767598296b7290411ccdccd8d4b62da218d6beef6c514c9c92a66b66412103a46b2b307246c570117ef89dafd2b5b2918aff1443efd2013129d32afd3c1ddfffffffff01c29598000000000017a9141dd2057ac46558d7cb422038515d118ca9bd65438700000000
   ```

2. next, spend the utxo from the previous output.

   ```
   $ ./spend-sv-p2sh --privkey=cTQMnqAB2BRmekZqhXxg9PgxqwN6Dz6pefYrAUrX6HDzX2c9nKXJ --to=bchtest:qqfzckruhgshztq8s3twl3q7rhut9ycvjg4xwgatgq --hash=3725333c5fdb082f33af6b1aa60f9c7b33132974e7170073990f89af03c802b5 --idx=0 --value=9999810 --feerate=0.0001 --pkscript=a9141dd2057ac46558d7cb422038515d118ca9bd654387
   
   // will get:
   0100000001b502c803af890f99730017e7742913337b9c0fa61a6baf332f08db5f3c332537000000000651045151959cffffffff01428e9800000000001976a914122c587cba21712c078456efc41e1df8b2930c9288ac00000000
   ```

3. Now, you can broadcast these two transaction via website or your own full node.

   ```
   $ bitcoin-cli sendrawtransaction 010000000140d7fb586f7647681fb2f89de43096d70c20df45fcb0bb6e56b46d858b0085d9000000006a473044022031fa8df85b0f1aa67e53d3fe09d4a25894da5e62a7692acb87da273c96b20728022060c633c3767598296b7290411ccdccd8d4b62da218d6beef6c514c9c92a66b66412103a46b2b307246c570117ef89dafd2b5b2918aff1443efd2013129d32afd3c1ddfffffffff01c29598000000000017a9141dd2057ac46558d7cb422038515d118ca9bd65438700000000
   
   $ bitcoin-cli sendrawtransaction 
   0100000001b502c803af890f99730017e7742913337b9c0fa61a6baf332f08db5f3c332537000000000651045151959cffffffff01428e9800000000001976a914122c587cba21712c078456efc41e1df8b2930c9288ac00000000
   ```

#### Solutions:

1. `Not recommended`. Create a transaction with the opcode `OP_CHECKDATASIGVERIFY `/ `OP_MUL` (in txin). The necessay elements include `message`, `signature` and `public key`. But the generate transaction is not standard, so the full node with default configuration will not accept this transaction.

   related script:

   - create-datasig and spend-datasig for bitcoin-abc
   - create-mul and spend-mul for bitcoin-sv

2. `Recommended` Create transaction with P2SH output script(containing `OP_CHECKDATASIG `/ `OP_MUL` digest). The sigscript for spending the previous transaction will carry opcode `OP_CHECKDATASIG` / `OP_MUL` packaged in redeem script. And it is a standard transaction, all full node with default configuration will relay it.

   related script:

   - create-datasig-p2sh and spend-datasig-p2sh for bitcoin-abc
   - create-sv-p2sh and spend-sv-p2sh for bitcoin-sv

