package main

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*

需求：
	假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。
	在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。
	如果余额不足，则回滚事务。

建表语句：
	-- 账户表
	CREATE TABLE accounts (
		id INT PRIMARY KEY AUTO_INCREMENT,
		balance DECIMAL(15,2) NOT NULL DEFAULT 0.00,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

	-- 交易记录表
	CREATE TABLE transactions (
		id INT PRIMARY KEY AUTO_INCREMENT,
		from_account_id INT NOT NULL,
		to_account_id INT NOT NULL,
		amount DECIMAL(15,2) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/

type Accounts struct {
	Id        int `gorm:"primaryKey"`
	Balance   float32
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Transactions struct {
	Id            int `gorm:"primaryKey"`
	FromAccountId int
	ToAccountId   int
	Amount        float32
	CreatedAt     *time.Time
}

func (account *Accounts) queryBalance(db *gorm.DB) float32 {
	db.Model(account).Find(account)
	return account.Balance
}

// 主要为了练习ORM库事务操作，暂不考虑余额的并发锁操作
func (fromAccount *Accounts) transfer(db *gorm.DB, toAccount *Accounts, amount float32) bool {
	err := db.Transaction(func(tx *gorm.DB) error {

		// 余额不足
		if fromAccount.queryBalance(tx) <= amount {
			return errors.New("余额不足")
		}

		if err := tx.Model(fromAccount).Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		if err := tx.Model(toAccount).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		transactions := Transactions{
			FromAccountId: fromAccount.Id,
			ToAccountId:   toAccount.Id,
			Amount:        amount,
		}
		if err := tx.Create(&transactions).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		fmt.Println("转账失败", err)
		return false
	}

	return true
}

func initData(db *gorm.DB) []Accounts {
	// 创建用户
	accounts := []Accounts{
		{Balance: 1000},
		{Balance: 1000},
	}

	db.CreateInBatches(accounts, len(accounts))
	return accounts
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 初始化数据
	accounts := initData(db)

	// 查询余额
	fmt.Printf("A的余额：%v, B的余额：%v \n", accounts[0].queryBalance(db), accounts[1].queryBalance(db))

	// 转账100元
	result := accounts[0].transfer(db, &accounts[1], 100)

	if result {
		fmt.Println("转账成功")
	} else {
		fmt.Println("转账失败")
	}

	// 查询余额
	fmt.Printf("A的余额：%v, B的余额：%v \n", accounts[0].queryBalance(db), accounts[1].queryBalance(db))

	// 转账901元
	result = accounts[0].transfer(db, &accounts[1], 901)

	if result {
		fmt.Println("转账成功")
	} else {
		fmt.Println("转账失败")
	}

	// 查询余额
	fmt.Printf("A的余额：%v, B的余额：%v \n", accounts[0].queryBalance(db), accounts[1].queryBalance(db))
}
