{{$gold := 16771846}}
{{$red := 16711680}}
{{$mainOpts := cslice "build" "waffen" "adepten" "ether" "aspekte" "wahrnehmung" "tipps" "argents"}}
{{$buildOpts := cslice "dps" "support" "tank" "gefährte" "comp" "pvp"}}
{{$embedMain := cembed
	"title" "Optionen"
	"description" (joinStr "\n" $mainOpts.StringSlice)
	"color" $red
}}
{{$embedBuild := cembed 
	"title" "Sub-Optionen"
	"description" (joinStr "\n" $buildOpts.StringSlice)
	"color" $red
}}
{{if (eq (len .Args) 2 3)}}
	{{$skyOpt := $mainOpts}}
	{{$skyArg := index .CmdArgs 0}}
	{{$skyArg := (lower $skyArg)}}
	{{if in $skyOpt $skyArg}}
		{{if eq $skyArg "waffen"}}
			{{$embed := cembed
				"title" "Waffenempfehlungen"
				"description" "Listet Waffen auf, die nach Spielerempfehlung notwendig und daher zu empfehlen sind.\n\nAus diesem Grund fehlen ein paar Klassen. Die Reihenfolge sollte man jedoch selbst entscheiden.\n\n**Lichthüter:**\n- Telurgans Stab (Support)\n\n**Klangschmied:**\n- Khella (Support, Comp, PvP)\n\n**Hainwandler:**\n- Sigil (DPS, Tank)\n- Widerstandskraft-Waffe (Tank)\n\n**Paladin:**\n- Wehklage und Leid (Tank, DPS, PvP)\n\n**Nekromant:**\n- Morana (DPS)\n- Morodea (DPS, PvP)\n\n**Berserker:**\n- Eisenherz (Mobs DPS, PvP)\n\n**Revenant:**\n- Baron (DPS)\n- Ner'gal (PvP)"
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "ether"}}
			{{$embed := cembed
				"title" "Ether (Terra)"
				"description" "Ether ist die Währung, die zum Freischalten von Terra-Symbolen jeder Klasse benötigt wird (Scarlet, Gold, Emerald). Um einer Terra-Symbol-Mission beizutreten wird 1000 Ether benötigt.\n\n**Etherquellen:** (Etherblumen)\n- Terra (Region)\n- Herz der Stadt (Team)\n- Giftiges Ödland (Team)\n- Alte Pfade (Team)\n- Schlucht von Terra (Einheit)"
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "adepten"}}
			{{$embed := cembed
				"title" "Adepten (Glauben)"
				"description" "- Beende niemals die Reise eines Adepten vorzeitig.\n- Nutze Charisma- & Konstitutionsrelliquien um den Glauben zu maximieren.\n- Sende deine Adepten immer auf Reisen.\n- 8 Adepten sind maximal erlaubt, für zusätzlichen Glauben schicke überflüssige Adepten weg.\n- Behalte immer die Adepten mit den höchsten Eigenschaften\n\n**Notizen:**\n- Mehr als 8 Adepten reduzieren das Glaubenseinkommen um 50%.\n- Die letzten 30% einer Reise bringen am meisten Glauben."
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "tipps"}}
			{{$embed := cembed
				"title" "Generelle Tipps"
				"description" "**Welche Klasse sollte man zuerst freischalten?**\nSpieler empfehlen Hainwandler, Pyromant oder Revenant für's Leveln.\n\nSie bieten vernünftigen Schaden und Überlebensskills.\n\n**Kathedrale und Turm des Wissens:**\nNormalerweise wird empfohlen Credits nicht für die Kathedrale zu verschwenden, solange man levelt.\n\nDen Turm des Wissens kannst du immer aufwerten, nur benutze dafür nicht zu viele Reizmittel.\n\n**Aelerium:**\nDiese Währung wird zum freischalten von A9-Talenten benutzt. Sie kann in Thea (Region) gefarmt werden, zum Beispiel durch's besiegen von Kairax oder andere Events.\n\n**Halle der Trophäen - Priorität:**\n- Leibarzt\n- Meister wiederbeleben (und Thea-1)\n- Schatzsucher\n- Artefakt (z.B. Schwachstellen-Scanner)\n\n**Wichtig:** Spiele immer für's Team, passe dich daran an, was das Team braucht."
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "aspekte"}}
			{{$embed := cembed
				"title" "Aspektformen"
				"description" "Aspekte: Standhaftigkeit, Wut, Gnade, Magie, Rache.\n\nZuerst hast du 3 Aspekte freigeschaltet (Standhaftigkeit, Wut, Gnade). Um die übrigen freizuschalten musst du zwei Championavatare besiegen.\n\n**Aspektpriorität:**\nJeder Aspekt braucht seine Zeit. Um alle 5 Aspekte vollständig durchzuskillen braucht man etwa 3,1M Wahrnehmung (Cog).\n\n**Oberste Priorität: Gnade**\nBei dieser Prio sind sich alle einig, man sollte zuerst Gnade freischalten. Der goldene 'Node' gibt einen großen Vorteil durch Elixiere, mit denen man die Aspektform gratis einnehmen kann.\n\n**Nach Gnade:**\nEs gibt viele verschiedene Ansichten, was als nächstes priorisiert werden sollte, dabei hängt es vor allem vom eigenen Spielstil ab."
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "wahrnehmung"}}
			{{$embed := cembed
				"title" "Wahrnehmung"
				"description" "Wahrnehmung ist die Währung die man zum Entwickeln der Aspekte des 'Elder God' benötigt.\n\n**Wie erhält man Wahrnehmung:**\n- Anordnungsbelohnungen.\n- Erfolgreiches Beenden von Trainings- und Champion-Avatar, Verzerrungen.\n- Besiegen von Bossen in überfallenen Regionen (Event).\n\n**Notizen:**\n- vollständig einen Aspekt durchzuskillen erhöht die Wahrnehmung, die man erhält.\n- Das 'Siegel der Taten' in der Kathedrale erhöht ebenfalls die Wahrnehmung die man vom Besiegen von Bossen in den invadierten Regionen erhält."
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "argents"}}
			{{$embed := cembed
				"title" "Argents"
				"description" "Argents ist die Premiumwährung von Skyforge, die man zum Kaufen von legendären Waffen, Reittieren usw. benötigt.\n\n**Wie kommt man an Argents:**\n- In den top 50% der täglichen Operation sein (zurückgesetzt alle 2 Wochen).\n- Belohnung für das Beenden der Gold-Anordnung des Training Avatars.\n- Teilnahme in Inghar.\n- Kaufen von Argentspaketen für Echtgeld."
				"color" $gold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "build"}}
			{{$builds := $buildOpts}}
			{{if eq (len .Args) 2}}
				{{sendMessage nil $embedBuild}}
			{{else}}
				{{$role := index .CmdArgs 1}}
				{{$role := lower $role}}
				{{if in $builds $role}}
					{{if eq $role "dps"}}
						{{$embed := cembed
							"title" "DPS Build"
							"description" "DPS Build von **Leo Scrame**\n\n**Notizen:**\nDie ersten zwei Symbol-Sets variieren abhängig davon, ob man wenige starke, oder viele schwache Treffer abbekommt.\n\nDas letzte Symbol-Set ist effektiv gegen schwächere Gegner.\n\n**Empfohlene Artefakte:**\n- Schwarze Flamme\n- Schwachstellen-Scanner\n\n**Adeptenreliquien:**\nEs wird empfohlen entweder Hauptschaden und kritischen Schaden auszubalancieren, oder sich auf Hauptschaden zu fokussieren.\n\n**Empfohlene göttliche Waffen:**\n- Cyrus' Klingen\n- Transformator des Integrators\n- Hammer von Nerion"
							"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920986307629572106/full-dps-build-de.png")
							"color" $gold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "support"}}
						{{$embed := cembed
							"title" "Support Build"
							"description" "Support Build von **Sam Sinner**, **Leo Scrame** und **Laama Laamanen**.\n\n**Notiz:**\nIgnoriere den Ehre-Stat der Ausrüstung.\n\n**Empfohlene Artefakte:**\n- Schwingen des Himmels (Machtbuffs)\n- Modul 'Überlastung' (mit Unterstützungsmodus)\n\n**Empfohlene göttliche Waffen:**\n- Machavanns Schutz (Machtbuffs)\n- Heridas Flügel"
							"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920986880391131206/full-support-build-de.png")
							"color" $gold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "tank"}}
						{{$embed := cembed
							"title" "Tank Build"
							"description" "Tank Build von **Leo Scrame**.\n\n**Empfohlene Artefakte:**\n- Aegis des Ozeans\n- Modul 'Überlastung' (mit Unerbittlichkeit der Aspekte)\n- Vipernarmband\n- Undurchdringlicher Schild\n\n**empfohlene göttliche Waffen:**\n- Protheus' Dreizack\n- Schwert von Laertes"
							"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/925371728207695902/full-tank-build-de.png")
							"color" $gold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "gefährte" "comp"}}
						{{$embed := cembed
							"title" "Gefährten Build"
							"description" "Gefährten-Build von **Sam Sinner**, **Leo Scrame** und **Zanzuro Mizoru**.\n\n**Notiz:**\nDieser Build empfiehlt sich für den Klangschmied, man kann ihn aber auch auf anderen Klassen spielen.\n\nMan kann das Armband durch Inbrust der Aspekte ersetzen.\n\n**Empfohlene Artefakte:**\n- Modul 'Überlastung'\n- Hypermotor (Klangschmied)\n\n**Empfohlene göttliche Waffen:**\n- Cyrus' Klingen\n- Transformator des Integrators"
							"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920986112732856320/full-companion-build-de.png")
							"color" $gold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "pvp"}}
						{{$embed := cembed
							"title" "PvP Build"
							"description" "PvP Build von **Zanzuro Mizoru**\n\n**Notiz:**\nFür zusätzliche Verteidigung kann man den Saphir durch Licht der Fürbitte ersetzen, oder um die Verteidigung der Gegner zu reduzieren den Rubin durch den Strahler des Untergangs.\n\n**Empfohlene Artefakte:**\n- Schrifttafel der Verzweiflung\n- Vipernarmband\n- Trewangs Idol\n\n**Adeptenreliquien:**\nDie meisten PvP-Spieler empfehlen Verteidigung."
							"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920986546604240916/full-pvp-build-de.png")
							"color" $gold
						}}
						{{sendMessage nil $embed}}
					{{end}}
				{{else}}
					{{sendMessage nil $embedBuild}}
				{{end}}
			{{end}}
		{{end}}
	{{else}}
		{{sendMessage nil $embedMain}}
	{{end}}
{{else}}
	{{sendMessage nil $embedMain}}
{{end}}
