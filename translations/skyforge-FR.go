{{$colorGold := 16771846}}
{{$colorRed := 16711680}}
{{$embedMain := cembed 
"title" "Options"
"description" "build \narmes \nadeptes \nether \naspects \ncog \nastuces \nargents"
"color" $colorRed}}
{{if (eq (len .Args) 2 3)}}
{{$skyOpt := (cslice "build" "armes" "adeptes" "ether" "aspects" "cog" "astuces" "argents")}}
{{$skyArg := index .CmdArgs 0}}
{{$skyArg := (lower $skyArg)}}
{{if in $skyOpt $skyArg}}
{{if eq $skyArg "armes"}}
{{$embed := cembed 
"title" "Recommandations pour les armes"
"description" "Les armes listées ci-dessous sont considérées (selon les joueurs) comme prioritaires, essentielles.\n\nC'est pourquoi toutes les classes ne sont pas listées. Réfléchissez toujours à quelle arme vous voulez en premier.\n\n**Lumancien:**\n- Baguette de Telurgan (Support, PvP)\n\n**Décibel:**\n- Khella (Support, Comp, PvP)\n\n**Saronide:**\n- Sigil (DPS, Tank)\n- Arme d'endurance (Tank)\n\n**Paladin:**\n- Chagrin et Lamentation (Tank, DPS, PvP)\n\n**Nécromancien:**\n- Morana (DPS)\n- Morodea (DPS, PvP)\n\n**Berserker:**\n- Coeur de Fer (DPS, PvP)\n\n**Revenant:**\n- Baron (DPS)\n- Ner'gal (PvP)"
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "ether"}}
{{$embed := cembed 
"title" "Éther (Terra)"
"description" "L'Éther est la ressource utilisée pour débloquer vos Symboles de classe sur Terra (Écarlate, Or, Émeraude). Pour débloquer la mission d'un symbole de terra, vous avez besoin de 1000 d'Éther.\n\n**Source d'Éther** (Fleur d'Éther)\n- Terra (Région)\n- Coeur de la Ville (Groupe)\n- Terres Gastes Toxiques (Groupe)\n- Pistes Anciennes (Groupe)\n- Faille Terranne (Unité)"
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "adeptes"}}
{{$embed := cembed 
"title" "Adeptes (Foi)"
"description" "- Ne jamais annuler un voyage pour récupérer les récompenses.\n- Ajouter des reliques avec du Charisme & de la Constitution pour récupérer plus de foi.\n- Toujours envoyer ses adeptes en voyage.\n- Seulement 8 adeptes sont autorisés.\n- Toujours garder les adeptes avec les stats les plus hautes.\n\n**Notes:**\n- Avoir plus de 8 adeptes réduit la foi gagnée de 50%.\n- Lors des derniers 25% de son voyage, un adepte ramène plus de foi.\n- Plus vous gardez vos adeptes longtemps, plus leurs stats vont évoluer."
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "astuces"}}
{{$embed := cembed 
"title" "Petites Astuces"
"description" "**Quelle classe débloquer en premier?**\nLes joueurs recommandent le Saronide, le Feurioso ou le Revenant pendant que vous gagnez en rang, déblocables avec les Étincelles de Transformation.\n\nCes classes vous permettent de survivre tout en infligeant de bons dégâts.\n\n**Cathédrale et Tour de la Connaissane:**\nLes joueurs vous recommandent généralement de ne pas gaspiller vos crédits pendant que vous montez en niveau ; vous aurez besoin de ces crédits pour d'autres choses (ie: l'équipement) et vous allez obtenir des fidèles pour améliorer le rang de votre Cathédrale en récompense.\n\nAugmentez votre Tour de la Connaissance avec les ressources que vous avez, mais n'utilisez pas beaucoup de Stimulants..\n\n**Aelerium:**\nIl s'agit de la ressource utilisée pour débloquer les talents Aelerium-9 (A9). Ils peuvent être récupérés à Théa (Région), en tuant Kairax ou en participant aux évènements de la région.\n\n**Quel bonus débloquer en priorité dans la salle des trophées:**\n- Médecin Personnel\n- Ramener le maître à la vie (et Théa-1)\n- Chasseur de trésors\n- Glyphe d'Asterius (KoE Farm)\n\n**Important:** Il est primordial d'avoir toujours l'esprit d'équipe, d'être bon joueur et de s'adapter aux besoins du groupe."
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "aspects"}}
{{$embed := cembed 
"title" "Formes d'Aspects"
"description" "Aspects: Ténacité, Rage, Pitié, Magie, Vengeance.\n\nAu début, vous avez 3 aspects disponibles (Ténacité, Rage, Pitié). Pour débloquer les deux autres aspects, vous devez vaincre deux Avatars Champions.\n\n**Priorité des Aspects:**\nChaque aspect prend du temps ; pour terminer les 5 aspects, vous avez besoin d'un total de 3.1 million de Cognition.\n\nPour cette raison, les joueurs recommandent de compléter vos aspects dans cet ordre:\n\n**Premier Aspect : la Pitié**\nCeci est le point sur lequel tout le monde s'entend concernant les aspects : complétez tout d'abord la Pitié. Le nodule d'or apporte d'énormes avantages, dont les élixirs, qui vous permettent d'utiliser gratuitement votre forme d'Aspect.\n\n**Après la Pitié:**\nIl y a beaucoup de débats sur l'aspect à prioriser après la Pitié, cela va en grande partie dépendre de votre gameplay. On vous recommande fortement de vérifier chaque nodule d'or pour vous aider à vous décider."
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "cog"}}
{{$embed := cembed 
"title" "Cognition"
"description" "La Cognition est la ressource utilisée pour développer vos Aspects.\n\n**Comment gagner de la Cognition:**\n- En récompense de directive.\n- En terminant l'Avatar d'entraînement, l'Avatar Champion et les distortions.\n- En tuant les boss des régions envahies (évènemement).\n\n**Notes:**\n- En terminant chaque Aspect (Nodule d'or), vous gagnerez plus de Cognition.\n- En activant le Sceau des Exploits dans votre Cathédrale, vous gagnerez plus de Cognition en tuant les boss des régions envahies."
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "argents"}}
{{$embed := cembed 
"title" "Argents"
"description" "Les Argents sont la monnaie payante utilisée pour acheter des armes légendaires, des tenues, des montures et autres\n\n**Comment obtenir des Argents:**\n- Se positionner dans le top 50% des opérations quotidiennes (se réinitialisent toutes les 2 semaines)\n- En terminant la directive de l'Avatar d'Entraînement.\n- En participant aux Inghar.\n- En l'achetant avec de la monnaie réelle."
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $skyArg "build"}}
{{$allowedRoles := (cslice "dps" "support" "tank" "companion" "pvp")}}
{{if eq (len .Args) 2}}
{{$embed := cembed 
"title" "Sub-options"
"description" "dps \nsupport \ntank \ncompanion \npvp"
"color" $colorRed}}
{{sendMessage nil $embed}}
{{else}}
{{$role := index .CmdArgs 1}}
{{$role := lower $role}}
{{if in $allowedRoles $role}}
{{if eq $role "dps"}}
{{$embed := cembed 
"title" "DPS Build"
"description" "DPS build by **Leo Scrame**\n\n**Note:**\nThe first two symbol sets below are based on protection, whether you want it to save you from a one shot or have protection from DoT.\n\nThe last symbol set is effective against mobs.\n\n**Recommended Artifact:**\n- Black Flame\n- Vulnerability Detector\n\n**Adept Relics:**\nIt is recommended to balance Main Damage and Critical Damage relics or go with Main Damage ones.\n\n**Recommended Divine Weapon:**\n- Nerion's Hammer\n- Cyrus's Blades\n- Integrator's Transformer"
"image" (sdict "url" "https://i.imgur.com/fw4W64p.png")
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $role "support"}}
{{ $embed := cembed 
"title" "Support Build"
"description" "Support Build by **Sam Sinner**, **Leo Scrame** and **Laama Laamanen**.\n\n**Note:**\nIgnore the Honor stat in the gear.\n\n**Recommended Artifact:**\n- Wings of Skies (Might buffs)\n- Overload Module (with Support Mode) (Optional)\n\n**Recommended Divine Weapon:**\n- Machavann's Guard (Might buff)\n- Herida's Wings"
"image" (sdict "url" "https://i.imgur.com/Gj2lc0b.png")
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $role "tank"}}
{{$embed := cembed 
"title" "Tank Build"
"description" "Tank build by **Leo Scrame**.\n\n**Note:**\nMany Artifacts work well with a Tank build, so we decided to list all the useful ones in order.\n\n**Recommended Artifact:**\n- Aegis of the Ocean\n- Overload Module (with Implacability)\n- Viper Wristlet\n- Impenetrable Shield\n\n**Recommended Divine Weapon:**\n- Protheus's Trident\n- Laertes's Sword"
"image" (sdict "url" "https://i.imgur.com/emugFG7.png")
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $role "companion"}}
{{$embed := cembed 
"title" "Companion Build"
"description" "Companion build by **Sam Sinner**, **Leo Scrame** and **Zanzuro Mizoru**.\n\n**Note:**\nThis build is perfect with Soundweaver. However, you can also use it with other classes such as Lightbinder.\n\nYou can also replace the Bracelet with Zeal.\n\n**Recommended Artifact:**\n- Overload Module\n- Hyper Engine (Soundweaver)\n\n**Recommended Divine Weapon:**\n- Cyrus's Blades"
"image" (sdict "url" "https://i.imgur.com/T0VhzCR.png")
"color" $colorGold}}
{{sendMessage nil $embed}}
{{else if eq $role "pvp"}}
{{$embed := cembed 
"title" "PvP Build"
"description" "PvP build by **Zanzuro Mizoru**\n\n**Note:**\nFor additional defense, you can replace the Sapphire with Beacon or to decrease opponents defense, replace Ruby with Emitter.\n\n**Recommended Artifact:**\n- Glyph of Despair\n- Viper Wristlet\n- Trewang's Idol\n\n**Adept Relics:**\nMany PvP players recommend to use Defense relics."
"image" (sdict "url" "https://i.imgur.com/n6yYkYa.png")
"color" $colorGold}}
{{sendMessage nil $embed}}
{{end}}
{{else}}
{{$embed := cembed 
"title" "Invalid"
"description" "Allowed: dps, support, tank, companion, pvp"
"color" $colorRed}}
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
