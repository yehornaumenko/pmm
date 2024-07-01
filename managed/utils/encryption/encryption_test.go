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

package encryption

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncryption(t *testing.T) {
	secret := "password1"

	cipherText, err := Encrypt(secret)
	require.NoError(t, err)
	require.NotEmpty(t, cipherText)
	decryptedSecret, err := Decrypt(cipherText)
	require.NoError(t, err)
	require.Equal(t, secret, decryptedSecret)

	c := &DatabaseConnection{
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "postgres",
		Password: "",
		EncryptedItems: []EncryptedDatabase{
			{
				Database: "pmm-managed",
				Tables: []EncryptedTable{
					{
						Table:          "agents",
						Identificators: []string{"agent_id"},
						Columns: []EncryptedColumn{
							{Column: "username", Handler: func() {}},
							{Column: "password", Handler: func() {}},
						},
					},
				},
			},
		},
	}

	ctx := context.Background()
	require.NoError(t, EncryptDB(ctx, c))
	require.NoError(t, DecryptDB(ctx, c))
}
