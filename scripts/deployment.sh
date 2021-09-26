scp -v -r ~/codespace/comm COMM-UAT:/root/comm
rsync -rv -e ssh --include '*/' --include='.git/*' ~/codespace/comm COMM-UAT:/root/comm
