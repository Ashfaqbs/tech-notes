## ** Kali Linux**

**Kali Linux** is a Debian-based distribution designed specifically for **penetration testing**, **ethical hacking**, and **digital forensics**. It includes a vast collection of **security tools** and is highly regarded by cybersecurity professionals.

### **Key Features of Kali Linux:**
1. **Purpose-Built for Security Testing**:
   - Kali Linux is preloaded with hundreds of **penetration testing tools** for tasks like network scanning, exploitation, forensic analysis, and vulnerability testing.
   - Tools like **Metasploit**, **Nmap**, **Wireshark**, and **John the Ripper** come pre-installed.

2. **Based on Debian**:
   - Since Kali is built on **Debian**, it inherits many of Debian's strengths, such as stability, the APT package manager, and security features.
   
3. **Rolling Release**:
   - Kali follows a **rolling release** model, meaning it continuously receives updates without needing to reinstall new versions.
   
4. **Customized for Security**:
   - Kali has special features, such as **Live Boot** (can run entirely from a USB drive without installing it), **forensic mode**, and **stealth mode**.

5. **Pre-configured for Pentesting**:
   - Networking features like **monitor mode** and **packet injection** are enabled by default for wireless testing.
   - It comes with many tools pre-configured and ready for use in penetration testing scenarios.

### **Practical Examples: Basic Commands**

#### **1. Package Management with APT**
Since Kali is based on Debian, it also uses the **APT** package manager.

Update package lists:
```bash
sudo apt update
```

Install Metasploit (pre-installed, but for practice):
```bash
sudo apt install metasploit-framework
```

Upgrade all installed packages:
```bash
sudo apt upgrade
```

#### **2. Using Security Tools**
Launch **Nmap** for network scanning:
```bash
sudo nmap -sP 192.168.1.0/24
```

Start **Metasploit Framework**:
```bash
sudo msfconsole
```

#### **3. Service Management with Systemd**
Start the **SSH** service (commonly used during pentests):
```bash
sudo systemctl start ssh
```

Enable SSH to start at boot:
```bash
sudo systemctl enable ssh
```

