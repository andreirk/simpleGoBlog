package dbrepo

import (
	"context"
	"errors"
	"time"
	"web3/models"

	"golang.org/x/crypto/bcrypt"
)

// Add functions we need here for accessing
// the database (They can be used by all
// handlers)

// Function used to add new posts
// Create a receiver m for model
// which links to the Postgres repository
func (m *postgresDBRepo) InsertPost(newPost models.Post) error {
	// If this transaction doesn't occur in
	// 5 seconds cancel
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Pass values to query using
	// placeholders
	query := `INSERT INTO posts(title, content, user_id) VALUES($1, $2, $3)`

	// Execute the query
	_, err := m.DB.ExecContext(ctx, query,
		newPost.Title,
		newPost.Content,
		newPost.UserID)
	if err != nil {
		return err
	}

	return nil
}

// 27 Gets user data by id
func (m *postgresDBRepo) GetUserByID(id int) (models.User, error) {

	// Cancel after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT name, email, password, acct_created, last_login, user_type, id FROM users WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, query, id)

	// Get user data and store in model
	var u models.User
	err := row.Scan(
		&u.Name,
		&u.Email,
		&u.Password,
		&u.AcctCreated,
		&u.LastLogin,
		&u.UserType,
		&u.ID,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

// 27 Used to modify a user
func (m *postgresDBRepo) UpdateUser(u models.User) error {
	// Cancel after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `UPDATE users SET name=$1, email=$2, last_login=$3, user_type=$4`

	_, err := m.DB.ExecContext(ctx, query,
		u.Name,
		u.Email,
		time.Now(),
		u.UserType)

	if err != nil {
		return err
	}
	return nil
}

// 27 Authenticate the user login data
func (m *postgresDBRepo) AuthenticateUser(email, testPassword string) (int, string, error) {
	// Cancel after 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// If correct login we store id
	var id int
	var hashedPW string

	query := `SELECT id, password FROM users WHERE email = $1`

	row := m.DB.QueryRowContext(ctx, query, email)

	// Put results in variables
	err := row.Scan(&id, &hashedPW)

	if err != nil {
		return id, "", err
	}

	// Compare password provided to the
	// stored password
	// (Cast to slice of bytes)
	err = bcrypt.CompareHashAndPassword([]byte(hashedPW), []byte(testPassword))

	if err == bcrypt.ErrMismatchedHashAndPassword {
		// If passwords don't match
		return 0, "", errors.New("password is incorrect")
	} else if err != nil {
		// If another error occurred
		return 0, "", err
	}

	// Correct email and pw were entered
	return id, hashedPW, nil
}
