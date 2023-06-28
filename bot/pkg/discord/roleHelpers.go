package discord

import (
	"bot/pkg/models"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

const (
	worldTierColor = 2003199
	classColor     = 2142890
	alertColor     = 14423100
)

var (
	worldTierRoles = []string{
		"World Tier 1",
		"World Tier 2",
		"World Tier 3",
		"World Tier 4",
	}

	classRoles = []string{
		"Barbarian",
		"Druid",
		"Sorcerer",
		"Rogue",
		"Necromancer",
	}

	alertRoles = []string{
		"Morning",
		"Day",
		"Afternoon",
		"Evening",
	}
)

func generateRoleParams(name string, color int) discordgo.RoleParams {
	falsePointer := false
	truePointer := true
	return discordgo.RoleParams{
		Name:        name,
		Color:       &color,
		Hoist:       &falsePointer,
		Permissions: nil,
		Mentionable: &truePointer,
	}
}

func createRole(name string, color int, guildId string, s *discordgo.Session, guildRoles []*discordgo.Role) (*discordgo.Role, error) {
	params := generateRoleParams(name, color)
	if len(guildRoles) > 0 {
		for _, role := range guildRoles {
			if role.Name == name {
				log.Printf("Role: %v already exists for guild: %v", role.Name, guildId)
				return role, nil
			}
		}
	}
	time.Sleep(300 * time.Millisecond)
	return s.GuildRoleCreate(guildId, &params)
}

func createRolesFromList(roles []string, color int, guildId string, s *discordgo.Session, guildRoles []*discordgo.Role) error {
	for _, role := range roles {
		createdRole, err := createRole(role, color, guildId, s, guildRoles)
		if err != nil {
			return err
		}

		log.Printf("Created role: %v for guild: %v", createdRole.Name, guildId)

		newRole := models.RoleModel{
			GuildId: guildId,
			RoleId:  createdRole.ID,
			Name:    createdRole.Name,
		}
		err = newRole.CreateRole()
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateRoles(s *discordgo.Session, guildId string) error {
	guild := models.GuildModel{}
	_, err := guild.GetGuildByGuildId(guildId)
	if err != nil {
		return err
	}

	guildRoles, err := s.GuildRoles(guildId)
	if err != nil {
		return err
	}

	err = createRolesFromList(worldTierRoles, worldTierColor, guildId, s, guildRoles)
	if err != nil {
		return err
	}

	err = createRolesFromList(classRoles, classColor, guildId, s, guildRoles)
	if err != nil {
		return err
	}

	err = createRolesFromList(alertRoles, alertColor, guildId, s, guildRoles)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRoles(s *discordgo.Session, guildId string) error {
	roles, err := models.GetRolesByGuildId(guildId)
	if err != nil {
		return err
	}

	for _, role := range roles {
		err = s.GuildRoleDelete(guildId, role.RoleId)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Deleted role: %v for guild: %v", role.Name, guildId)
	}

	return nil
}

func SetRole(roleName, guildId, userId string, s *discordgo.Session) error {
	roles, err := models.GetRolesByGuildId(guildId)
	if err != nil {
		return err
	}

	for _, role := range roles {
		if role.Name != roleName {
			continue
		}
		err = s.GuildMemberRoleAdd(guildId, userId, role.RoleId)
		if err != nil {
			return err
		}
		break
	}

	return nil
}

func UnsetAllRoles(guildId, userId string, s *discordgo.Session) error {
	roles, err := models.GetRolesByGuildId(guildId)
	if err != nil {
		return err
	}

	for _, role := range roles {
		err = s.GuildMemberRoleRemove(guildId, userId, role.RoleId)
		if err != nil {
			return err
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}
