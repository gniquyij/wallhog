echo '''
alias hg="wallhog () { echo '$@'; unset -f '$(go env GOPATH)'/bin/wallhog; }; '$(go env GOPATH)'/bin/wallhog"
''' >> $HOME/.bash_profile