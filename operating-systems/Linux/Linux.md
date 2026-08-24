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



Both `tar` and `zip` are used for **archiving** and **compressing** files, but they work differently and serve distinct purposes. Here’s a detailed comparison:

### **1. `tar` (Tape Archive)**
- **Primary Use**: `tar` is used to **combine multiple files** into a single archive file, often called a **tarball**. It can compress the archive, but it doesn’t do so by default.
- **Common File Extensions**:
  - `.tar`: An archive created with `tar` without compression.
  - `.tar.gz` or `.tgz`: A `tar` archive compressed with **gzip**.
  - `.tar.bz2`: A `tar` archive compressed with **bzip2**.

- **Common Commands**:
  - **Create an Archive (`-c`)**:
    ```bash
    tar -cvf archive.tar file1 file2 directory/
    ```
    - Creates a `tar` archive named `archive.tar` containing `file1`, `file2`, and the `directory/`.
  - **Compress with gzip (`-z`)**:
    ```bash
    tar -czvf archive.tar.gz file1 file2 directory/
    ```
    - Creates a gzip-compressed archive `archive.tar.gz`.
  - **Extract an Archive (`-x`)**:
    ```bash
    tar -xvf archive.tar
    ```
    - Extracts the contents of `archive.tar`.
  - **List Contents (`-t`)**:
    ```bash
    tar -tvf archive.tar
    ```
    - Lists the contents of the `tar` archive.

- **Strengths**:
  - Can archive **directories** and **preserve file permissions** and structures, making it ideal for backup and distribution of software source code.
  - Often used in combination with compression tools like `gzip` or `bzip2` for effective compression.

- **Weakness**:
  - Requires an additional compression tool for reducing the archive size (`gzip`, `bzip2`, etc.).
  - Extracting a specific file from a large tarball can be slower because `tar` has to read the entire archive to locate the file.



### **2. `zip`**
- **Primary Use**: `zip` is used to **compress** files individually and combine them into a single `.zip` archive. Each file is compressed separately, and the archive contains a directory of all the compressed files.
- **Common File Extensions**:
  - `.zip`: A compressed archive containing one or more files.

- **Common Commands**:
  - **Create a Zip Archive**:
    ```bash
    zip archive.zip file1 file2 directory/
    ```
    - Creates a `zip` archive named `archive.zip` containing `file1`, `file2`, and the contents of `directory/`.
  - **Add Files to an Existing Archive**:
    ```bash
    zip archive.zip file3
    ```
    - Adds `file3` to the existing `archive.zip`.
  - **Extract a Zip Archive**:
    ```bash
    unzip archive.zip
    ```
    - Extracts the contents of `archive.zip`.
  - **List Contents**:
    ```bash
    unzip -l archive.zip
    ```
    - Lists the contents of the `zip` archive.

- **Strengths**:
  - Compresses each file separately, so if one file is corrupted, others are still accessible.
  - More commonly used on **Windows** and cross-platform scenarios.
  - Supports partial extraction, which means we can extract only specific files without processing the entire archive.

- **Weakness**:
  - Does not preserve Unix-specific file permissions and metadata, making it less suitable for system backups.
  - Compresses files individually, so it may not be as efficient as `tar` combined with compression for large numbers of small files.

### **Summary of Differences**:

| Feature               | `tar`                                 | `zip`                              |
|-----------------------|---------------------------------------|------------------------------------|
| **Usage**             | Archiving with optional compression   | Compression and archiving          |
| **Compression**       | Requires external tools (`gzip`, etc.)| Built-in file-by-file compression  |
| **File Extension**    | `.tar`, `.tar.gz`, `.tar.bz2`         | `.zip`                             |
| **Handling Directories** | Archives directory structure well  | Flattens directories unless using `-r` |
| **Platform**          | Common in Unix/Linux                  | More common in cross-platform use  |
| **File Permissions**  | Preserves Unix permissions            | Does not preserve Unix permissions |
| **Performance**       | Efficient for archiving many files    | Can be less efficient with many small files |
| **Partial Extraction**| Must read the entire archive          | Can extract specific files easily  |

### **When to Use Which?**
- **Use `tar`**: For creating backups of directory structures, preserving permissions, and combining with compression tools for efficient storage. Ideal for **Linux/Unix** systems.
- **Use `zip`**: For sharing files across different platforms (especially with Windows users), or when we need easy and flexible compression. Ideal for **cross-platform** compatibility.


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



Ethical Hacking CommandsCommon Across Multiple OS

nmap

Scenario: Network scanning and mapping.
Usage: nmap -sP 192.168.1.0/24

tcpdump

Scenario: Capture and analyze network packets.
Usage: tcpdump -i eth0 -w capture.pcap

netcat (nc)

Scenario: Establish network connections and transfer data.
Usage: nc -lvp 4444 (Listen for incoming connections on port 4444).

hydra

Scenario: Brute-force login credentials.
Usage: hydra -l admin -P passwords.txt ssh://192.168.1.10

john the ripper

Scenario: Password cracking.
Usage: john --wordlist=passwords.txt hashfile

metasploit

Scenario: Exploitation framework for discovering, exploiting, and validating vulnerabilities.
Usage: msfconsole




### **Comparison Between Debian, Kali, and Arch Linux**:

| **Feature**                | **Debian**                         | **Kali Linux**                     | **Arch Linux**                     |
|----------------------------|------------------------------------|------------------------------------|------------------------------------|
| **Main Focus**              | Stability and wide package support | Penetration testing and security   | Customizability and simplicity     |
| **Package Manager**         | `APT`                              | `APT`                              | `Pacman` (with AUR support)        |
| **Release Model**           | Stable (also Testing and Unstable) | Rolling                            | Rolling                            |
| **Pre-installed Software**  | Minimal, but with extensive repos  | Security tools (Nmap, Metasploit)  | Minimalist; user installs everything|
| **Customization**           | Moderate                          | Moderate                           | High                               |
| **Best For**                | Servers, workstations, and developers | Ethical hackers and cybersecurity experts | Advanced users who want full control|

---

### **Summary**:
- **Debian**: Great for stability, security, and broad software support. Often used for both servers and desktops.
- **Kali Linux**: Tailored for cybersecurity professionals and penetration testers with pre-installed tools for ethical hacking.
- **Arch Linux**: Provides total control and customization for users who want a minimalist system, ideal for advanced Linux users.

Each of these distributions has its strengths, so choosing the right one depends on the use case and comfort level with Linux.
