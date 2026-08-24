why both **Linux** and **Unix** scripts often have the `.sh` extension, even though they might use different shells like **Bash**, **Ksh**, or **Sh**.

### **1. The `.sh` Extension**
- The `.sh` extension is traditionally used for **shell scripts**, regardless of whether we're using **Bash**, **Sh**, **Ksh**, or any other shell. It stands for "**shell script**."
  
- It’s a **convention** rather than a strict rule. The `.sh` extension simply indicates that the file contains shell commands, but it doesn't force the system to use any specific shell to interpret the script. The actual shell that runs the script is determined by either:
  - The **shebang (`#!`)** line at the top of the script, or
  - The shell used to execute the script from the command line.

### **2. Linux and Unix Shells**
- **Linux** primarily uses **Bash** as its default shell, but it can also use **Sh**, **Zsh**, **Ksh**, etc. 
- **Unix** systems (like AIX or Solaris) might use **Sh** or **Ksh**, but modern Unix systems often support **Bash** too.
  
Even though **Bash** and **Sh** are different shells, both can execute scripts with a `.sh` extension. The file extension doesn't determine the shell—**the shebang does**.

### **3. Shebang (`#!`) Defines the Shell**
The **first line** of a shell script is crucial because it tells the system **which shell** to use to run the script. This is called the **shebang** (or hashbang), and it looks like this:

```bash
#!/bin/bash
```

or

```sh
#!/bin/sh
```

- If we want wer script to explicitly use **Bash**, we'd start the script with `#!/bin/bash`.
- If we want it to use **Ksh**, we'd start with `#!/bin/ksh`.

### **4. Why Use `.sh` for Both?**
- **Convenience and convention**: Developers often use `.sh` for any shell script, regardless of the specific shell (Bash, Ksh, etc.) because it's a common and recognized extension.
- **Interchangeability**: Many shells (like **Sh**, **Bash**, and **Ksh**) are compatible with one another for basic commands. A script written for `sh` will often run in `bash` or `ksh` without issues.
  
### **5. Example**:
Here's an example script that uses **Ksh** but still has a `.sh` extension:
```ksh
#!/bin/ksh
echo "This script is running with KornShell!"
```
And here's one that uses **Bash** but also uses the `.sh` extension:
```bash
#!/bin/bash
echo "This script is running with Bash!"
```

### **6. What Happens Without a Shebang?**
If there is **no shebang** in the script, the shell that interprets the script will be the one we’re currently using in wer terminal. For example:
- If we're in a **Bash** shell and run `./script.sh`, Bash will interpret the script.
- If we're in a **Ksh** shell and run the same script, **Ksh** will interpret it.

### **Summary**:
- The `.sh` extension is a **general convention** for shell scripts across **Unix**, **Linux**, and other systems.
- **Shebang (`#!`)** at the top of the script defines **which shell** (Bash, Ksh, Sh, etc.) will execute the script.
- Even though Linux often uses **Bash** and Unix might use **Ksh** or **Sh**, both systems commonly use `.sh` as a universal extension for shell scripts.








Let's break down wer script first, and then I'll address the error we're encountering.

### **wer Script Explanation**

```bash
#!/bin/bash
pid=$(ps -ef | grep 'SB-TelegramBOT.jar' | grep -v grep | awk '{print $2}')
if [ -n "$pid" ]; then
  kill -9 $pid
fi
```

#### **Stop Script Explanation**:

1. **Shebang (`#!/bin/bash`)**:
   - This tells the system to use the **Bash** shell to run the script.

2. **`pid=$(ps -ef | grep 'SB-TelegramBOT.jar' | grep -v grep | awk '{print $2}')`**:
   - **`ps -ef`**: This command lists all running processes on the system.
   - **`grep 'SB-TelegramBOT.jar'`**: Filters the output of `ps -ef` to only include lines that contain `SB-TelegramBOT.jar`. This is likely wer Spring Boot application.
   - **`grep -v grep`**: Ensures that the line containing the `grep` command itself is **excluded** from the results.
   - **`awk '{print $2}'`**: Extracts the second column from the output, which contains the **process ID (PID)** of the running `SB-TelegramBOT.jar` application.
   - **`pid=$(...)`**: The output (the PID) is stored in the variable `pid`.

3. **`if [ -n "$pid" ]; then`**:
   - **`-n "$pid"`**: This checks if `pid` is non-empty (i.e., if a process ID was found).

4. **`kill -9 $pid`**:
   - If a PID was found, this command sends a **SIGKILL (kill -9)** signal to forcefully terminate the process.

#### **What This Script Does**:
- The script searches for the process associated with the **`SB-TelegramBOT.jar`** application.
- If it finds the process, it forcefully **kills** it using `kill -9` (which terminates the process immediately, regardless of what it's doing).






### why we need to use `./file.sh` instead of just `file.sh` when we're in the same directory as the script:

### **1. The Current Directory (`.`) and Executable Files**:
When we're in a folder and want to run a script using `./file.sh`, the `./` tells the shell to look for the script in the **current directory** (`.` stands for "current directory").

- **`./file.sh`** means "run the script `file.sh` in the current directory."
- If we just type `file.sh`, the shell will search for `file.sh` in the directories listed in wer **`$PATH`** environment variable, not in the current folder. Since the current directory is usually **not** in `$PATH` by default (for security reasons), the shell won’t find wer script unless we explicitly tell it to look in the current directory by using `./`.

### **2. Why Not `file.sh`?**
- **Security**: By default, most shells do not include the current directory (`.`) in the **`$PATH`** variable. This prevents malicious or unintended scripts in wer current directory from accidentally running when we call commands.
  
- **Without `./`**, the shell searches through the directories listed in **`$PATH`** (e.g., `/bin`, `/usr/bin`), but the **current directory is usually not included** in `$PATH`.

### **3. The `./` Prefix**:
- The `./` prefix explicitly tells the shell: **"Look for this script in the current directory."**
- Example:
  ```bash
  ./file.sh
  ```
  This will run the `file.sh` script located in the current directory.

### **4. Making a Script Globally Accessible**:
If we want to be able to run `file.sh` from anywhere without needing to use `./`, we have two options:
  
#### **a. Add the Script’s Directory to the `$PATH`**:
  - we can add the directory containing `file.sh` to wer **`$PATH`** environment variable, so the shell will always look there.
  - Add the following line to wer `.bashrc` or `.zshrc` file (depending on wer shell):
    ```bash
    export PATH=$PATH:/path/to/wer/script/directory
    ```
  - After that, we can just run `file.sh` without needing `./`.

#### **b. Move the Script to a Directory Already in `$PATH`**:
  - Move the script to a directory that’s already in wer **`$PATH`**, like `/usr/local/bin`:
    ```bash
    sudo mv file.sh /usr/local/bin/
    ```
  - Once it’s in `/usr/local/bin`, we can call `file.sh` from anywhere.

---

### **Summary**:
- we need to use `./file.sh` to run a script in the **current directory** because the shell does not look in the current directory by default.
- If we want to run the script without `./`, we can either:
  - Add the script’s directory to the **`$PATH`**, or
  - Move the script to a directory that’s already in **`$PATH`**.









#### we want a script in **Linux**, **Unix**, and **Windows** to run a **Spring Boot Java program** (a `.jar` file) with specific **heap size** settings and **port number**, along with handling situations where the Java executable might not be in the standard path. 

### 1. **Linux (Bash) – `.sh` Script**
We'll create a Bash script to set **min/max heap** size, specify a **port** for the Spring Boot app, and handle custom **JAVA_HOME** paths (like IBM Java).

```bash
#!/bin/bash

# Define the JAR file location and Java options (heap size, port)
JAR_FILE="wer-springboot-app.jar"
JAVA_OPTS="-Xms256m -Xmx1024m"
SPRING_OPTS="--server.port=8080"

# Set JAVA_HOME if needed (use specific Java version if not on PATH)
# we can replace this path with wer IBM Java or other specific Java versions.
if [ -z "$JAVA_HOME" ]; then
  export JAVA_HOME="/opt/ibm/java8"  # Change this path as necessary
fi

# Ensure the Java binary is available
if [ ! -x "$JAVA_HOME/bin/java" ]; then
  echo "Java not found at $JAVA_HOME"
  exit 1
fi

# Run the Java Spring Boot JAR
$JAVA_HOME/bin/java $JAVA_OPTS -jar $JAR_FILE $SPRING_OPTS
```

### **Explanation**:
- **JAVA_OPTS**: Sets **minimum heap (`-Xms`)** and **maximum heap (`-Xmx`)**.
- **SPRING_OPTS**: Sets the **port number** for the Spring Boot app.
- **JAVA_HOME**: If Java isn’t in the default system path, we set it explicitly.
- **Executes**: The `java` command runs the **Spring Boot JAR** with the options provided.

### 2. **Unix (SH) – `.sh` Script**
This is similar to the Linux script, but for older **Unix** systems like AIX, Solaris, or HP-UX, where we might need to adjust the `JAVA_HOME` path or specific shell nuances.

```sh
#!/bin/sh

# Define the JAR file location and Java options (heap size, port)
JAR_FILE="wer-springboot-app.jar"
JAVA_OPTS="-Xms512m -Xmx2048m"
SPRING_OPTS="--server.port=9090"

# Set JAVA_HOME if it's not set (e.g., IBM Java, AIX-specific)
if [ -z "$JAVA_HOME" ]; then
  JAVA_HOME="/usr/java8_64"  # Customize as per wer environment
fi

# Ensure the Java executable exists
if [ ! -x "$JAVA_HOME/bin/java" ]; then
  echo "Java not found at $JAVA_HOME"
  exit 1
fi

# Execute the Java JAR
$JAVA_HOME/bin/java $JAVA_OPTS -jar $JAR_FILE $SPRING_OPTS
```

### **Explanation**:
- Same as the Linux script, but tweaked for older Unix systems.
- **JAR_FILE** is the Spring Boot JAR file.
- **JAVA_OPTS** sets **heap sizes**, and **SPRING_OPTS** defines the port.
- **JAVA_HOME** is used to ensure the correct Java version is called.

### 3. **Windows (BAT) – `.bat` Script**
For Windows, we’ll use a `.bat` script to do the same: run a Java Spring Boot JAR with heap size settings and a specific port.

```bat
@echo off

REM Define the JAR file location and Java options (heap size, port)
set JAR_FILE=wer-springboot-app.jar
set JAVA_OPTS=-Xms256m -Xmx1024m
set SPRING_OPTS=--server.port=8081

REM Set JAVA_HOME if needed (e.g., IBM Java or custom path)
IF NOT DEFINED JAVA_HOME (
    set JAVA_HOME=C:\Program Files\IBM\Java80  REM Update to wer Java path
)

REM Check if Java exists in the specified JAVA_HOME
IF NOT EXIST "%JAVA_HOME%\bin\java.exe" (
    echo Java not found at %JAVA_HOME%
    exit /b 1
)

REM Run the Spring Boot JAR with the Java options
"%JAVA_HOME%\bin\java" %JAVA_OPTS% -jar %JAR_FILE% %SPRING_OPTS%
```

### **Explanation**:
- **JAVA_OPTS** and **SPRING_OPTS** set the heap size and server port, respectively.
- **JAVA_HOME** is set if it’s not already defined, which is useful if we’re using **IBM Java** or any non-default Java location.
- We check if the **java.exe** binary exists in the specified path before running the JAR.

---

### **Common Points**:
- Each script sets **heap sizes** (`-Xms` for minimum heap, `-Xmx` for maximum heap) and a specific **Spring Boot port** (`--server.port`).
- They all ensure that **Java is correctly set up** by checking for the `JAVA_HOME` environment variable. If **Java is in a custom location** (e.g., IBM Java), we adjust the path accordingly.
- The scripts are designed to work in **Linux**, **Unix**, and **Windows** environments, so we should be able to run wer Spring Boot JAR on any of these systems.

### **Customizing**:
- we need to **replace paths** and **heap sizes** depending on wer environment. For instance, change `JAVA_HOME` if we’re using **IBM Java** or a specific version installed in a custom location.
- we can adjust the **heap sizes** based on the memory requirements of wer Spring Boot app.

---









