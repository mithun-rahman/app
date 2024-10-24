#!/bin/bash

export CGO_CXXFLAGS="-I/opt/homebrew/include"
export CGO_LDFLAGS="-L/opt/homebrew/lib -L/opt/homebrew/opt/openblas/lib -ldlib -lopenblas -llapack"

go run main.go

#export CGO_ENABLED=1
#
#export CGO_CPPFLAGS="-I/usr/local/Cellar/dlib/19.24.6/include"
#export CGO_LDFLAGS="-L/usr/local/Cellar/dlib/19.24.6/lib -ldlib"
#
#cd /Users/pathaoltd/go/src/app
#
#go run main.go


#brew uninstall dlib
#brew uninstall --force dlib
#
#brew autoremove
#brew cleanup

#/usr/local/bin/brew uninstall --force $(/usr/local/bin/brew list --formula)
#/usr/local/bin/brew uninstall --force $(/usr/local/bin/brew list --cask)
#/usr/local/bin/brew cleanup

#sudo rm -rf /usr/local/Homebrew
#sudo rm -rf /usr/local/bin/brew

#/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
#
#echo 'eval "$(/opt/homebrew/bin/brew shellenv)"' >> /Users/pathaoltd/.zprofile
#eval "$(/opt/homebrew/bin/brew shellenv)"

#file $(brew --prefix openblas)/lib/libopenblas.dylib
#
#file $(which pkg-config)