<h2>C'est quoi `bufio` ? </h2> 

`bufio` est un package qui fournit des `Readers` et `Writers` bufferisés.

Ca signifie qu'il lit/écrit en mémoire temporaire avant d'envoyer ou afficher 

<h2>Pourquoi `NewScanner` ? </h2> 

`bufio.NewScanner(os.Stdin)` qui permet de scanner l'entrée clavier ligne par ligne

`os.Stdin` = entrée standard = donc ce que l'utilisateur tape au clavier


`"variable".Scan()` = attend que l'utilisateur tape quelque chose