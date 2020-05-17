echo 'export WALLHOG_PATH='${PWD}'' >> $HOME/.bash_profile
go build wh.go
#echo '''
#wallhog () { '${PWD}'/wh; }
#alias wh=wallhog
#''' >> $HOME/.bash_profile