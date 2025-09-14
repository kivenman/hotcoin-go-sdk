package main

import (
	"fmt"
	"log"

	hotcoin "github.com/hotcoin/go-sdk"
)

func main() {
	fmt.Println("=== HOTCOIN Go SDK 示例 ===")

	// 创建客户端
	client := hotcoin.NewClient("your_api_key", "your_secret_key")
	client.SetDebug(true)

	// 示例1：获取合约列表
	fmt.Println("\n1. 获取合约列表:")
	contracts, err := client.Market.GetContracts("")
	if err != nil {
		log.Printf("获取合约列表失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 个合约\n", len(contracts))
		if len(contracts) > 0 {
			fmt.Printf("第一个合约: %s\n", contracts[0].Code)
		}
	}

	// 示例2：获取K线数据
	fmt.Println("\n2. 获取K线数据:")
	klines, err := client.Market.GetKline("solusdt", "1min", 5)
	if err != nil {
		log.Printf("获取K线数据失败: %v", err)
	} else {
		fmt.Printf("获取到 %d 条K线数据\n", len(klines))
		if len(klines) > 0 {
			fmt.Printf("最新K线: 开盘价=%s, 收盘价=%s, 成交量=%s\n",
				klines[len(klines)-1].Open, klines[len(klines)-1].Close, klines[len(klines)-1].Volume)
		}
	}

	// 示例3：获取深度数据
	fmt.Println("\n3. 获取深度数据:")
	depth, err := client.Market.GetDepth("solusdt", "")
	if err != nil {
		log.Printf("获取深度数据失败: %v", err)
	} else {
		fmt.Printf("买盘档位: %d, 卖盘档位: %d\n", len(depth.Bids), len(depth.Asks))
		if len(depth.Bids) > 0 && len(depth.Asks) > 0 {
			fmt.Printf("最优买价: %s, 最优卖价: %s\n", depth.Bids[0][0], depth.Asks[0][0])
		}
	}

	// 示例4：获取24小时行情统计
	fmt.Println("\n4. 获取24小时行情统计:")
	tickers, err := client.Market.GetTicker("solusdt")
	if err != nil {
		log.Printf("获取行情统计失败: %v", err)
	} else {
		if len(tickers) > 0 {
			// 查找 SOL/USDT 合约
			for _, ticker := range tickers {
				if ticker.TickerID == "solusdt" {
					fmt.Printf("SOL/USDT 24h统计: 最新价=%s, 最高价=%s, 最低价=%s, 成交量=%s, 涨跌幅=%s%%\n",
						ticker.LastPrice, ticker.High, ticker.Low, ticker.Size24, ticker.Fluctuation)
					break
				}
			}
		}
	}

	fmt.Println("\n注意: 请设置真实的API密钥来测试需要认证的接口")
	fmt.Println("更多示例请参考 README.md 文档")
}
