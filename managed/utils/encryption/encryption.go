// Copyright (C) 2023 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package encryption contains functions to encrypt/decrypt items or DB.
package encryption

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"os"
	"slices"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// DefaultEncryptionKeyPath contains default PMM encryption key path.
const DefaultEncryptionKeyPath = "/srv/pmm-encryption.key"

var (
	// ErrEncryptionNotInitialized is error in case of encryption is not initialized.
	ErrEncryptionNotInitialized = errors.New("encryption is not initialized")
	// DefaultEncryption is the default implementation of encryption.
	DefaultEncryption = New(DefaultEncryptionKeyPath)
)

// New create encryption, if key on path doesnt exists will be generated.
func New(keyPath string) *Encryption {
	e := &Encryption{}
	e.Path = keyPath

	bytes, err := os.ReadFile(e.Path)
	switch {
	case os.IsNotExist(err):
		err = e.generateKey()
		if err != nil {
			logrus.Errorf("Encryption: %v", err)
		}
	case err != nil:
		logrus.Errorf("Encryption: %v", err)
	default:
		e.Key = string(bytes)
	}

	primitive, err := e.getPrimitive()
	if err != nil {
		logrus.Errorf("Encryption: %v", err)
	}
	e.Primitive = primitive

	return e
}

// Encrypt is wrapper around DefaultEncryption.Encrypt.
func Encrypt(secret string) (string, error) {
	return DefaultEncryption.Encrypt(secret)
}

// Encrypt returns input string encrypted.
func (e *Encryption) Encrypt(secret string) (string, error) {
	if e == nil || e.Primitive == nil {
		return secret, ErrEncryptionNotInitialized
	}

	cipherText, err := e.Primitive.Encrypt([]byte(secret), []byte(""))
	if err != nil {
		return secret, err
	}

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

// EncryptItems is wrapper around DefaultEncryption.EncryptItems.
func EncryptItems(ctx context.Context, db *sql.DB, tables []Table) error {
	return DefaultEncryption.EncryptItems(ctx, db, tables)
}

// EncryptItems will encrypt all columns provided in DB connection.
func (e *Encryption) EncryptItems(ctx context.Context, db *sql.DB, tables []Table) error {
	if len(tables) == 0 {
		return errors.New("target tables/columns not defined")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	for _, table := range tables {
		res, err := table.Read(tx)
		if err != nil {
			return err
		}

		for k, v := range res.SetValues {
			for i, val := range v {
				var encrypted any
				var err error
				switch table.Columns[i].CustomHandler {
				case nil:
					encrypted, err = encryptColumnStringHandler(e, val)
				default:
					encrypted, err = table.Columns[i].CustomHandler(e, val)
				}

				if err != nil {
					return err
				}
				res.SetValues[k][i] = encrypted
			}
			data := slices.Concat([]any{}, v)
			data = slices.Concat(data, res.WhereValues[k])
			_, err := tx.Exec(res.Query, data...)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Decrypt is wrapper around DefaultEncryption.Decrypt.
func Decrypt(cipherText string) (string, error) {
	return DefaultEncryption.Decrypt(cipherText)
}

// Decrypt returns input string decrypted.
func (e *Encryption) Decrypt(cipherText string) (string, error) {
	if e == nil || e.Primitive == nil {
		return cipherText, ErrEncryptionNotInitialized
	}
	if cipherText == "" {
		return cipherText, nil
	}
	decoded, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return cipherText, fmt.Errorf("value %s is probably not encrypted, error: %w", cipherText, err)
	}
	secret, err := e.Primitive.Decrypt(decoded, []byte(""))
	if err != nil {
		return cipherText, err
	}

	return string(secret), nil
}

// DecryptItems is wrapper around DefaultEncryption.DecryptItems.
func DecryptItems(ctx context.Context, db *sql.DB, tables []Table) error {
	return DefaultEncryption.DecryptItems(ctx, db, tables)
}

// DecryptItems will decrypt all columns provided in DB connection.
func (e *Encryption) DecryptItems(ctx context.Context, db *sql.DB, tables []Table) error {
	if len(tables) == 0 {
		return errors.New("target tables/columns not defined")
	}

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback() //nolint:errcheck

	for _, table := range tables {
		res, err := table.Read(tx)
		if err != nil {
			return err
		}

		for k, v := range res.SetValues {
			for i, val := range v {
				var decrypted any
				var err error
				switch table.Columns[i].CustomHandler {
				case nil:
					decrypted, err = decryptColumnStringHandler(e, val)
				default:
					decrypted, err = table.Columns[i].CustomHandler(e, val)
				}

				if err != nil {
					return err
				}
				res.SetValues[k][i] = decrypted
			}
			data := slices.Concat([]any{}, v)
			data = slices.Concat(data, res.WhereValues[k])
			_, err := tx.Exec(res.Query, data...)
			if err != nil {
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
