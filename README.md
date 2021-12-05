# Skyforge-YAGPDB
I've created custom commands that will provide useful information on Discord about the MMO game [Skyforge](https://sf.my.games/en) using YAGPDB Bot [Custom Commands](https://docs.yagpdb.xyz/commands/custom-commands) feature.

- Skyforge command outputs many useful information for players
- PvP command to display Happy Hours and next active map
- PvP Happy Hour auto announcement command

To use any of the commands here, first you'll need to add [YAGPDB bot](https://yagpdb.xyz/) to your Discord server. They provide detailed documentation on how to get started with the bot [here](https://docs.yagpdb.xyz/getting-started).

# Skyforge Command
To install this custom command:
- Login to your YAGPDB Dashbord. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Go to "Ungrouped" tab or create a new Group if you'd like this command to be used by specific roles only.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: command
- Trigger: skyforge

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/skyforge-command.go.tmpl), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![skyforge-dashboard](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/skyforge-dashboard-command.jpg?token=AEE6UZV356RSNMGIGOSHYM3BVQBCM)

Click on "Save" and done! That's it. Try it on your Discord server with `-skyforge`

Here is a preview of the command `-skyforge`:

![skyforge-command](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/skyforge-command-demo.jpg?token=AEE6UZXLSJLUUC7IL576SPLBVQCT6)

# PvP Command
To install this custom command:
- Login to your YAGPDB Dashbord. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Go to "Ungrouped" tab or create a new Group if you'd like this command to be used by specific roles only.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: command
- Trigger: pvp

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/pvp-command.go.tmpl), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![pvp-dashboard](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/pvp-dashboard-command.jpg?token=AEE6UZVOBIQ4PUAE337KH4TBVQBR6)

Click on "Save" and done! That's it. Try it on your Discord server with `-pvp`

Here is a preview of the command `-pvp`:

![pvp-command-demo](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/pvp-command-demo.jpg?token=AEE6UZRWDCGZZ64IWRPOWIDBVQCAA)

### IMPORTANT (for PvP Command)
If you are playing the game on NA server, change the region from `EU` to `NA` in the first code line.

So the result would be:
```
{{ $region := "NA" }}
```

# PvP Auto Announce
To install this custom command:
- Login to your YAGPDB Dashbord. ([Here](https://yagpdb.xyz/manage))
- Navigate to Core -> Custom Commands.
- Click on "Create a new Custom Command".

Choose the following options:
- Trigger type: Hourly interval
- Interval: 1
- Channel: [Choose the channel you would like the bot to announce Happy Hours in]

Then go [here](https://github.com/Samillion/skyforge-yagpdb/blob/main/pvp-announce-interval.go.tmpl), click on "Raw" then copy all the code there and paste it in the "Response" field, as shown here:

![pvp-announce-dashboard](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/pvp-announce-dashboard.jpg?token=AEE6UZSIVLNJ4UMXXC3HHLDBVQDVQ)

Click on "Save" and done!

Here is a preview of the PvP Auto Announce command:

![pvp-announce-demo](https://raw.githubusercontent.com/Samillion/skyforge-yagpdb/main/previews/pvp-autoannounce-demo.jpg?token=AEE6UZTWFSGCMNXLQXPGEFDBVQGBY)

### IMPORTANT (for PvP Auto Announce)
If you would like the bot to announce it for NA region times, in the first lines of the code change `Europe/Berlin` to `US/Pacific`.

So the result would be:
```
{{ $tz := "US/Pacific" }}
```

[Only once] You ***have to*** save or run this command at the beginning of any hour of the day, otherwise the announcements won't be accurate later on.

This is a limitatation of YAGPDB not having the option to control when exactly you can run the command. So for example, if you save it on `4:35:20 PM`, the script will check every hour on `X:35:20` and that will mean Happy Hour announcements will be late as well.

So at any hour of the day, just save or run the command at `HH:00:00` so that it would run at the start of every hour.
