2. UNIX/Linux

Bash Shell (Common across Linux, macOS, UNIX)

ls

Scenario: List files and directories in the current directory.
Usage: ls -l (Long listing with details).

cd

Scenario: Change the current directory.
Usage: cd /var/log

cp

Scenario: Copy files or directories.
Usage: cp file.txt /home/user/backup/

mv

Scenario: Move or rename files or directories.
Usage: mv oldfile.txt newfile.txt

rm

Scenario: Remove files or directories.
Usage: rm -r /path/to/directory/

cat

Scenario: Concatenate and display the contents of a file.
Usage: cat /etc/passwd

grep

Scenario: Search for text patterns in files.
Usage: grep 'error' /var/log/syslog

find

Scenario: Search for files and directories.
Usage: find / -name "*.conf"

chmod

Scenario: Change file permissions.
Usage: chmod 755 script.sh

chown

Scenario: Change file owner and group.
Usage: chown user:group file.txt

ps

Scenario: Display currently running processes.
Usage: ps aux

kill

Scenario: Terminate a process by its PID.
Usage: kill -9 1234

df

Scenario: Display disk space usage.
Usage: df -h

top

Scenario: Monitor system processes and resource usage.
Usage: top

nano

Scenario: Edit files directly in the terminal.
Usage: nano /etc/hosts

tar

Scenario: Archive files and directories.
Usage: tar -czvf archive.tar.gz /path/to/directory

ssh

Scenario: Securely connect to a remote server.
Usage: ssh user@remote_host

iptables

Scenario: Configure IP packet filter rules (firewall).
Usage: iptables -A INPUT -p tcp --dport 22 -j ACCEPT

sudo

Scenario: Execute a command with elevated (root) privileges.
Usage: sudo apt update