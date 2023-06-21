package models

import (
	"database/sql"
	"diablo_iv_tool/pkg/database"
	"fmt"
	"github.com/Masterminds/squirrel"
)

type RoleModel struct {
	Id      string `json:"id"`
	GuildId string `json:"guild_id"`
	RoleId  string `json:"role_id"`
	Name    string `json:"name"`
}

func (role *RoleModel) CreateTable() error {
	_, err := database.Database.Exec(`
		CREATE TABLE IF NOT EXISTS roles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			guild_id TEXT NOT NULL,
			role_id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL,
			FOREIGN KEY(guild_id) REFERENCES guilds(guild_id)
		);
	`)
	if err != nil {
		return err
	}

	return nil
}

func appendRolesToList(list []RoleModel, rows *sql.Rows) ([]RoleModel, error) {
	var role RoleModel
	err := rows.Scan(
		&role.Id,
		&role.GuildId,
		&role.RoleId,
		&role.Name,
	)
	if err != nil {
		return nil, err
	}

	list = append(list, role)
	return list, nil
}

func getRolesListFromRows(rows *sql.Rows) ([]RoleModel, error) {
	var list []RoleModel
	var err error
	list, err = appendRolesToList(list, rows)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		list, err = appendRolesToList(list, rows)
		if err != nil {
			return nil, err
		}
	}
	return list, nil
}

func (role *RoleModel) getRoleByQuery(query squirrel.SelectBuilder) error {
	rows, err := database.SelectHelper(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Scan(
		&role.Id,
		&role.GuildId,
		&role.RoleId,
		&role.Name,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetRolesByGuildId(guildId string) ([]RoleModel, error) {
	rows, err := database.SelectHelper(
		squirrel.
			Select("*").
			From("roles").
			Where(squirrel.Eq{"guild_id": guildId}),
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles, err := getRolesListFromRows(rows)
	return roles, err
}

func (role *RoleModel) GetRoleById(id string) (*RoleModel, error) {
	return role, role.getRoleByQuery(
		squirrel.
			Select("*").
			From("roles").
			Where(squirrel.Eq{"id": id}),
	)
}

func (role *RoleModel) GetRoleByRoleId(roleId string) (*RoleModel, error) {
	return role, role.getRoleByQuery(
		squirrel.
			Select("*").
			From("roles").
			Where(squirrel.Eq{"role_id": roleId}),
	)
}

func (role *RoleModel) CreateRole() error {
	result, err := squirrel.
		Insert("roles").
		Columns(
			"guild_id",
			"role_id",
			"name",
		).
		Values(
			role.GuildId,
			role.RoleId,
			role.Name,
		).
		RunWith(database.Database).Exec()
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	role.Id = fmt.Sprintf("%d", id)
	return nil
}

func (role *RoleModel) DeleteRole() error {
	_, err := squirrel.
		Delete("roles").
		Where(squirrel.Eq{"id": role.Id}).
		RunWith(database.Database).Exec()
	return err
}

func (role *RoleModel) UpdateRole() error {
	roleQuery := squirrel.Update("roles")

	if role.RoleId != "" {
		roleQuery = roleQuery.Set("role_id", role.RoleId)
	}

	if role.Name != "" {
		roleQuery = roleQuery.Set("name", role.Name)
	}

	if role.GuildId != "" {
		roleQuery = roleQuery.Set("channel", role.GuildId)
	}

	_, err := roleQuery.
		Where(squirrel.Eq{"id": role.Id}).
		RunWith(database.Database).Exec()
	return err
}
