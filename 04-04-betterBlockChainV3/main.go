package main

import "fmt"

//mac/linux 下面
//cd 04-00-block-chain/04-01-firstBlockChain
//go run *.go

//windows下面编译运行
//go build
//start xxx.exe
//记得代码后面+ time.Sleep(time.Second*100) or fmt.Scanf()

func main() {
	blockChain := NewBlockChain()
	blockChain.AddBlock("班长向班花转了50枚比特币")

	it := blockChain.NewIterator()
	//调用迭代器
	for {
		//Next 返回当前区块，并且指针左移
		block := it.Next()

		fmt.Printf("===========================\n")
		fmt.Printf("当前区块哈希值： %x\n", block.Hash)
		fmt.Printf("前区块哈希值：%x\n", block.PrevHash)
		fmt.Printf("区块数据： %s\n", block.Data)
		fmt.Printf("当前区块随机数: %d\n", block.Nonce)
		fmt.Printf("当前区块难度值: %d\n", block.Difficulty)
		fmt.Printf("===========================\n\n")

		if len(block.PrevHash) == 0 {
			fmt.Println("区块链编译结束")
			break
		}
	}

	//time.Sleep(time.Second * 100)

}
