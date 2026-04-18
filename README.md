# FunCodex
Simple text encoder made in Lua *(a.k.a. best scripting language on Earth)*

## Main concepts
This program is able to hide a small binary code inside a normal text alternating similar characters (like the latin a and the cyrillic a) between them.

## Disclaimer
I'm making this small project just to try NeoVim and to see if I like its workflow, I know that this program is impractical for many reasons and should not be used in a critical context where messages should be truly hidden, also, its biggest flaw is the **limited space** of binary code you can store in an actual sentence, which makes it work *badly* in small text messages.

This doesn't mean that it's *useless* though, I mean look at it, if you want to tell someone something but you're not brave enough to do so, just send them a FunCodex message and pretend that you did something ;)

Apart from that, I don't even know if there's already something like this, if so, let me know, I want to specify that I'm not trying to copy anyone, I had this idea on my own but I also acknowledge the fact that the web is vast and it's easy to have similar ideas.

## How to use
0. Install Lua

1. Clone the repository wherever you want
```sh
git clone "https://github.com/DanyiYK/FunCodex/"
```

2. Go in the src directory
```sh
cd ./FunCodex/src
```

3. Crypt your first message!
```sh
lua main.lua crypt "\"The quick brown fox jumps over the lazy dog\" is an English-language pangram – a sentence that contains all 26 letters of the English alphabet.\"" "example"
CRYPTED TEXT:
"Тhе quiсk brown fох jumрs оver the lаzy dоg" is аn English-languаgе рangrаm – а sеntеnce that contains all 26 letters of the English alphabet."
```

4. Decrypt the message
To decrypt a message, copy the result of the encryption and paste it in the decrypt command, like this:
```sh
lua main.lua decrypt "\"Тhе quiсk brown fох jumрs оver the lаzy dоg\" is аn English-languаgе рangrаm – а sеntеnce that contains all 26 letters of the English alphabet."
DECRYPTED TEXT:
"example"
```
