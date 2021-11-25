# A Documentation
A library to manage Blocks

## Structs

### Block
```
type Block struct {
    Hash            cipher.SHA256
    SequenceNo      unint64
    Depth           unint64             // The depth will be the position at which this block is stored in tree.
    Parent          *Block
    Children        []*Block
    
    Transactions    []cipher.SHA256              // List of transaction ids
	UxIdSpent       map[cipher.SHA256]*[]btye    // outputs spent / destroyed
	UxIdCreated     map[cipher.SHA256]*[]btye    // ouputs being created
}
```

### BlockTree
In order to use the library we will create a BlockTree instance. Each block tree instance holds a tree of blocks and exposes various methods to interact with them. 
```
type BlockTree struct {
    Root               *Block                               // Serves as the root of the block tree
    TransactionsMap    map[cipher.SHA256]*Transactions      // Global TxId to *Transactions Map (Transactions struct is defined in [coin/Transactions.go]) 
}
```


## Routines

### CheckIfUnspentOutputExistsInSpent
- Checks if an unspent output is spent in a block
```
func (b *Block) CheckIfUnspentOutputExistsInSpent(uxId cipher.SHA256) string {
}
```

### CheckIfUnspentOutputExistsInCreated
- Checks if an unspent output is created in a block
```
func (b *Block) CheckIfUnspentOutputExistsInCreated(uxId cipher.SHA256) string {
}
```


### AddBlock
- Adds a block to the block tree
```
func (bt *BlockTree) AddBlock(b *Block) string {
}
```

### GetBlockDepth
- Returns the number of blocks between the given block and the root of the block tree
```
func (bt *BlockTree) GetBlockDepth(b *Block) uint64 {
}
```

### GetAllBlocks
- Returns an list of all blocks from the root to the end of the tree. This function performs a *depth first traversal* of the whole tree returns the list of all blocks 
  in the order they are found. 
```
func (bt *BlockTree) GetAllBlocks() []*Block {
}
```

### CheckIfUnspectOutputExists
- Traverses a tree from root to the given Block can check if the unspent out was destroyed on it's way from Root to the given block.
- This function can *return* any of the following codes:
  -> "NeverExisted": The unspent output never existed 
  -> "Spent": The unspent output exists but was spent 
  -> "Available": The unspent output is available to spend on the given block
  
```
func (bt *BlockTree) CheckIfUnspectOutputExists(UnspectOutput cipher.SHA256, targetBlock *Block) string {
}
```


