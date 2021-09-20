#!/bin/bash
cd /usr/local/src
apt-get update && apt-get install -yq gnupg2 lsb-release wget git curl
wget -O - https://files.freeswitch.org/repo/deb/debian-release/fsstretch-archive-keyring.asc | apt-key add -
echo "deb http://files.freeswitch.org/repo/deb/debian-release/ `lsb_release -sc` main" > /etc/apt/sources.list.d/freeswitch.list
echo "deb-src http://files.freeswitch.org/repo/deb/debian-release/ `lsb_release -sc` main" >> /etc/apt/sources.list.d/freeswitch.list
apt-get update -y
apt-get build-dep freeswitch -y
# built from source
wget https://github.com/signalwire/freeswitch/archive/refs/tags/v1.10.6.tar.gz -O freeswitch-1.10.6.tar.gz
tar -xzvf freeswitch-1.10.6.tar.gz
cd freeswitch-1.10.6
./bootstrap.sh -j
mv modules.conf modules.conf.origin
cat >modules.conf<< EOF
applications/mod_commands
applications/mod_conference
applications/mod_av
applications/mod_avmd
applications/mod_dptools
applications/mod_distributor
applications/mod_spandsp
applications/mod_spy
applications/mod_valet_parking
applications/mod_voicemail
applications/mod_voicemail_ivr
applications/mod_callcenter
applications/mod_fifo
applications/mod_valet_parking
asr_tts/mod_flite
codecs/mod_opus
dialplans/mod_dialplan_xml
endpoints/mod_sofia
event_handlers/mod_event_socket
formats/mod_sndfile
loggers/mod_console
loggers/mod_logfile
EOF
./configure -C --prefix=/usr/local --with-rundir=/run/freeswitch --with-logfiledir=/var/log/freeswitch/ --enable-64 --with-openssl && make && make install
mv /usr/local/etc/freeswitch /usr/local/etc/freeswitch.origin
