package dbclient

import (
	"../model"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
)

type IBoltClient interface {
	OpenBoltDb()
	QueryAccount(accountId string) (model.Account, error)
	Seed()
	Check() bool
	Close()
}

//reel implementation
type BoltClient struct {
	boldDB *bolt.DB
}

func (bc *BoltClient) OpenBoltDb() {
	var err error
	bc.boldDB, err = bolt.Open("accounts.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
}

//Start seeding account
func (bc *BoltClient) Seed() {
	bc.initializeBucket()
	bc.seedAccounts()
}

//Create an "AccountBucket" in our BoltDB.will overrite existing bucket of the same name
func (bc *BoltClient) initializeBucket() {
	bc.boldDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("AccountBucket"))
		if err != nil {
			return fmt.Errorf("create bucket failed: %s", err)
		}
		return nil
	})
}

//Seed (n) make-believe account objects into the AccountBucket bucket.
func (bc *BoltClient) seedAccounts() {
	total := 100
	for i := 0; i < total; i++ {
		//Generate a key 10000 or larger
		key := strconv.Itoa(10000 + i)

		//Create a instance of our Account struce
		acc := model.Account{
			Id:   key,
			Name: "Person_" + strconv.Itoa(i),
		}

		//Serialize the struct to JSON
		jsonBytes, _ := json.Marshal(acc)

		//Write the data to AccountBucke
		bc.boldDB.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("AccountBucket"))
			err := b.Put([]byte(key), jsonBytes)
			return err
		})
	}
	fmt.Printf("Seeded %v fake accounts...\n", total)
}

func (bc *BoltClient) QueryAccount(accountId string) (model.Account, error) {
	//Allocate an empty Account instance
	account := model.Account{}

	err := bc.boldDB.View(func(tx *bolt.Tx) error {
		//Read the bucket from the DB
		b := tx.Bucket([]byte("AccountBucket"))

		//Read the value identified by our accountId
		accountBytes := b.Get([]byte(accountId))
		if accountBytes == nil {
			return fmt.Errorf("No account found for " + accountId)
		}
		//Unmarshal the returned bytes into the account struct created at the top of the function
		json.Unmarshal(accountBytes, &account)

		return nil
	})
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}

// 测试是否能够连接上BoltDb
func (bc *BoltClient) Check() bool {
	return bc.boldDB != nil
}

func (bc *BoltClient) Close() {
	bc.Close()
}
