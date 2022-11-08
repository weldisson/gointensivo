package database

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/weldisson/gointensivo/internal/order/entity"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

type OrderRepositoryTestSuite struct {
	suite.Suite
	Db *sql.DB
}

// executado sempre que é chamado uma suite de testes
func (suite *OrderRepositoryTestSuite) SetupSuite() {
	// abre conexão com banco
	db, err := sql.Open("sqlite3", ":memory:")
	suite.NoError(err)
	// cria tabela orders
	_, err = db.Exec("CREATE TABLE orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id))")
	suite.Db = db
}

func (suite *OrderRepositoryTestSuite) TearDownTest() {
	// fecha conexão com banco
	suite.Db.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuite))
}

func (suite *OrderRepositoryTestSuite) TestGivenAnOrder_WhenSave_ThenShouldSaveOrder() {
	// nova ordem
	order, err := entity.NewOrder("123", 10.0, 2.0)
	// não pode ter erros
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult entity.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price from orders where id = ?", order.ID).
		Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
