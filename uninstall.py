import os

os.system('sudo systemctl disable sfpocedge.service')
os.system('sudo rm /lib/systemd/system/sfpocedge.service')