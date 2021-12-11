{{$colorGold := 16771846}}
{{$colorRed := 16711680}}
{{$embedMain := cembed 
	"title" "Options"
	"description" "build \nweapons \nadepts \nether \naspects \ncog \ntips \nargents \ndisclaimer"
	"color" $colorRed
}}
{{if (eq (len .Args) 2 3)}}
	{{$skyOpt := (cslice "build" "weapons" "adepts" "ether" "aspects" "cog" "tips" "argents" "disclaimer")}}
	{{$skyArg := index .CmdArgs 0}}
	{{$skyArg := (lower $skyArg)}}
	{{if in $skyOpt $skyArg}}
		{{if eq $skyArg "weapons"}}
			{{$embed := cembed 
				"title" "Weapon Recommendations"
				"description" "Listed weapons are considered (players opinion) to be priority, the essential ones.\n\nThat is why not all classes are listed. You should still prioritize which weapon to get first.\n\n**Lightbinder:**\n- Telurgan Rod (Support, PvP)\n\n**Soundweaver:**\n- Khella (Support, Comp, PvP)\n\n**Grovewalker:**\n- Sigil (DPS, Tank)\n- Endurance Weapon (Tank)\n\n**Paladin:**\n- Lament and Sorrow (Tank, DPS, PvP)\n\n**Necromancer:**\n- Morana (DPS)\n- Morodea (DPS, PvP)\n\n**Berserker:**\n- Iron Heart (DPS, PvP)\n\n**Revenant:**\n- Baron (DPS)\n- Ner'gal (PvP)"
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "ether"}}
			{{$embed := cembed 
				"title" "Ether (Terra)"
				"description" "Ether is the currency you need to unlock Terra Symbols for each class (Scarlet, Gold, Emerald). To enter a Terra Symbol mission you need 1000 Ether.\n\n**Sources of Ether:** (Ether Flower)\n- Terra (Region)\n- Heart of the City (Group)\n- Toxic Wasteland (Group)\n- Ancient Trails (Group)\n- Terrain Rift (Squad)"
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "adepts"}}
			{{$embed := cembed 
				"title" "Adepts (Faith)"
				"description" "- Never cancel a journey to collect rewards.\n- Add Charisma & Constitution relics for max faith return.\n- Send Adepts on journey always.\n- Only 8 adepts are allowed.\n- Always keep adepts with higher stats.\n\n**Notes:**\n- Having more than 8 adepts reduces faith gain by 50%.\n- Last 25% of journey brings a high return of faith.\n- Adept stats will increase the longer you keep them."
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "tips"}}
			{{ $embed := cembed 
				"title" "General Tips"
				"description" "**Which class to unlock first?**\nPlayers recommend Grovewalker, Firestarter or Revenant while leveling, unlock with Sparks of Transformation.\n\nThey provide decent damage and give you survivability.\n\n**Cathedral and Tower of Knowledge:**\nUsually recommended to not waste credits on Cathedral while leveling, because you'll need credits for many things (ie: gear) and you will get items to improve your cathedral rank as rewards.\n\nTower of Knowledge, upgrade it as you can, just don't use many Stimulants.\n\n**Aelerium:**\nThat is the currency to unlock the A9 talents. Can be farmed in Thea (Region), for example by killing Kairax and doing region events.\n\n**Halls of Trophies Priority:**\n- Personal Medic\n- Revive Master (and Thea-1)\n- Treasure Hunter\n- Asterius's Glyph (KoE Farm)\n\n**Important:** Always be a teamplayer and a good sport, adapt to group needs."
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "aspects"}}
			{{$embed := cembed 
				"title" "Aspect Forms"
				"description" "Aspects: Tenacity, Rage, Mercy, Magic, Vengeance.\n\nAt first, you have 3 aspects unlocked (Tenacity, Rage, Mercy). To unlock the remaining Aspects you need to clear Champion Avatar twice.\n\n**Aspect Priority:**\nEach Aspect takes time, to fully develop all 5 Aspects you need 3.1 million Cognition total.\n\nFor that reason, players recommend to prioritize.\n\n**First Priority: Mercy**\nThis is probably the only agreed upon fact when developing Aspects, to work on Mercy first. The gold node brings huge advantages, most important: Elixirs. They allow you to use your Aspect Form for free.\n\n**After Mercy:**\nThere are many theories on what to prioritize next, it will mostly depend on your gameplay. We highly recommend checking each Aspect gold node to help you decide."
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "cog"}}
			{{$embed := cembed 
				"title" "Cognition"
				"description" "Cognition is the currency used to develop your Aspects.\n\n**How to earn Cognition:**\n- Directive rewards.\n- Finishing Training and Champion Avatar, Distortions.\n- Killing invaded region bosses (Event).\n- Participating in Pantheon Wars battles.\n\n**Notes:**\n- Fully developing each Aspect (Gold Node), will increase the amount of Cognition you earn.\n- Having 'Seal of Deeds' active in your Cathedral increases the amount of Cognition you earn from killing invaded region bosses."
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "argents"}}
			{{$embed := cembed 
				"title" "Argents"
				"description" "Argents is the premium currency used in Skyforge to buy legendary weapons, styles, mounts and so on.\n\n**How to get Argents:**\n- Being top 50% in daily operations. (Resets every 2 weeks)\n- Reward for finishing the gold directive of Training Avatar.\n- Participating in Inghar.\n- Participating in Pantheon Wars.\n- Buying Argents packs for money."
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "disclaimer"}}
			{{$embed := cembed 
				"title" "Disclaimer"
				"description" "All information provided by `-skyforge` is voluntary effort, not official. It is based on players testing, so result may vary depending on changes in the future.\n\nIt is meant to guide you in the correct direction.\n\nThis project was started by Sam Sinner and Leo Scrame.\n\n**Contributors:**\n- Laama Laamanen\n- Zanzuro Mizoru\n- Reinna Sigma\n- Istani Revenant\n- Nova Char\n\n**Translators:**\n- DE: Leo Scrame\n- FR: Jynn Zemxil"
				"color" $colorGold
			}}
			{{sendMessage nil $embed}}
		{{else if eq $skyArg "build"}}
			{{$allowedRoles := (cslice "dps" "support" "tank" "companion" "pvp")}}
			{{if eq (len .Args) 2}}
				{{$embed := cembed 
					"title" "Sub-options"
					"description" "dps \nsupport \ntank \ncompanion \npvp"
					"color" $colorRed
				}}
				{{sendMessage nil $embed}}
			{{else}}
				{{$role := index .CmdArgs 1}}
				{{$role := lower $role}}
				{{if in $allowedRoles $role}}
					{{if eq $role "dps"}}
						{{$embed := cembed 
							"title" "DPS Build"
							"description" "DPS build by **Leo Scrame**\n\n**Note:**\nThe first two symbol sets below are based on protection, whether you want it to save you from a one shot or have protection from DoT.\n\nThe last symbol set is effective against mobs.\n\n**Recommended Artifact:**\n- Black Flame\n- Vulnerability Detector\n\n**Adept Relics:**\nIt is recommended to balance Main Damage and Critical Damage relics or go with Main Damage ones.\n\n**Recommended Divine Weapon:**\n- Nerion's Hammer\n- Cyrus's Blades\n- Integrator's Transformer"
							"image" (sdict "url" "https://i.imgur.com/vLdUlRc.png")
							"color" $colorGold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "support"}}
						{{ $embed := cembed 
							"title" "Support Build"
							"description" "Support Build by **Sam Sinner**, **Leo Scrame** and **Laama Laamanen**.\n\n**Note:**\nIgnore the Honor stat in the gear.\n\n**Recommended Artifact:**\n- Wings of Skies (Might buffs)\n- Overload Module (with Support Mode) (Optional)\n\n**Recommended Divine Weapon:**\n- Machavann's Guard (Might buff)\n- Herida's Wings"
							"image" (sdict "url" "https://i.imgur.com/H5V9B6k.png")
							"color" $colorGold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "tank"}}
						{{$embed := cembed 
							"title" "Tank Build"
							"description" "Tank build by **Leo Scrame**.\n\n**Note:**\nMany Artifacts work well with a Tank build, so we decided to list all the useful ones in order.\n\n**Recommended Artifact:**\n- Aegis of the Ocean\n- Overload Module (with Implacability)\n- Viper Wristlet\n- Impenetrable Shield\n\n**Recommended Divine Weapon:**\n- Protheus's Trident\n- Laertes's Sword"
							"image" (sdict "url" "https://i.imgur.com/7pJg7GL.png")
							"color" $colorGold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "companion"}}
						{{$embed := cembed 
							"title" "Companion Build"
							"description" "Companion build by **Sam Sinner**, **Leo Scrame** and **Zanzuro Mizoru**.\n\n**Note:**\nThis build is perfect with Soundweaver. However, you can also use it with other classes such as Lightbinder.\n\nYou can also replace the Bracelet with Zeal.\n\n**Recommended Artifact:**\n- Overload Module\n- Hyper Engine (Soundweaver)\n\n**Recommended Divine Weapon:**\n- Cyrus's Blades"
							"image" (sdict "url" "https://i.imgur.com/46LViCe.png")
							"color" $colorGold
						}}
						{{sendMessage nil $embed}}
					{{else if eq $role "pvp"}}
						{{$embed := cembed 
							"title" "PvP Build"
							"description" "PvP build by **Zanzuro Mizoru**\n\n**Note:**\nFor additional defense, you can replace the Sapphire with Beacon or to decrease opponents defense, replace Ruby with Emitter.\n\n**Recommended Artifact:**\n- Glyph of Despair\n- Viper Wristlet\n- Trewang's Idol\n\n**Adept Relics:**\nMany PvP players recommend to use Defense relics."
							"image" (sdict "url" "https://i.imgur.com/Gqgq8oj.png")
							"color" $colorGold
						}}
						{{sendMessage nil $embed}}
					{{end}}
					{{else}}
					{{$embed := cembed 
						"title" "Invalid"
						"description" "Allowed: dps, support, tank, companion, pvp"
						"color" $colorRed
					}}
					{{sendMessage nil $embed}}
					{{end}}
				{{end}}
		{{end}}
	{{else}}
		{{sendMessage nil $embedMain}}
	{{end}}
{{else}}
{{sendMessage nil $embedMain}}
{{end}}
