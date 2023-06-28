
# Diablo IV Tool

A discord bot providing time trackers and (yet to come) tools for Diablo IV.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

[Add this bot to your server](https://discord.com/api/oauth2/authorize?client_id=1120791758406701056&permissions=275146427456&scope=bot)

## Acknowledgements

Diablo IV's world bosses spawns in 4 separate timespans:
- Morning (4:30 - 6:30)
- Day (10:30 - 12:30)
- Afternoon (16:30 - 18:30)
- Evening (22:30 - 00:30)

## Features

- Upcoming boss timer
- Upcoming bosses list
- Upcoming boss notifications per timespan
- Upcoming helltide timer
- Upcoming helltides list
- More to come...

## Usage

Upon joining the server, the bot will create the following roles:
- Barbarian
- Druid
- Sorcerer
- Rogue
- Necromancer
- World Tier 1
- World Tier 2
- World Tier 3
- World Tier 4
- Morning
- Day
- Afternoon
- Evening

While most of these roles are not used by the bot (this is only cosmetic yet), 
the 4 last roles, which are corresponding to the timespans of the world bosses spawns, are used to notify the users.

The bot will automatically notify when the next world boss spawns in less than 60 minutes.
This notification includes a mention to the role corresponding to the timespan of the spawn.

Bosses spawns are announced in the channel set by the `/channel` command.

All dates and time provided by the bot are in UTC. Unless specified otherwise by the `/timezone` command.

Every command is a discord Slash Command, command sent in the chat will not work.

### Info commands

- `/boss`: Get upcoming boss timer
- `/bosses`: Get upcoming bosses list
- `/helltide`: Get upcoming helltide timer
- `/helltides`: Get upcoming helltides list

### Conf commands

- `/channel`: Set the channel to send notifications to
- `/timezone <TZ Identifier>`: Set the timezone to use for dates and times (e.g. `Europe/Paris`)

### User/Role commands

- `/alert <span>`: (morning, day, afternoon, evening) Invocating user get attributed the corresponding role to receive notifications
- `/class <role>`: (barbarian, druid, sorcerer, rogue, necromancer) Invocating user get attributed the corresponding class role
- `/wt <world tier>`: (1, 2, 3, 4) Invocating user get attributed the corresponding world tier role

## Environment Variables

To run this bot locally, you will need to add the following environment variables to your .env file

`DBIVTOOL_BOT_TOKEN`: The discord bot token

`DBIVTOOL_DB_DRIVER`: The database driver (`sqlite3`)

`DBIVTOOL_DB_DSN`: The database DSN (`file:./db.sqlite3?_foreign_keys=ON`)

`DBIVTOOL_REDIS_HOST`: The redis host (`localhost`), if using the provided redis server, this variable should be set to `redis`

`DBIVTOOL_REDIS_PORT`: The redis port (`6379`)

`DBIVTOOL_REDIS_PASSWORD`: The redis password 

`DBIVTOOL_REDIS_DB`: The redis database number (`0`)

`DBIVTOOL_VOLUME_PATH_HOST`: The host path to the persistent folder

`DBIVTOOL_VOLUME_PATH_CONTAINER`: The container path to the persistent folder

## Deployment

To deploy this project run

```bash
docker-compose up -d
```

To deploy this project with a custom redis server, run

```bash
docker-compose up -d bot fetcher
```

## Feedback

If you have any feedback, please create an issue on this repository.

## Credits

By [@millefeuille](https://www.github.com/Millefeuille42)

This bot works using the following APIs:  
- https://d4builds.gg/ 
- https://worldstone.io/
- https://helltides.com/
