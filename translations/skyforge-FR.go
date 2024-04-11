{{ $options := sdict 
	"armes" (sdict 
		"title" "Recommandations pour les armes"
		"description" "Les armes listées ci-dessous sont considérées selon les joueurs comme prioritaires.\n\nC'est pourquoi toutes les classes ne sont pas listées. Réfléchissez toujours à quelle arme vous voulez en premier\n\n**Lumancien:**\n- Telurgan (Support, PvP)\n\n**Décibel:**\n- Khella (Support, Comp, PvP)\n\n**Saronide:**\n- Sigil (DPS, Tank)\n- Arme d'endurance (Tank)\n\n**Paladin:**\n- Chagrin et Lamentation (Tank, DPS, PvP)\n\n**Nécromancien:**\n- Morana (DPS)\n- Morodea (DPS, PvP)\n\n**Berserker:**\n- Coeur de Fer (DPS, PvP)\n\n**Revenant:**\n- Baron (DPS)\n- Ner'gal (PvP)"
	)
	"ether" (sdict 
		"title" "Éther (Terra)"
		"description" "L'Éther est la ressource utilisée pour débloquer vos Symboles de classe sur Terra (Écarlate, Or, Émeraude). Pour débloquer la mission d'un symbole, vous avez besoin de 1000 d'Éther.\n\n**Source d'Éther** (Fleur d'Éther)\n- Terra\n- Coeur de la Ville\n- Terres Gastes Toxiques\n- Pistes Anciennes\n- Faille Terranne"
	)
	"adeptes" (sdict
		"title" "Adeptes (Foi)"
		"description" "- Ne jamais annuler un voyage en cours\n- Ajouter des reliques avec du Charisme & de la Constitution pour avoir plus de foi\n- Toujours envoyer ses adeptes en voyage\n- Seuls 8 adeptes sont autorisés\n- Toujours garder les adeptes avec les stats les plus hautes\n\n**Notes:**\n- Avoir plus de 8 adeptes réduit la foi gagnée de 50%\n- Lors des derniers 30% de son voyage, un adepte ramène plus de foi\n- Plus vous gardez vos adeptes longtemps, plus leurs stats vont évoluer"
	)
	"astuces" (sdict 
		"title" "Astuces"
		"description" "**Quelle classe débloquer en premier?**\nLes joueurs recommandent le Saronide, le Feurioso ou le Revenant pendant que vous gagnez en rang ; ces classes octroient de la survibabilité et des dégâts\n\n**Cathédrale et Tour de la Connaissane:**\nLes joueurs vous recommandent de ne pas gaspiller vos crédits ; vous en aurez besoin pour d'autres choses (ie: l'équipement) et vous obtiendrez des fidèles en récompenses de directives pour améliorer votre Cathédrale. Augmentez votre Tour de la Connaissance quand vous le pouvez, mais n'utilisez pas de Stimulants.\n\n**Aelerium:**\nIl s'agit de la ressource utilisée pour débloquer les talents Aelerium-9 (A9). L'Aelerium peut être récupéré à Théa, en tuant Kairax ou en participant aux évènements.\n\n**Priorités dans la Salle des Trophées:**\n- Médecin Personnel\n- Ramener le maître à la vie (et Théa-1)\n- Chasseur de trésors\n- Glyphe d'Asterius (KoE Farm)\n\n**Important:** Il est primordial d'avoir l'esprit d'équipe, d'être bon joueur et de s'adapter aux besoins du groupe"
	)
	"aspects" (sdict 
		"title" "Formes d'Aspects"
		"description" "Aspects: Ténacité, Rage, Pitié, Magie, Vengeance.\n\nAu début, vous avez 3 aspects disponibles (Ténacité, Rage, Pitié). Pour débloquer les deux autres, vous devez tuer deux Avatars Champions.\n\n**Priorité des Aspects:**\nChaque aspect est long à monter ; pour tous les terminer, vous avez besoin de 3,1 million de Cognition. C'est pour ça que les joueurs recommandent de compléter vos aspects dans cet ordre:\n\n**Premier Aspect : la Pitié**\nTous les joueurs sont d'accord là-dessus. Son nodule d'or apporte d'énormes avantages, dont les élixirs qui remboursent transformation en Aspect.\n\n**Après la Pitié:**\nIl y a beaucoup de débats sur quel aspect monter ensuite, cela dépendra en grande partie de votre gameplay. On vous recommande de vérifier chaque nodule d'or pour vous aider à vous décider."
	)
	"cog" (sdict 
		"title" "Cognition"
		"description" "La Cognition est la ressource utilisée pour développer vos Aspects\n\n**Comment gagner de la Cognition:**\n- Sur vos directives.\n- En terminant l'Avatar d'entraînement, l'Avatar Champion et les distortions\n- En tuant les boss des régions envahies (évènemement)\n\n**Notes:**\n- En complétant chaque Aspect, vous gagnerez plus de Cognition\n- Activez le Sceau des Exploits dans votre Cathédrale pour gagner plus de Cognition en tuant les boss des régions envahies"
	)
	"argents" (sdict 
		"title" "Argents"
		"description" "Les Argents sont la monnaie payante utilisée pour acheter des armes légendaires, des tenues, des montures, etc\n\n**Comment en obtenir:**\n- Se positionner dans le top 50% des opérations quotidiennes (se réinitialisent toutes les 2 semaines)\n- Terminer la directive de l'Avatar d'Entraînement\n- Participer aux Inghar\n- En acheter avec de la monnaie réelle"
	)
	"disclaimer" (sdict 
		"title" "Disclaimer"
		"description" "All information provided by `-skyforge` is voluntary effort, not official. It is based on players testing, so result may vary depending on changes in the future.\n\nIt is meant to guide you in the correct direction.\n\nThis project was started by Sam Sinner and Leo Scrame.\n\nProject details on [Github](https://github.com/Samillion/skyforge-yagpdb)."
	)
}}

{{ $roles := sdict 
	"dps" (sdict 
		"title" "Build DPS"
		"description" "Build proposé par Leo Scrame\n\n**Note:**\nLes deux premiers sets de symboles montrés sous le build sont basés sur de la protection contre les 'one-shots' (1er) ou contre des dots (2ieme). Le dernier set est efficace contre les monstres.\n\n**Artefacts:**\n- Flamme Noire\n- Détecteur de Vulnérabilité\n\n**Reliques:**\nIl est recommandé d'équilibrer les reliques avec des Dégâts Principaux et des Dégâts Critiques, ou de mettre uniquement des Dégâts Principaux.\n\n**Armes Divines:**\n- Marteau de Nerion\n- Lames de Cyrus\n- Intégrateur"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920994016303120394/full-dps-build-fr.png")
	)
	"support" (sdict 
		"title" "Build Support"
		"description" "Build proposé par Sam Sinner, Leo Scrame et Laama Laamanen.\n\n**Note:**\nVous pouvez ignorer la statistique d'Honneur de cet équipement.\n\n**Artefacts:**\n- Ailes des cieux (buffs de puissance)\n- Module Surcharge (avec le Mode Soutien) (Optionnel)\n\n**Armes Divines:**\n- Machavann\n- Ailes d'Herida"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920994378300923955/full-support-build-fr.png")
	)
	"tank" (sdict 
		"title" "Build Tank"
		"description" "Build proposé par Leo Scrame\n\n**Note:**\nBeaucoup d'artefacts sont utiles pour tanker, donc nous avons décidé de lister les plus utiles.\n\n**Artefacts:**\n- Égide de l'Océan\n- Module Surcharge (avec Implacabilité)\n- Bracelet de Vipère\n- Bouclier impénétrable\n\n**Armes Divines:**\n- Trident de Prothéus\n- Épée de Laërte"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/925371802627231784/full-tank-build-fr.png")
	)
	"comp" (sdict 
		"title" "Build Compagnon"
		"description" "Build proposé par Sam Sinner, Leo Scrame et Zanzuro Mizoru.\n\n**Note:**\nCe build marche bien avec le Décibel. Mais vous pouvez l'utiliser avec d'autres classes\n\nVous pouvez aussi remplacer le Bracelet avec le Zèle.\n\n**Artefacts:**\n- Module Surcharge\n- Moteur Hyper (Décibel)\n\n**Arme Divine:**\n- Lames de Cyrus"
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920993893217099806/full-companion-build-fr.png")
	)
	"pvp" (sdict 
		"title" "Build PVP"
		"description" "Build proposé par Zanzuro Mizoru\n\n**Note:**\nPour avoir plus de défense, vous pouvez remplacer le Saphir avec la Balise de Protection. Pour diminuer la défense de vos ennemis, remplacez le Rubis avec l'Emetteur du Désespoir\n\n**Artefacts:**\n- Glyphe du Désespoir\n- Bracelet de Vipère\n- Trewang\n\n**Reliques:**\nBeaucoup de joueurs PvP utilisent des reliques de Défense."
		"image" (sdict "url" "https://cdn.discordapp.com/attachments/913303283110649916/920994248160071700/full-pvp-build-fr.png")
	)
}}

{{ $category := cslice "build" "armes" "adeptes" "ether" "aspects" "cog" "astuces" "argents" "disclaimer"}}
{{ $builds := cslice "dps" "support" "tank" "compagnon" "comp" "pvp" }}

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
				{{ $role = reReplace "compagnon" $role "comp" }}
				{{ if $roles.HasKey $role }}
					{{ $embed = $roles.Get $role }}
					{{ $embed.Set "color" $gold }}
				{{ end }}
			{{ end }}
		{{ end }}
	{{ end }}
{{ end }}

{{ sendMessage nil (cembed $embed) }}
