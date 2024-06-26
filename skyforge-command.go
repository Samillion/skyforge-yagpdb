{{ $options := sdict 
	"weapons" (sdict 
		"title" "Weapon Recommendations"
		"description" "Listed weapons are considered (players opinion) to be priority.\n\nThat is why not all classes are listed. You should still prioritize which weapon to get first.\n\n**Lightbinder:**\n- Telurgan Rod (Support, PvP)\n\n**Soundweaver:**\n- Khella (Support, Comp, PvP)\n\n**Grovewalker:**\n- Sigil (DPS, Tank)\n- Endurance Weapon (Tank)\n\n**Paladin:**\n- Lament and Sorrow (Tank, DPS, PvP)\n\n**Necromancer:**\n- Morana (DPS)\n- Morodea (DPS, PvP)\n\n**Berserker:**\n- Iron Heart (DPS, PvP)\n\n**Revenant:**\n- Baron (DPS)\n- Ner'gal (PvP)"
	)
	"ether" (sdict 
		"title" "Ether (Terra)"
		"description" "It is the currency you need to unlock Terra Symbols for each class (Scarlet, Gold, Emerald). To enter a Terra Symbol mission you need 1000 Ether.\n\n**Sources:** (Ether Flower)\n- Terra (Region)\n- Heart of the City (Group)\n- Toxic Wasteland (Group)\n- Ancient Trails (Group)\n- Terrain Rift (Squad)"
	)
	"adepts" (sdict
		"title" "Adepts (Faith)"
		"description" "- Never cancel a journey to collect rewards.\n- Add Charisma & Constitution relics for max faith return.\n- Send Adepts on journey always.\n- Only 8 adepts are allowed.\n- Always keep adepts with higher stats.\n\n**Notes:**\n- Having more than 8 adepts reduces faith gain by 50%.\n- Last 30% of journey brings a high return of faith.\n- Adept stats will increase over time."
	)
	"tips" (sdict 
		"title" "General Tips"
		"description" "**Which class to unlock first?**\nPlayers recommend Grovewalker, Firestarter or Revenant, unlock with Sparks of Transformation.\n\nThey provide decent damage and give survivability.\n\n**Cathedral and Tower of Knowledge:**\nUsually recommended to not waste credits on Cathedral while leveling, because you'll need credits for many things (ie: gear) and you will get items to improve your cathedral rank as rewards.\n\nTower of Knowledge, upgrade it as you can, just don't use many Stims.\n\n**Aelerium:**\nIt is the currency to unlock A9 talents. Can be farmed in Thea (Region), for example by killing Kairax and doing region events.\n\n**Halls of Trophies Priority:**\n- Personal Medic\n- Revive Master (and Thea-1)\n- Treasure Hunter\n- Asterius's Glyph (KoE Farm)\n\n**Important:** Always be a teamplayer and a good sport, adapt to group needs."
	)
	"aspects" (sdict 
		"title" "Aspect Forms"
		"description" "Aspects: Tenacity, Rage, Mercy, Magic, Vengeance.\n\nAt first you have 3 aspects unlocked (Tenacity, Rage, Mercy). To unlock the remaining Aspects you need to clear Champion Avatar twice.\n\n**Aspect Priority:**\nEach Aspect takes time, to fully develop all 5 you need 3.1 million Cognition.\n\nFor that reason, players recommend to prioritize.\n\n**First Priority: Mercy**\nThis is probably the only agreed upon fact when developing Aspects, to work on Mercy first. The gold node brings huge advantages, most important: Elixirs. They allow you to use your Aspect Form for free.\n\n**After Mercy:**\nThere are theories on what to prioritize next, it will mostly depend on your gameplay. We highly recommend checking each Aspect gold node to help you decide."
	)
	"cog" (sdict 
		"title" "Cognition"
		"description" "It is the currency used to develop your Aspects.\n\n**Sources:**\n- Directive rewards.\n- Finishing Avatar missions, Distortions.\n- Killing invaded region bosses (Event).\n\n**Notes:**\n- Fully developing each Aspect (Gold Node), will increase the amount of Cognition you earn.\n- Having Seal of Deeds active in your Cathedral increases the amount of Cognition you earn from killing invaded region bosses."
	)
	"argents" (sdict 
		"title" "Argents"
		"description" "It is the premium currency used in Skyforge to buy legendary weapons, styles, mounts and so on.\n\n**How to get Argents:**\n- Being top 50% in daily operations. (Resets every 2 weeks)\n- Reward for finishing the gold directive of Training Avatar.\n- Participating in Inghar.\n- Buying Argents packs for money."
	)
	"disclaimer" (sdict 
		"title" "Disclaimer"
		"description" "All information provided by `-skyforge` is voluntary effort, not official. It is based on players testing, so result may vary depending on changes in the future.\n\nIt is meant to guide you in the correct direction.\n\nThis project was started by Sam Sinner and Leo Scrame.\n\nProject details on [Github](https://github.com/Samillion/skyforge-yagpdb)."
	)
}}

{{ $roles := sdict 
	"dps" (sdict 
		"title" "DPS Build"
		"description" "DPS build by **Leo Scrame**\n\n**Note:**\nThe first two symbol sets below are based on protection, whether you want it to save you from a one shot or have protection from DoT.\n\nThe last symbol set is effective against mobs.\n\n**Recommended Artifact:**\n- Black Flame\n- Vulnerability Detector\n\n**Adept Relics:**\nIt is recommended to balance Main Damage and Critical Damage relics or go with Main Damage ones.\n\n**Recommended Divine Weapon:**\n- Nerion's Hammer\n- Cyrus's Blades\n- Integrator's Transformer"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920977294670057542/full-dps-build-en.png")
	)
	"support" (sdict 
		"title" "Support Build"
		"description" "Support Build by **Sam Sinner**, **Leo Scrame** and **Laama Laamanen**.\n\n**Note:**\nIgnore the Honor stat in the gear.\n\n**Recommended Artifact:**\n- Wings of Skies (Might buffs)\n- Overload Module (with Support Mode) (Optional)\n\n**Recommended Divine Weapon:**\n- Machavann's Guard (Might buff)\n- Herida's Wings"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920977603387588638/full-support-build-en.png")
	)
	"tank" (sdict 
		"title" "Tank Build"
		"description" "Tank build by **Leo Scrame**.\n\n**Note:**\nMany Artifacts work well with a Tank build, so we decided to list all the useful ones in order.\n\n**Recommended Artifact:**\n- Aegis of the Ocean\n- Overload Module (with Implacability)\n- Viper Wristlet\n- Impenetrable Shield\n\n**Recommended Divine Weapon:**\n- Protheus's Trident\n- Laertes's Sword"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/925371668409499698/full-tank-build-en.png")
	)
	"comp" (sdict 
		"title" "Companion Build"
		"description" "Companion build by **Sam Sinner**, **Leo Scrame** and **Zanzuro Mizoru**.\n\n**Note:**\nThis build is perfect with Soundweaver. However, you can also use it with other classes such as Lightbinder.\n\nYou can also replace the Bracelet with Zeal.\n\n**Recommended Artifact:**\n- Overload Module\n- Hyper Engine (Soundweaver)\n\n**Recommended Divine Weapon:**\n- Cyrus's Blades"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920976802749505536/full-companion-build-en.png")
	)
	"pvp" (sdict 
		"title" "PvP Build"
		"description" "PvP build by **Zanzuro Mizoru**\n\n**Note:**\nFor additional defense, you can replace the Sapphire with Beacon or to decrease opponents defense, replace Ruby with Emitter.\n\n**Recommended Artifact:**\n- Glyph of Despair\n- Viper Wristlet\n- Trewang's Idol\n\n**Adept Relics:**\nMany PvP players recommend to use Defense relics."
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920977445329448960/full-pvp-build-en.png")
	)
}}

{{ $category := cslice "build" "weapons" "adepts" "ether" "aspects" "cog" "tips" "argents" "disclaimer"}}
{{ $builds := cslice "dps" "support" "tank" "companion" "comp" "pvp" }}

{{ $catList := "" }}
{{ range $category }}{{- $catList = print $catList "- " (title .) "\n" -}}{{- end }}
{{ $main := sdict "title" "Options" "description" $catList }}

{{ $buildList := "" }}
{{ range $builds }}{{- $buildList = print $buildList "- " (title .) "\n" -}}{{- end }}
{{ $build := sdict "title" "Build Options" "description" $buildList }}

{{ $embed := $main }}
{{ $args := .CmdArgs }}
{{ $gold := 16771846 }}

{{ if (ge (len $args) 1) }}
	{{ $arg := index .CmdArgs 0 | lower }}
	{{ if in $category $arg }}
		{{ if and (ne $arg "build") ($options.HasKey $arg) }}
			{{ $embed = $options.Get $arg }}
			{{ $embed.Set "color" $gold }}
		{{ else if eq $arg "build" }}
			{{ $embed = $build }}
			{{ if ge (len $args) 2 }}
				{{ $role := index .CmdArgs 1 | lower }}
				{{ $role = reReplace "companion" $role "comp" }}
				{{ if $roles.HasKey $role }}
					{{ $embed = $roles.Get $role }}
					{{ $embed.Set "color" $gold }}
				{{ end }}
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}

{{ sendMessage nil (cembed $embed) }}
