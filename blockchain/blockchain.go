package blockchain

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Data 			map[string]interface{}
	Hash 			string
	PreviousHash 	string
	Timestamp 		time.Time
	Pow 			int
}

type Blockchain struct {
	GenesisBlock 	Block
	Chain 			[]Block
	Difficulty 		int
}

func NewBlockchain(difficulty int) *Blockchain {
	genesisBlock := Block{
		Hash: "0",
		Timestamp: time.Now(),
	}

	return &Blockchain{
		GenesisBlock: genesisBlock,
		Chain: []Block{genesisBlock},
		Difficulty: difficulty,
	}
}

func (b *Blockchain) AddBlock(from string, to string, amount float64) {
	blockData := map[string]interface{}{
		"from": from,
		"to": to,
		"amount": amount,
	}
	lastBlock := &b.Chain[len(b.Chain) - 1]
	newBlock := Block{
		Data: blockData,
		PreviousHash: lastBlock.Hash,
		Timestamp: time.Now(),
	}
	newBlock.mine(b.Difficulty)
	b.Chain = append(b.Chain, newBlock)
}

func (b Blockchain) IsValid() bool {
	for i := range b.Chain[1:] {
		previousBlock := b.Chain[i]
		currentBlock := b.Chain[i+1]
		if currentBlock.Hash != currentBlock.calculateHash() || currentBlock.PreviousHash != previousBlock.Hash {
			return false
		}
	}
	return true
}

func (b *Block) calculateHash() string  {
	data,_ := json.Marshal(b.Data)
	blockData := b.PreviousHash + string(data) + b.Timestamp.String() + strconv.Itoa(b.Pow)
	blockHash := sha1.Sum([]byte(blockData))
	return fmt.Sprintf("%x", blockHash)
}

func (b *Block) mine(difficulty int) {
	prefix := strings.Repeat("0", difficulty)

	for !strings.HasPrefix(b.Hash, prefix) {
		b.Pow++
		b.Hash = b.calculateHash()
	}
}