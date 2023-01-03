package BLC

import (
	"flag"
	"fmt"
	"os"
	"runtime"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {

	fmt.Println("欢迎来到杜岸峰的区块链...")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("运行本区块链您首先需要创建区块链并生成创世区块.")
	fmt.Println("以下命令供您使用：")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
	fmt.Println("createblockchain                          ----> 输入数据创建一个创建一个创世区块")
	fmt.Println("view                                       ----> 查看链中的所有区块交易信息")
	fmt.Println("indexblock -key []byte{}                        ----> 输入区块hash，返回一个区块的信息")
	fmt.Println("newblock -data string                         ----> 输入数据创建一个新的区块 string")
	fmt.Println("stop                                    ----> 输入数据创建一个新的区块 string")
	fmt.Println("balance -address ADDRESS                            ----> Back the balance of the address you input")
	fmt.Println("blockchaininfo                                      ----> Prints the blocks in the chain")
	fmt.Println("send -from FROADDRESS -to TOADDRESS -amount AMOUNT  ----> Make a transaction and put it into candidate block")
	fmt.Println("mine                                                ----> Mine and add a block to the chain")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
}

func (cli *CommandLine) createBlockChain() {
	blockchain := CreateBlockChain()
	fmt.Println("创世区块的hash是:", blockchain.Tip)
}

func (cli *CommandLine) view() {
	chain := ReturnChain()
	chain.ViewChainData()
}
func (cli *CommandLine) newblock(data string) {
	fmt.Println("传入的数据是:", data)
	chain := ReturnChain()
	chain.AddBlock(data)
}
func (cli *CommandLine) indexblock(key []byte) {
	chain := ReturnChain()
	fmt.Println("输入的has是:", key)
	chain.SingleCheck(key)
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}
func (cli *CommandLine) stop() {
	chain := ReturnChain()
	chain.DBStop()
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	flag.NewFlagSet("createblockchain", flag.ExitOnError)
	flag.NewFlagSet("view", flag.ExitOnError)
	flag.NewFlagSet("stop", flag.ExitOnError)
	getBlockChainInfoCmd := flag.NewFlagSet("newblock", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("indexblock", flag.ExitOnError)

	//createBlockChainOwner := createBlockChainCmd.String("data", "", "The address refer to the owner of blockchain")
	//balanceAddress := balanceCmd.String("address", "", "Who need to get balance amount")
	sendFromAddress := sendCmd.String("key", "", "输入区块hash值")
	datastring := getBlockChainInfoCmd.String("data", "区块数据", "数据交易数据")
	//sendToAddress := sendCmd.String("to", "", "Destination address")
	//sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "createblockchain":
		cli.createBlockChain()
		//err := createBlockChainCmd.Parse(os.Args[2:])
		//Handle(err)

	case "view":
		cli.view()
		//err := balanceCmd.Parse(os.Args[2:])
		//Handle(err)

	case "newblock":
		err := getBlockChainInfoCmd.Parse(os.Args[2:])
		Handle(err)

	case "indexblock":
		err := sendCmd.Parse(os.Args[2:])
		Handle(err)

	case "stop":
		cli.stop()

	default:
		cli.printUsage()
		runtime.Goexit()
	}

	//if sendCmd.Parsed() {
	//	if *sendFromAddress == "" || *sendToAddress == "" || *sendAmount <= 0 {
	//		sendCmd.Usage()
	//		runtime.Goexit()
	//	}
	//	cli.send(*sendFromAddress, *sendToAddress, *sendAmount)
	//}

	if getBlockChainInfoCmd.Parsed() {
		if *datastring == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.newblock(string(*datastring))
	}

	if sendCmd.Parsed() {
		if *sendFromAddress == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.indexblock([]byte(*sendFromAddress))
	}
}
