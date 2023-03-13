import os

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