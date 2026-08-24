## ** Debian**

**Debian** is one of the oldest and most stable **Linux distributions**. It is the foundation for several popular distributions like **Ubuntu** and **Kali Linux**. Debian is known for its **stability**, **security**, and **wide software availability**, making it ideal for servers, desktops, and various development environments.

### **Key Features of Debian:**
1. **Stability**:
   - Debian focuses on stability over cutting-edge features, which makes it a great choice for **servers** or users who need a highly reliable system.
   - It follows a predictable **release cycle**: **stable**, **testing**, and **unstable** (Sid). The **stable** branch is rock-solid, while **testing** is more current but still stable for many users.

2. **Package Management**:
   - Debian uses the **APT** package management system (`apt`, `dpkg`).
   - It provides **.deb** packages for software installation.
   
3. **Wide Software Support**:
   - With a massive repository of over 50,000 packages, Debian offers access to a broad range of free and open-source software.

4. **Security**:
   - Debian includes robust security features and receives **timely updates** to patch vulnerabilities.
   - The **Debian Security Team** actively ensures that security patches are available for all supported versions.

5. **Multiple Architectures**:
   - Debian supports many architectures beyond x86, such as ARM, MIPS, and PowerPC, making it useful for embedded devices and servers.

### **Practical Examples: Basic Commands**

#### **1. Package Management with APT**
Update the list of available packages:
```bash
sudo apt update
```

Install a package (e.g., `nginx`):
```bash
sudo apt install nginx
```

Upgrade installed packages to newer versions:
```bash
sudo apt upgrade
```

#### **2. Service Management with Systemd**
Start a service:
```bash
sudo systemctl start nginx
```

Enable a service to start at boot:
```bash
sudo systemctl enable nginx
```

#### **3. Basic Monitoring and Disk Space**
Check memory usage:
```bash
free -h
```

View disk usage:
```bash
df -h
```
