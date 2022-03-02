# wordlebot
This project propose a Wordle solver (french version)

## Introduction

This project is a golang version of [David Louapre's article](https://scienceetonnante.com/2022/02/13/comment-craquer-le-jeu-wordle-sutom/).

It uses a [french lexical database](http://www.lexique.org/) (140.000 words with frequency in movies or books).
It will select only the 4096 most frequent 5 letters words as a playground (`cmd/extract`).

Then it will compute the best first guess word (`cmd/first`)
And eventually for each game pattern, will provide the best new guess (`cmd/guess`)


## How to use it

1. Download lexical dabase
```
cd assets
curl -O http://www.lexique.org/databases/Lexique383/Lexique383.tsv
```

2. Parse tsv
```
go run cmd/extract/main.go -d -l assets/Lexique383.tsv
```

3. Compute the best word to start with
```
go run cmd/first/main.go -d -l assets/words.txt
```

4. Visualize a word statistics
```
go run cmd/chart/main.go -d -l assets/taris.stat
open bar.html
```

![](stat.png)

5. Start a game

Start a game on your browser (https://wordle.louan.me/)

<img src="game-start.png" width="500">

```
go run cmd/next/main.go -d -l assets/words.txt
```

Then use the words suggested by `wordlebot`
![](cli.png)

Et voila !

<img src="game-win.png" width="500">

6. Run the game simulator (optional)

```
go run cmd/simulator/main.go -d -l assets/words.txt
INFO[0000] start with local words list assets/words.txt
DEBU[0028] SUCCESS ✅ Found word gadjo in 4 attempts
DEBU[0029] SUCCESS ✅ Found word poter in 2 attempts
DEBU[0034] FAILURE ❌ Couldn't find word calao in less than 6 attempts
DEBU[0035] SUCCESS ✅ Found word putti in 4 attempts
...
```

## References

* http://www.lexique.org/
* https://scienceetonnante.com/2022/02/13/comment-craquer-le-jeu-wordle-sutom/
* https://www.youtube.com/watch?v=iw4_7ioHWF4&t=569s
* https://www.youtube.com/watch?v=fRed0Xmc2Wg