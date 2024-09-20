## Skyforge-YAGPDB
- [Overview](#overview)
- [Contributors & Translators](#contributors--translators)
- [Skyforge Command](#skyforge-command)
- [PvP Command](#pvp-command)
- [Disclaimer](#disclaimer)
- [License](#license)

### Translations
- [Skyforge Command (DE)](https://github.com/Samillion/skyforge-yagpdb/tree/main/translations#skyforge-de-command)
- [Skyforge Command (FR)](https://github.com/Samillion/skyforge-yagpdb/tree/main/translations#skyforge-fr-command)

## Overview
I've created custom commands that will provide useful information on Discord about the MMO game [Skyforge](https://sf.my.games/en) using YAGPDB Bot [Custom Commands](https://docs.yagpdb.xyz/commands/custom-commands) feature.

- Skyforge command outputs many useful information for players
- PvP command to display Happy Hours and next active map

To use any of the commands here, first you'll need to add [YAGPDB bot](https://yagpdb.xyz/) to your Discord server. They provide detailed documentation on how to get started with the bot [here](https://docs.yagpdb.xyz/getting-started).

> [!IMPORTANT]
> All the infromation you'd need to install and run these commands are listed below, so please make sure to read each section **fully and thoroughly** before adding any of the commands to your Discord server.

## Contributors & Translators
This project was started by Sam Sinner and Leo Scrame.

| Contributors | Translators |
| ------ | ------ |
| Laama Laamanen | DE: Leo Scrame |
| Zanzuro Mizoru | FR: Jynn Zemxil |
| Reinna Sigma | QA: Phiozof Barophilo |
| Istani Revenant | QA: Sky Darcancia |
| Nova Char | QA: Shocobo Kurochi |

## Skyforge Command
This command will output the relevant data about the MMO game Skyforge with the following options:
```
-skyforge
-skyforge weapons
-skyforge adepts
-skyforge ether
-skyforge aspects
-skyforge cog
-skyforge tips
-skyforge argents
-skyforge build
-skyforge build dps
-skyforge build support
-skyforge build tank
-skyforge build companion
-skyforge build pvp
-skyforge disclaimer
```

To install this custom command:
- Login to your YAGPDB Dashboard. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Go to "Ungrouped" tab or create a new Group if you'd like this command to be used by specific roles only.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: command
- Trigger: skyforge

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/skyforge-command.go), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![skyforge-dashboard](https://i.imgur.com/HXHfqQE.jpeg)

Click on "Save" and done! That's it. Try it on your Discord server with `-skyforge`

Here is a preview of the command `-skyforge`:

![image](https://github.com/Samillion/skyforge-yagpdb/assets/17427046/390399fa-490a-4532-9a39-deac233d4458)

## PvP Command
This command will output data relevant to PvP Happy Hours, schedule and active maps with the following options:
```
-pvp
-pvp next
```

To install this custom command:
- Login to your YAGPDB Dashboard. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Go to "Ungrouped" tab or create a new Group if you'd like this command to be used by specific roles only.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: command
- Trigger: pvp

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/pvp-command.go), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![pvp-dashboard](https://i.imgur.com/l9LilzK.jpeg)

Click on "Save" and done! That's it. Try it on your Discord server with `-pvp`

Here is a preview of the command `-pvp`:

![image](https://github.com/Samillion/skyforge-yagpdb/assets/17427046/f0d61e79-e39c-46e5-8f99-3e0b4bd07c91)

> [!IMPORTANT]
> For the PvP Command, if you are playing the game on NA server, change the region from `EU` to `NA` in the first code line.

So the result would be:
```go
{{ $region := "NA" }}
```

## Disclaimer
All the information provided by `-skyforge` and the other commands is voluntary effort, not official. It is based entirely on players testing, so result may vary depending on changes in the future. It is merely meant to guide you in the correct direction.

## License
Feel free to use, copy, modify or even take credit for the code as you please, just don't sell it or charge anyone. Let all use it as they like for free.
