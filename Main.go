package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
)

var server string
var port string
var user string
var password string
var database string

type StructureTable struct {
	Company                            string `json:"Company"`
	RC                                 string `json:"RC"`
	OC                                 string `json:"OC"`
	Channel                            string `json:"Channel"`
	ProductCode                        string `json:"ProductCode"`
	DrCr                               string `json:"DrCr"`
	GLAccountNo                        string `json:"GLAccountNo"`
	GLAccountName                      string `json:"GLAccountName"`
	Activity                           string `json:"Activity"`
	Tax                                string `json:"Tax"`
	InterCo                            string `json:"InterCo"`
	Future1                            string `json:"Future1"`
	Future2                            string `json:"Future2"`
	Currency                           string `json:"Currency"`
	AmountEnteredDebit                 string `json:"AmountEnteredDebit"`
	AmountEnteredCredit                string `json:"AmountEnteredCredit"`
	LoanAccountNumber                  string `json:"LoanAccountNumber"`
	GroupReferenceNumber               string `json:"GroupReferenceNumber"`
	OriginalTransactionReferenceNumber string `json:"OriginalTransactionReferenceNumber"`
	TransactionReferenceNumber         string `json:"TransactionReferenceNumber"`
	GLGroupCodeCoA                     string `json:"GLGroupCodeCoA"`
	TransactionPostingDate             string `json:"TransactionPostingDate"`
	EffectiveDate                      string `json:"EffectiveDate"`
	JournalEntryDescription            string `json:"JournalEntryDescription"`
}

func main() {
	StartTime := time.Now()
	log.Println("Start at " + (StartTime.String()))
	godotenv.Load()
	server = os.Getenv("EnvServer")
	port = os.Getenv("EnvUser")
	user = os.Getenv("EnvPassword")
	password = os.Getenv("EnvPort")
	database = os.Getenv("EnvDatabase")
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", server, user, password, port, database)
	startRowW := os.Getenv("startRow")
	startRow,errs := strconv.Atoi(startRowW)
	if errs != nil {
		log.Fatal("convert start row error", errs.Error())
	}
	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}

	InsertToSql(os.Getenv("PathFile"), startRow, db)
	log.Println("End at" + (time.Now().String()))
	executionMin := time.Now().Sub(StartTime)
	fmt.Println("execution times ", executionMin)

	defer db.Close()
}

func InsertToSql(path string, startRow int, db *sql.DB) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	//fileContent := ""

	for i, line := range lines[startRow:] {
		columns := StructureTable{}
		columns.Company = strings.TrimSpace(string(line[0:10]))
		columns.RC = strings.TrimSpace(string(line[10:20]))
		columns.OC = strings.TrimSpace(string(line[20:30]))
		columns.Channel = strings.TrimSpace(string(line[30:40]))
		columns.ProductCode = strings.TrimSpace(string(line[40:55]))
		columns.DrCr = strings.TrimSpace(string(line[55:65]))
		columns.GLAccountNo = strings.TrimSpace(string(line[65:85]))
		columns.GLAccountName = strings.TrimSpace(string(line[85:140]))
		columns.Activity = strings.TrimSpace(string(line[140:150]))
		columns.Tax = strings.TrimSpace(string(line[150:160]))
		columns.InterCo = strings.TrimSpace(string(line[160:170]))
		columns.Future1 = strings.TrimSpace(string(line[170:180]))
		columns.Future2 = strings.TrimSpace(string(line[180:190]))
		columns.Currency = strings.TrimSpace(string(line[180:199]))
		columns.AmountEnteredDebit = strings.ReplaceAll(strings.TrimSpace(string(line[199:224])),",","")
		columns.AmountEnteredCredit = strings.ReplaceAll(strings.TrimSpace(string(line[224:249])),",","")
		columns.LoanAccountNumber = strings.TrimSpace(string(line[249:275]))
		columns.GroupReferenceNumber = strings.TrimSpace(string(line[275:305]))
		columns.OriginalTransactionReferenceNumber = strings.TrimSpace(string(line[305:345]))
		columns.TransactionReferenceNumber = strings.TrimSpace(string(line[345:375]))
		columns.GLGroupCodeCoA = strings.TrimSpace(string(line[375:395]))
		columns.TransactionPostingDate = strings.TrimSpace(string(line[395:425]))
		columns.EffectiveDate = strings.TrimSpace(string(line[425:445]))
		columns.JournalEntryDescription = strings.TrimSpace(string(line[445:(len(line) - 1)]))

		if InsertToTable(db, columns) != true {
			log.Println("error at row" + strconv.Itoa(i))
		}
		//log.Println(fileContent)
	}

	log.Println("Inserted successfully ", len(lines), " recods")

	return nil
	//return os.WriteFile(path, []byte(fileContent), 0644)
}

func InsertToTable(db *sql.DB, columns StructureTable) bool {
	ctx := context.Background()

	err := db.PingContext(ctx)
	if err != nil {
		log.Fatal("Ping database failed:", err.Error())
		return false
	}

	_, sqlerr := db.ExecContext(ctx, `insert into `+os.Getenv("TabletoInsert")+` values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, columns.Company, columns.RC, columns.OC, columns.Channel, columns.ProductCode, columns.DrCr, columns.GLAccountNo, columns.GLAccountName, columns.Activity, columns.Tax, columns.InterCo, columns.Future1, columns.Future2, columns.Currency, columns.AmountEnteredDebit, columns.AmountEnteredCredit, columns.LoanAccountNumber, columns.GroupReferenceNumber, columns.OriginalTransactionReferenceNumber, columns.TransactionReferenceNumber, columns.GLGroupCodeCoA, columns.TransactionPostingDate, columns.EffectiveDate, columns.JournalEntryDescription)
	if sqlerr != nil {
		log.Fatal(sqlerr.Error())
		return false
	}
	return true
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
