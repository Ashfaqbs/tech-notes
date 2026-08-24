## **Shell & Types of Shells (Bash vs Zsh)**
### **1️⃣ What is a Shell?**
A **shell** is a command-line interpreter that allows users to interact with the operating system. It processes commands entered by the user and executes system-level tasks.

### **2️⃣ Types of Shells in Linux**
Linux has different types of shells, but the most commonly used are:

| Shell | Description |
|-------|------------|
| **Bash (Bourne Again Shell)** | Default shell in most Linux distributions, powerful and widely used. |
| **Zsh (Z Shell)** | An improved shell with better features, customization, and usability. |
| **Fish (Friendly Interactive Shell)** | User-friendly, auto-suggestions, but less script-compatible. |
| **Tcsh (TENEX C Shell)** | Based on C Shell, used in some Unix systems. |

---

## **3️⃣ Bash vs Zsh**
| Feature | Bash | Zsh |
|---------|------|-----|
| **Auto-Suggestions** | ❌ No | ✅ Yes |
| **Plugins & Customization** | 🔸 Limited | ✅ Extensive (Oh-My-Zsh, etc.) |
| **Better History Search** | ❌ No | ✅ Yes (prefix-based search) |
| **Spelling Correction** | ❌ No | ✅ Yes |
| **Globbing (wildcard expansion)** | 🔸 Basic | ✅ Advanced (`**` for recursive matching) |
| **Theme Support** | ❌ No | ✅ Yes (Powerlevel10k, etc.) |

### **Why Do People Prefer Zsh?**
- **Developers & Power Users:** Zsh provides better **auto-completion, plugins, themes, and scripting capabilities**.
- **Oh-My-Zsh Framework:** Makes Zsh more powerful with easy customization.
- **Better User Experience:** Features like **auto-suggestions, spelling correction, and better history navigation** make it a preferred choice.

---

## **4️⃣ What We Did to Switch from Bash to Zsh**
### **Step 1: Install Zsh & Auto-Suggestions**
```bash
sudo apt update && sudo apt install -y zsh zsh-autosuggestions
```
This installed:
- **Zsh** (`/usr/bin/zsh`)
- **Auto-Suggestions** (to show command hints)

### **Step 2: Change Default Shell to Zsh**
```bash
chsh -s $(which zsh)
```
- `which zsh` → Finds the installed location of Zsh (`/usr/bin/zsh`).
- `chsh -s` → Changes the **default login shell** for the user.

### **Step 3: Reboot & Login Again**
After logging back in, Zsh started but showed the **new user setup prompt** because no `.zshrc` file existed.

### **Step 4: Created `.zshrc` File**
```bash
echo "# My Zsh Configuration" > ~/.zshrc
source ~/.zshrc
```
- Prevented Zsh from asking for setup every time.

### **Step 5: Enabled Auto-Suggestions & Syntax Highlighting**
```bash
echo "source /usr/share/zsh-autosuggestions/zsh-autosuggestions.zsh" >> ~/.zshrc
sudo apt install zsh-syntax-highlighting -y
echo "source /usr/share/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh" >> ~/.zshrc
source ~/.zshrc
```
- Now, Zsh provides **suggestions and syntax highlighting**.

---

## **5️⃣ Understanding `/usr/bin/zsh` & Why It Points There**
### **What is `/usr/bin/zsh`?**
- When we ran:
  ```bash
  which zsh
  ```
  we got:
  ```
  /usr/bin/zsh
  ```
  - This means Zsh **is installed in `/usr/bin/`**.
  - **All system-installed executables (like Bash, Zsh, Python, etc.) go inside `/usr/bin/`**.
  - `/usr/bin/` is in the system's `$PATH`, so typing `zsh` runs the binary from this location.

### **Where Are Programs Stored in Linux?**
| Directory | Purpose |
|-----------|---------|
| **`/bin/`** | Essential system binaries (e.g., `bash`, `ls`, `cat`) |
| **`/usr/bin/`** | User binaries (installed via `apt`, `snap`, etc.) |
| **`/sbin/`** | System admin binaries (e.g., `fdisk`, `iptables`) |
| **`/usr/local/bin/`** | Manually installed binaries (not from package manager) |

---

## **Conclusion**
- **Zsh is an improved shell over Bash**, preferred for development due to better auto-suggestions, plugins, and customization.
- **We installed Zsh, set it as default, and added useful features**.
- **`which zsh` shows `/usr/bin/zsh`** because all installed system programs reside there.
