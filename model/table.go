package model

import (
	"database/sql"
	"time"

	"github.com/AdrianLungu/decimal"
	"github.com/pkg/errors"
)

// AccessToken represents store_api.access_token
type AccessToken struct {
	AccountID   int64     // account_id
	Token       string    // token
	IsValid     bool      // is_valid
	GeneratedAt time.Time // generated_at
}

// Create inserts the AccessToken to the database.
func (r *AccessToken) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO access_token (account_id, token, is_valid, generated_at) VALUES ($1, $2, $3, $4)`,
		&r.AccountID, &r.Token, &r.IsValid, &r.GeneratedAt)
	if err != nil {
		return errors.Wrap(err, "failed to insert access_token")
	}
	return nil
}

// GetAccessTokenByPk select the AccessToken from the database.
func GetAccessTokenByPk(db Queryer, pk0 int64) (*AccessToken, error) {
	var r AccessToken
	err := db.QueryRow(
		`SELECT account_id, token, is_valid, generated_at FROM access_token WHERE account_id = $1`,
		pk0).Scan(&r.AccountID, &r.Token, &r.IsValid, &r.GeneratedAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select access_token")
	}
	return &r, nil
}

// Item represents store_api.item
type Item struct {
	ID          int64           // id
	Name        string          // name
	Price       decimal.Decimal // price
	Description sql.NullString  // description
}

// Create inserts the Item to the database.
func (r *Item) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO item (name, price, description) VALUES ($1, $2, $3) RETURNING id`,
		&r.Name, &r.Price, &r.Description).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert item")
	}
	return nil
}

// GetItemByPk select the Item from the database.
func GetItemByPk(db Queryer, pk0 int64) (*Item, error) {
	var r Item
	err := db.QueryRow(
		`SELECT id, name, price, description FROM item WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Name, &r.Price, &r.Description)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select item")
	}
	return &r, nil
}

// Sale represents store_api.sale
type Sale struct {
	ID         int64           // id
	AccountID  int64           // account_id
	ItemID     int64           // item_id
	PaidAmount decimal.Decimal // paid_amount
	SoldAt     time.Time       // sold_at
}

// Create inserts the Sale to the database.
func (r *Sale) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO sale (account_id, item_id, paid_amount, sold_at) VALUES ($1, $2, $3, $4) RETURNING id`,
		&r.AccountID, &r.ItemID, &r.PaidAmount, &r.SoldAt).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert sale")
	}
	return nil
}

// GetSaleByPk select the Sale from the database.
func GetSaleByPk(db Queryer, pk0 int64) (*Sale, error) {
	var r Sale
	err := db.QueryRow(
		`SELECT id, account_id, item_id, paid_amount, sold_at FROM sale WHERE id = $1`,
		pk0).Scan(&r.ID, &r.AccountID, &r.ItemID, &r.PaidAmount, &r.SoldAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select sale")
	}
	return &r, nil
}

// UserAccount represents store_api.user_account
type UserAccount struct {
	ID           int64     // id
	Email        string    // email
	Gender       string    // gender
	Birthday     time.Time // birthday
	Password     string    // password
	RegisteredAt time.Time // registered_at
}

// Create inserts the UserAccount to the database.
func (r *UserAccount) Create(db Queryer) error {
	err := db.QueryRow(
		`INSERT INTO user_account (email, gender, birthday, password, registered_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		&r.Email, &r.Gender, &r.Birthday, &r.Password, &r.RegisteredAt).Scan(&r.ID)
	if err != nil {
		return errors.Wrap(err, "failed to insert user_account")
	}
	return nil
}

// GetUserAccountByPk select the UserAccount from the database.
func GetUserAccountByPk(db Queryer, pk0 int64) (*UserAccount, error) {
	var r UserAccount
	err := db.QueryRow(
		`SELECT id, email, gender, birthday, password, registered_at FROM user_account WHERE id = $1`,
		pk0).Scan(&r.ID, &r.Email, &r.Gender, &r.Birthday, &r.Password, &r.RegisteredAt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select user_account")
	}
	return &r, nil
}

// Username represents store_api.username
type Username struct {
	AccountID   int64  // account_id
	LowerName   string // lower_name
	DisplayName string // display_name
}

// Create inserts the Username to the database.
func (r *Username) Create(db Queryer) error {
	_, err := db.Exec(
		`INSERT INTO username (account_id, lower_name, display_name) VALUES ($1, $2, $3)`,
		&r.AccountID, &r.LowerName, &r.DisplayName)
	if err != nil {
		return errors.Wrap(err, "failed to insert username")
	}
	return nil
}

// GetUsernameByPk select the Username from the database.
func GetUsernameByPk(db Queryer, pk0 int64) (*Username, error) {
	var r Username
	err := db.QueryRow(
		`SELECT account_id, lower_name, display_name FROM username WHERE account_id = $1`,
		pk0).Scan(&r.AccountID, &r.LowerName, &r.DisplayName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select username")
	}
	return &r, nil
}
