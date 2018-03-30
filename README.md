# simplekeysd

## Setup go
  brew install golang
  mkdir -p ~/go/bin ~/go/pkg ~/go/src
  
  #Add to .bash_profile:
  export GOPATH="$HOME/go"
  export PATH="/usr/local/opt/go@1.10/bin:$GOPATH/bin:$PATH"
  

go get -u github.com/kardianos/govendor
