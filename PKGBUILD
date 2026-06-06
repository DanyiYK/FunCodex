pkgname=funcodex
pkgver=0.2.0
pkgrel=1
pkgdesc="Simple text encoder made in Lua"
arch=('x86_64')
url="https://github.com/DanyiYK/FunCodex"
license=('MIT')
depends=('lua')

package() {
  cd "$startdir"
  install -Dm644 funcodex.1 \
    "$pkgdir/usr/share/man/man1/funcodex.1"

  install -d "$pkgdir/usr/share/$pkgname/src"
  cp -a src/* "$pkgdir/usr/share/$pkgname/src/"

  install -Dm755 /dev/stdin \
    "$pkgdir/usr/bin/$pkgname" <<'EOF'
#!/bin/sh
exec lua /usr/share/funcodex/src/main.lua "$@"
EOF
}
