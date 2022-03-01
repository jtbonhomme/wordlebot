# wordlebot
This project propose a Wordle solver (french version)

## Introduction

This project is a golang version of [David Louapre's article]https://scienceetonnante.com/2022/02/13/comment-craquer-le-jeu-wordle-sutom/().

It uses a [french lexical(http://www.lexique.org/)] (140.000 words with frequency) as database.
It will select only the most frequent 5 letters words as a playground (`cmd/db`).

Then it will compute the best first guess word (`cmd/first`)
And eventually for each game pattern, will provide the best new guess (`cmd/guess`)

You need to download the lexical database with:
```
cd assets
curl -O http://www.lexique.org/databases/Lexique383/Lexique383.tsv
```

## How to use it


## References

* http://www.lexique.org/
* https://scienceetonnante.com/2022/02/13/comment-craquer-le-jeu-wordle-sutom/
* https://www.youtube.com/watch?v=iw4_7ioHWF4&t=569s
* https://www.youtube.com/watch?v=fRed0Xmc2Wg