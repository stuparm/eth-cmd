# evm storage

The purpose of this repository is to analyze storage usage at the block and transaction levels. Specifically, it aims to evaluate the extent of storage redundancy across transactions within a single block.

To illustrate a simplified, hypothetical scenario, consider a block containing four transactions:

- tx1: A sends ETH to B
- tx2: B sends ETH to A
- tx3: A sends ETH to B
- tx4: B sends ETH to A
In this case, while the block contains four transactions, only two storage values are modified:
- A.balance
- B.balance
  
This concept becomes intesting when considering widely-used protocols like Uniswap. For example, if multiple transactions in a block invoke the swap() function on the same liquidity pool, they might access and modify the same storage slot. These transactions will touch and potentionally change the same storage slot and I would like to analyze those.