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
	fmt.Println("createblockchain  -address string                        ----> 输入数据创建一个创建一个创世区块")
	fmt.Println("view                                       ----> 查看链中的所有区块交易信息")
	fmt.Println("indexblock -key []byte{}  没有实现                      ----> 输入区块hash，返回一个区块的信息")
	fmt.Println("send     -from []string  -to []string -amount []string                        ----> 输入数据创建一个新的区块 string")
	fmt.Println("stop                                    ----> 输入数据创建一个新的区块 string")
	fmt.Println("getbalance -address string                                   ----> 输入数据创建一个新的区块 string")
	fmt.Println("--------------------------------------------------------------------------------------------------------------")
}

func (cli *CommandLine) createBlockChain(address string) {
	CreateBlockChain(address)
	//fmt.Println("创世区块的hash是:", blockchain.Tip)
}

func (cli *CommandLine) view() {
	chain := ReturnChain()
	chain.ViewChainData()
}
func (cli *CommandLine) send(from, to, amount []string) {
	//fmt.Println("传入的form:", from)
	//fmt.Println("传入的to:", to)
	//fmt.Println("传入的amount:", amount)
	//fmt.Println("传入的数据是:", txns)
	chain := ReturnChain()
	//println("开始创建交易组")
	transactions := chain.CreateTransactions(from, to, amount)
	println(len(transactions))
	if len(transactions) == 0 {
		println("传入的交易无效，无法进行挖矿")
	} else {
		println("开始挖矿")
		chain.Mine(transactions)
	}

}
func (cli *CommandLine) indexblock(key []byte) {
	chain := ReturnChain()
	//fmt.Println("输入的has是:", key)
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

func (cli *CommandLine) getBalance(address string) {
	println("输入的地址是：", address)
	chain := ReturnChain()
	balance, _ := chain.GetBalance(address)
	fmt.Println(address, "该地址拥有的金额为：", balance)
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	CreateBlockChainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	flag.NewFlagSet("view", flag.ExitOnError)
	flag.NewFlagSet("stop", flag.ExitOnError)
	getBlockChainInfoCmd := flag.NewFlagSet("send", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("indexblock", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getbalance", flag.ExitOnError)

	//createBlockChainOwner := createBlockChainCmd.String("data", "", "The address refer to the owner of blockchain")
	//balanceAddress := balanceCmd.String("address", "", "Who need to get balance amount")
	sendFromAddress := sendCmd.String("key", "", "输入区块hash值")
	createChain := CreateBlockChainCmd.String("address", "Daf", "输入创世区块初始化货币的地址")
	fromaddress := getBlockChainInfoCmd.String("from", "发送者", "发送者数组")
	toaddress := getBlockChainInfoCmd.String("to", "接收者", "接收者数组")
	amounts := getBlockChainInfoCmd.String("amount", "区块数据", "发送金额数组")

	hexAddress := getBalanceCmd.String("address", "Daf", "输入地址")
	//sendToAddress := sendCmd.String("to", "", "Destination address")
	//sendAmount := sendCmd.Int("amount", 0, "Amount to send")

	switch os.Args[1] {
	case "createblockchain":
		err := CreateBlockChainCmd.Parse(os.Args[2:])
		Handle(err)

	case "view":
		cli.view()
		//err := balanceCmd.Parse(os.Args[2:])
		//Handle(err)

	case "send":
		err := getBlockChainInfoCmd.Parse(os.Args[2:])
		Handle(err)

	case "getbalance":
		err := getBalanceCmd.Parse(os.Args[2:])
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

	if getBlockChainInfoCmd.Parsed() {
		if *fromaddress == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		if *toaddress == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		if *amounts == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}

		//fmt.Println("输入的数据为：", *fromaddress)
		from := JsonToArray(*fromaddress)
		to := JsonToArray(*toaddress)
		amount := JsonToArray(*amounts)

		if len(from) == len(to) && len(to) == len(amount) {
			fmt.Println("传入数据正确")
			cli.send(from, to, amount)
		} else {
			fmt.Println("传入的数据有错误")
		}
		//fmt.Println(from[0])
		//fmt.Println(to[0])
		//fmt.Println(amount[0])

		//cli.send(from, to, amount)

	}

	if CreateBlockChainCmd.Parsed() {
		if *createChain == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.createBlockChain(*createChain)
	}

	if getBalanceCmd.Parsed() {
		if *hexAddress == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		fmt.Println("hexAddress的值是", *hexAddress)
		cli.getBalance(*hexAddress)
	}

	if sendCmd.Parsed() {
		if *sendFromAddress == "" {
			sendCmd.Usage()
			runtime.Goexit()
		}
		cli.indexblock([]byte(*sendFromAddress))
	}
}
