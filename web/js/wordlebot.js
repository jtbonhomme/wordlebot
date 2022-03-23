
let iframeDocument
let buttons
let keys = []
let firstGuess = "TARIS"
let log

function init() {
    console.log("init ...")
    iframeDocument = window.frames[0].document;
    buttons = iframeDocument.getElementsByTagName("BUTTON")
    keys['A'] = buttons[0]
    keys['B'] = buttons[25]
    keys['C'] = buttons[23]
    keys['D'] = buttons[12]
    keys['E'] = buttons[2]
    keys['F'] = buttons[13]
    keys['G'] = buttons[14]
    keys['H'] = buttons[15]
    keys['I'] = buttons[7]
    keys['J'] = buttons[16]
    keys['K'] = buttons[17]
    keys['L'] = buttons[18]
    keys['M'] = buttons[19]
    keys['N'] = buttons[26]
    keys['O'] = buttons[8]
    keys['P'] = buttons[9]
    keys['Q'] = buttons[10]
    keys['R'] = buttons[3]
    keys['S'] = buttons[11]
    keys['T'] = buttons[4]
    keys['U'] = buttons[6]
    keys['V'] = buttons[24]
    keys['W'] = buttons[21]
    keys['X'] = buttons[22]
    keys['Y'] = buttons[5]
    keys['Z'] = buttons[1]
    keys['ENTER'] = buttons[20]

    log = document.getElementById("log")

    typeWord(firstGuess)
}

function sleep(milliseconds) {
    const date = Date.now();
    let currentDate = null;
    do {
        currentDate = Date.now();
    } while (currentDate - date < milliseconds);
}

function typeWord(word) {
    if (word.length > 5) {
        return
    }
    logGuess(word)
    for (i=0; i < word.length; i++) {
        typeLetter(word[i].toUpperCase())
    }
    typeLetter('ENTER')
}

function logGuess(word) {
    if (word.length > 5) {
        return
    }
    log.innerText = word.toUpperCase()
}

function typeLetter(letter) {
    keys[letter].click()
}