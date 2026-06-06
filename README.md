# FunCodex
Simple text encoder made in Lua *(a.k.a. best scripting language on Earth)*

## Main concepts
This program is able to hide a small binary code inside a normal text alternating similar characters (like the latin a and the cyrillic a) between them.

## Disclaimer
I'm making this small project just to try NeoVim and to see if I like its workflow, I know that this program is impractical for many reasons and should not be used in a critical context where messages should be truly hidden, also, its biggest flaw is the **limited space** of binary code you can store in an actual sentence, which makes it work *badly* in small text messages.

This doesn't mean that it's *useless* though, I mean look at it, if you want to tell someone something but you're not brave enough to do so, just send them a FunCodex message and pretend that you did something ;)

Apart from that, I don't even know if there's already something like this, if so, let me know, I want to specify that I'm not trying to copy anyone, I had this idea on my own but I also acknowledge the fact that the web is vast and it's easy to have similar ideas.

# Installation

## Arch Linux

You can install FunCodex using the provided PKGBUILD.

```sh
git clone "https://github.com/DanyiYK/FunCodex/"
cd funcodex
makepkg -si

```

## Manual Installation

#### Requirements

* Lua 

Install Lua using your package manager:

* **Debian**

  ```sh
  sudo apt install lua
  ```

* **Fedora**

  ```sh
  sudo dnf install lua
  ```

---

#### Setup

```sh
git clone https://github.com/DanyiYK/FunCodex
cd FunCodex/src
```

---

## Usage

### 1. Encrypt a message

```sh
lua main.lua crypt "\"The quick brown fox jumps over the lazy dog\" is an English-language pangram – a sentence that contains all 26 letters of the English alphabet.\"" "example"
```

**Output:**

```
CRYPTED TEXT:
"Тhе quiсk brown fох jumрs оver the lаzy dоg" is аn English-languаgе рangrаm – а sеntеnce that contains all 26 letters of the English alphabet.
```

### 2. Decrypt a message

Copy the encrypted text and run:

```sh
lua main.lua decrypt "\"Тhе quiсk brown fох jumрs оver the lаzy dоg\" is аn English-languаgе рangrаm – а sеntеnce that contains all 26 letters of the English alphabet."
```

**Output:**

```
DECRYPTED TEXT:
"example"
```
