## Translations
- [Overview](#overview)
- [Skyforge-DE](#skyforge-de-command)
- [Skyforge-FR](#skyforge-fr-command)

## Overview
Translations for the `-skyforge` English command in DE (German) and FR (French) languages.

If you already have the English `-skyforge` command added, we suggest to add the new commands as `-sfde` for German and `-sffr` for French, which is all explained below in detail.

This was made possible thanks to the efforts of the following awesome people:
- DE: Leo Scrame
- FR: Jynn Zemxil
- QA: Phiozof Barophilo
- QA: Sky Darcancia

## Skyforge-DE Command
Dieser Befehl gibt relevante Informationen über das MMO Spiel "Skyforge" durch die folgenden Optionen wieder:

```
-skyforge
-skyforge Build
-skyforge Build Dps
-skyforge Build Support
-skyforge Build Tank
-skyforge Build Gefährte
-skyforge Build Pvp
-skyforge Waffen
-skyforge Adepten
-skyforge Ether
-skyforge Aspekte
-skyforge Wahrnehmung
-skyforge Tipps
-skyforge Argents
-skyforge Disclaimer
```

Um diesen Custom Befehl zu installieren folgen Sie diesen Schritten:
- Loggen Sie sich in Ihr YAGPDB Dashboard. ([Hier]( https://yagpdb.xyz/manage ))
- Navigieren Sie zu Core -> Custom Commands.
- Wenn Sie den Befehl nur für bestimmte Rollen nutzbar machen wollen, erstellen Sie eine neue Gruppe. Andernfalls bewegen Sie sich zum "Ungrouped" Reiter.
- Klicken Sie auf "Create a new Custom Command".

Wählen Sie die folgenden Optionen:
- Trigger type: command
- Trigger: skyforge

Hier können Sie einen anderen Trigger wählen, wenn auf Ihrem Server der Bot anders angesprochen werden soll. zB.: `sfde` -> `-sfde Ether`

Gehen Sie nun [hierhin](https://github.com/Samillion/skyforge-yagpdb/blob/main/translations/skyforge-DE.go), klicken Sie auf "Raw" und kopieren den Code in die "Response"-Zeile, wie es hier gezeigt wird:

![skyforge-dashboard](https://i.imgur.com/TAQs7wI.jpeg)

Speichern Sie nun mit dem "Save"-Knopf, um die Einbindung abzuschließen. Probieren Sie es nun auf Ihrem Discord Server indem Sie `-skyforge` eingeben.

Hier ist eine Vorschau auf den `-skyforge` Befehl:

![skyforge-command](https://i.imgur.com/aTEjD0M.jpeg)

## Skyforge-FR Command [Incomplete: Work in progress]
Cette commande affiche les informations pertinentes à propos du MMO Skyforge avec les options suivantes

```
-skyforge
-skyforge build
-skyforge build dps
-skyforge build support
-skyforge build tank
-skyforge build companion
-skyforge build pvp
-skyforge armes
-skyforge adeptes
-skyforge ether
-skyforge aspects
-skyforge cog
-skyforge astuces
-skyforge argents
```

Pour installer les commandes personnalisées:
- Connectez-vous au Tableau de Bord de YAGPDB ([Ici](https://yagpdb.xyz/manage))
- Rendez-vous dans l'onglet "Core" -> "Custom Commands"
- Recherchez l'onglet "Ungrouped" ou créez un nouveau groupe si vous voulez que cette commande ne soit utilisée que par des rôles spécifiques
- Cliquez sur "Create a new Custom Command"

Remplissez les options suivantes:
- Trigger type: Command
- Trigger: skyforge

Vous pouvez changer l'option "Trigger" pour qu'elle soit plus spécifique dans votre langue. Par exemple pour le français : `sffr` -> `sffr ether`.

Ensuite rendez-vous [ici](https://github.com/Samillion/skyforge-yagpdb/blob/main/translations/skyforge-FR.go), cliquez sur "Raw" pour copier tout le code qui s'y trouve et collez-le dans le champ "Response", comme montré ici:

![skyforge-dashboard](https://i.imgur.com/rWZNAUp.jpeg)

Cliquez sur "Save" et c'est terminé! Essayez-le sur votre serveur Discord avec `-skyforge`.

Voici une prévisualisation de la commande `-skyforge`:

![skyforge-command](https://i.imgur.com/j7oHCHO.jpeg)
