# Skyforge-YAGPDB
- [Overview](#overview)
- [Skyforge Command](#skyforge-command)
- [PvP Command](#pvp-command)
- [PvP Auto Announce](#pvp-auto-announce)
- [Contributors & Translators](#contributors--translators)
- [Disclaimer](#disclaimer)
- [License](#license)
- [Contact](#contact)

### Translations
- [Skyforge Command (DE)](https://github.com/Samillion/skyforge-yagpdb/tree/main/translations#skyforge-de-command)
- [Skyforge Command (FR)](https://github.com/Samillion/skyforge-yagpdb/tree/main/translations#skyforge-fr-command)

# Overview
I've created custom commands that will provide useful information on Discord about the MMO game [Skyforge](https://sf.my.games/en) using YAGPDB Bot [Custom Commands](https://docs.yagpdb.xyz/commands/custom-commands) feature.

- Skyforge command outputs many useful information for players
- PvP command to display Happy Hours and next active map
- PvP Happy Hour auto announcement command

To use any of the commands here, first you'll need to add [YAGPDB bot](https://yagpdb.xyz/) to your Discord server. They provide detailed documentation on how to get started with the bot [here](https://docs.yagpdb.xyz/getting-started).

### Important
All the infromation you'd need to install and run these commands are listed below, but please make sure to read each section **fully and thoroughly** before adding any of the commands to your Discord server.

# Skyforge Command
This command will output the relevant data about the MMO game Skyforge with the following options:
```
-skyforge
-skyforge build
-skyforge build dps
-skyforge build support
-skyforge build tank
-skyforge build companion
-skyforge build pvp
-skyforge weapons
-skyforge adepts
-skyforge ether
-skyforge aspects
-skyforge cog
-skyforge tips
-skyforge argents
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

![skyforge-command](https://i.imgur.com/zZS1y2P.jpeg)

# PvP Command
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

![pvp-command-demo](https://i.imgur.com/RH6xfaA.jpeg)

### IMPORTANT (for PvP Command)
If you are playing the game on NA server, change the region from `EU` to `NA` in the first code line.

So the result would be:
```go
{{ $region := "NA" }}
```

# PvP Auto Announce
This command will automatically announce the active PvP maps in the channel you specify on your Discord server during happy hours.

To install this custom command:
- Login to your YAGPDB Dashboard. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: Hourly interval
- Interval: 1
- Channel: [Choose the channel you would like the bot to announce Happy Hours in]

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/pvp-announce-interval.go), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![pvp-announce-dashboard](https://i.imgur.com/eALjoFD.jpeg)

Click on "Save" and done!

Here is a preview of the PvP Auto Announce command:

![pvp-announce-demo](https://i.imgur.com/s9MASbH.jpeg)

### IMPORTANT (for PvP Auto Announce)
If you would like the bot to announce it for NA region times, in the first line of code change `EU` to `NA`.

So the result would be:
```go
{{ $region := "NA" }}
```

### Also
[Only once] You ***have to*** save or run this command at the beginning of any hour of the day, otherwise the announcements won't be accurate later on.

This is a limitation of YAGPDB not having the option to control when exactly you can run the command. So for example, if you save it on `4:35:20 PM`, the command will check every hour on `X:35:20` and that will mean Happy Hour announcements will be late as well.

So at any hour of the day, just save or run the command at `HH:00:00` so that it would run at the start of every hour, as displayed here:

![interval-display-zero](https://i.imgur.com/OcRsrFu.jpeg)

## Contributors & Translators
This project was started by Sam Sinner and Leo Scrame.

**Contributors:**
- Laama Laamanen
- Zanzuro Mizoru
- Reinna Sigma
- Istani Revenant
- Nova Char

**Translators:**
- DE: Leo Scrame
- FR: Jynn Zemxil
- QA: Phiozof Barophilo
- QA: Sky Darcancia
- QA: Shocobo Kurochi

# Disclaimer
All the information provided by `-skyforge` and the other commands is voluntary effort, not official. It is based entirely on players testing, so result may vary depending on changes in the future. It is merely meant to guide you in the correct direction.

# License
Feel free to use, copy, modify or even take credit for the code as you please, just don't sell it or charge anyone. Let all use it as they like for free.

# Contact
If there are any issues, either open a ticket in this repository or contact me on Skyforge's official Discord server: **[EG]Sam Sinner**.
