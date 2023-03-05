rm mb4.rar
rar a mb4.rar mb4 pages/*.html assets/* -r -s
sshpass -p 'vMbXcWBTensq' scp mb4.rar ubuntu@162.19.152.177:/var/www/mb4/

sshpass -p 'vMbXcWBTensq' ssh ubuntu@162.19.152.177 "sudo systemctl stop mb4.service"
sshpass -p 'vMbXcWBTensq' ssh ubuntu@162.19.152.177 "cd /var/www/mb4/ ; unrar x -y mb4.rar"
sshpass -p 'vMbXcWBTensq' ssh ubuntu@162.19.152.177 "cd /var/www/mb4/ ; chmod +x mb4"
sshpass -p 'vMbXcWBTensq' ssh ubuntu@162.19.152.177 "sudo systemctl start mb4.service"
