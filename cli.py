import os
import sys

version = 'v0.1'

def uninstallService():
    os.system('sudo systemctl disable sfpocedge.service')
    os.system('sudo rm /lib/systemd/system/sfpocedge.service')
    os.system('sudo systemctl daemon-reload')

def installService():
    serviceTemplate = '''[Unit]
    Description=Run edge master
    After=multi-user.target
    [Service]
    Type=simple
    WorkingDirectory={}
    ExecStart={}/edge-app
    Restart=always
    RestartSec=10s
    [Install]
    WantedBy=multi-user.target
    '''

    curDir = os.getcwd()
    service = serviceTemplate.format(curDir, curDir)

    with open('sfpocedge.service', 'w') as f:
        f.write(service)

    os.system('sudo chmod 644 sfpocedge.service')
    os.system('sudo mv sfpocedge.service /lib/systemd/system/')
    os.system('sudo systemctl daemon-reload')
    os.system('sudo systemctl enable sfpocedge.service')

def buildApp():
    os.system('rm edge-app')
    os.system('go build -o edge-app main.go')
    os.system('wget "https://github.com/iot-standards-laboratory/ETRI-SFPOC-EDGE_front/releases/download/{}/web.tar.bz2"'.format(version))
    os.system('tar -xvf web.tar.bz2 && rm web.tar.bz2')
    os.system('mv web www')

if len(sys.argv) < 2:
    buildApp()
    installService()

elif sys.argv[1] == 'build':
    buildApp()

elif sys.argv[1] == 'install':
    installService()

elif sys.argv[1] == 'uninstall':
    uninstallService()
