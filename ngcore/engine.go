package ngcore

import (
    "fmt"
		"os"
		"os/exec"
)

const FSXML = `
<?xml version="1.0"?>
<document type="freeswitch/xml">
  <X-PRE-PROCESS cmd="set" data="dblocation=/dev/shm"/>
  <X-PRE-PROCESS cmd="set" data="useragent=Comm v0.0.0"/>
  <X-PRE-PROCESS cmd="set" data="commserver=127.0.0.1:9021"/>
  <!-- BEGIN: CONFIGURATION SECTION -->
  <section name="configuration" description="Various Configuration">
    <!-- ACL ACCESS CONTROL LIST -->
    <configuration name="acl.conf" description="Network Lists">
      <network-lists>
        <list name="WHITE.ACL.LIST" default="deny">
          <node type="allow" cidr="0.0.0.0/0"/>
        </list>
      </network-lists>
    </configuration>
    <!-- CONSOLE LOGGER -->
    <configuration name="console.conf" description="Console Logger">
      <mappings>
        <map name="all" value="console,debug,info,notice,warning,err,crit,alert"/>
      </mappings>
      <settings>
        <param name="colorize" value="true"/>
        <param name="loglevel" value="info"/>
      </settings>
    </configuration>
    <!-- DISTRIBUTOR -->
    <configuration name="distributor.conf" description="Distributor Configuration">
      <lists>
        <list name="dummies" total-weight="2">
          <node name="dummy1" weight="1"/>
          <node name="dummy2" weight="1"/>
        </list>
      </lists>
    </configuration>
    <!-- EVENT SOCKET -->
    <configuration name="event_socket.conf" description="Event Socket">
      <settings>
        <param name="listen-ip" value="127.0.0.1"/>
        <param name="listen-port" value="8021"/>
        <param name="password" value="ramdomstr"/>
      </settings>
    </configuration>
    <!--- LOGFILE -->
    <configuration name="logfile.conf" description="File Logging">
      <settings>
        <param name="rotate-on-hup" value="true"/>
      </settings>
      <profiles>
        <profile name="default">
          <settings>
            <param name="rollover" value="104857600"/>
              <param name="maximum-rotate" value="32"/>
            <param name="uuid" value="true" />
          </settings>
          <mappings>
            <map name="all" value="console,info,notice,warning,err,crit,alert"/>
          </mappings>
        </profile>
      </profiles>
    </configuration>
    <!--- MODULES -->
    <configuration name="modules.conf" description="Modules">
      <modules>
        <load module="mod_console"/>
        <load module="mod_logfile"/>
        <load module="mod_sofia"/>
        <load module="mod_distributor"/>
        <load module="mod_event_socket"/>
        <load module="mod_commands"/>
        <load module="mod_dptools"/>
        <load module="mod_dialplan_xml"/>
        <load module="mod_bcg729"/>
        <load module="mod_spandsp"/>
        <load module="mod_sndfile"/>
        <load module="mod_opus"/>
        <load module="mod_flite"/>
      </modules>
    </configuration>
    <configuration name="opus.conf">
      <settings>
        <param name="use-vbr" value="1"/>
        <param name="use-dtx" value="1"/>
        <param name="complexity" value="10"/>
        <param name="packet-loss-percent" value="10"/>
        <param name="keep-fec-enabled" value="1"/>
        <param name="use-jb-lookahead" value="1"/>
        <param name="maxaveragebitrate" value="64000"/>
        <param name="maxplaybackrate" value="48000"/>
        <param name="sprop-maxcapturerate" value="48000"/>
        <param name="adjust-bitrate" value="1"/>
      </settings>
    </configuration>
    <!--- SOFIA SIP -->
    <configuration name="sofia.conf" description="Sofia SIP Endpoint">
      <global_settings>
        <param name="log-level" value="0"/>
        <param name="debug-presence" value="0"/>
      </global_settings>
      <profiles>
        <profile name="defaultsip">
          <settings>
            <param name="debug" value="0"/>
            <param name="sip-trace" value="yes"/>
            <param name="sip-capture" value="no"/>
            <param name="user-agent-string" value="$${useragent}"/>
            <param name="username" value="root" />
            <param name="shutdown-on-fail" value="true"/>
            <param name="parse-all-invite-headers" value="true"/>
            <param name="sip-options-respond-503-on-busy" value="true"/>
            <param name="inbound-late-negotiation" value="true"/>
            <param name="disable-register" value="true"/>
            <param name="disable-transfer" value="false"/>
            <param name="manual-redirect" value="true"/>
            <param name="proxy-refer" value="false"/>
            <param name="disable-hold" value="true"/>
            <param name="proxy-hold" value="true"/>
            <param name="send-display-update" value="false"/>
            <param name="inbound-proxy-media" value="false"/>
            <param name="inbound-codec-negotiation" value="generous"/>
            <param name="codec-prefs" value="PCMA,PCMU,G729,OPUS"/>
            <param name="suppress-cng" value="true"/>
            <param name="pass-callee-id" value="false" />
            <param name="enable-timer" value="false"/>
            <param name="stun-enabled" value="false"/>
            <param name="t38-passthru" value="true"/>
            <param name="rtp-rewrite-timestamps" value="true"/>
            <param name="rtp-timer-name" value="soft"/>
            <param name="enable-100rel" value="true"/>
            <param name="dtmf-type" value="rfc2833"/>
            <param name="pass-rfc2833" value="true"/>
            <param name="rfc2833-pt" value="101"/>
            <param name="rtp-digit-delay" value="20"/>
            <param name="local-network-acl" value="rfc1918.auto"/>
            <param name="apply-nat-acl" value="nat.auto"/>
            <param name="dialplan" value="XML"/>
            <param name="context" value="default"/>
            <param name="dbname" value="$${dblocation}/defaultsip.db"/>
            <param name="rtp-ip" value="$${local_ip_v4}"/>
            <param name="sip-ip" value="$${local_ip_v4}"/>
            <param name="sip-port" value="5060"/>
          </settings>
        </profile>
      </profiles>
    </configuration>
    <!--- SPANDSP -->
    <configuration name="spandsp.conf" description="SpanDSP configuration">
    </configuration>
    <!-- FREESWITCH DEFAULT CONFIGURATION -->
    <configuration name="switch.conf" description="Core Configuration">
      <settings>
        <param name="core-db-name" value="$${dblocation}/commcore.db"/>
        <param name="max-sessions" value="24000"/>
        <param name="sessions-per-second" value="1200"/>
        <param name="rtp-start-port" value="16384"/>
        <param name="rtp-end-port" value="32767"/>
        <param name="min-idle-cpu" value="15"/>
        <param name="event-heartbeat-interval" value="30"/>
      </settings>
    </configuration>
  </section>
  <!-- END: CONFIGURATION SECTION -->
  <!-- BEGIN: DIALPLAN SECTION -->
  <section name="dialplan" description="Regex/XML Dialplan">
    <context name="redirected">
      <extension name="any_to_any">
        <condition regex="all">
          <regex field="${acl(${network_addr} WHITE.ACL.LIST)}" expression="true"/>
          <regex field="destination_number" expression="."/>
          <action application="sched_hangup" data="+3600 ALLOTTED_TIMEOUT"/>
          <action application="park"/>
          <anti-action application="hangup" data="CALL_REJECTED"/>
          <anti-action application="log" data="WARNING: YOU SHOULD NOT FOUND THIS MESSAGE - THE THING WENT WRONG"/>
        </condition>
      </extension>
    </context>
    <context name="default">
      <extension name="any_to_any">
        <condition regex="all">
          <regex field="${acl(${network_addr} WHITE.ACL.LIST)}" expression="true"/>
          <regex field="destination_number" expression="."/>
          <action application="sched_hangup" data="+3600 ALLOTTED_TIMEOUT"/>
          <action application="park"/>
          <anti-action application="hangup" data="CALL_REJECTED"/>
          <anti-action application="log" data="WARNING: YOU SHOULD NOT FOUND THIS MESSAGE - THE THING WENT WRONG"/>
        </condition>
      </extension>
    </context>
  </section>
  <!-- END: DIALPLAN SECTION -->
</document>
`

func Engine() {
		/*
    data, err := ioutil.ReadFile("../configs/freeswitch.xml")
    if err != nil {
        fmt.Println("File reading error", err)
		fmt.Println("--------------------------------------------------------------------------", string(FSXML))
        return
    }
    fmt.Println("Contents of file:", string(data))
		*/

		// create file
    f, err := os.Create("/usr/local/etc/freeswitch/freeswitch.xml")
    if err != nil {
        fmt.Println(err)
        return
    }
		// write string
    l, err := f.WriteString(FSXML)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
		// close the file
    fmt.Println(l, "bytes /usr/local/etc/freeswitch/freeswitch.xml written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }

		// start freeswitch
    fscmd := exec.Command("/usr/local/bin/freeswitch", "-nc")
    err = fscmd.Run()
    if err != nil {
			fmt.Println(err)
    }
		fmt.Println(l, "start freeswitch")
}
