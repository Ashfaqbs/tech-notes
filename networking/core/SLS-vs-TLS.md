## Working of SLS vs TLS

### Let me give you a proper story on how SLS and TLS work. A device can be a machine, a server, or even a phone—anything that communicates securely over the network. Here’s the explanation with a relatable scenario.

---

### **Scenario: Two Friends Sharing Secrets Safely**
Imagine two friends, **Ali** and **Rahul**, who want to share secrets securely over the internet. Depending on the situation, they use either **SLS** or **TLS**.

---

### **SLS Story: A Private Meeting in a Quiet Room**
- **Step 1: Secure but Less Standardized Communication**  
  Ali wants to share some personal secrets with Rahul. He chooses **SLS** (Session Layer Security) because it offers a secure conversation without a lot of formalities.

- **Step 2: Creating a Secure Session**  
  Ali and Rahul decide to meet in a quiet, private room. They check that no one is listening before talking.  
  - This is like **SLS establishing a secure session** without involving complex certificates or public key infrastructure (PKI).

- **Step 3: Sharing the Secrets**  
  Ali whispers the secrets, and Rahul listens carefully, ensuring no eavesdroppers. If either of them feels unsafe, they immediately stop the conversation.  
  - In **SLS**, the session is secured but often application-specific, meaning different apps might handle security differently.

- **Step 4: Trust Within the Session**  
  As long as the room is secure, Ali and Rahul trust the conversation. However, there's no external validation or certificate authority verifying the room's safety.  
  - This makes SLS **less standardized** compared to TLS.

---

### **TLS Story: A Secure Phone Call with Authentication**
- **Step 1: Highly Secure and Standard Communication**  
  Now, Ali wants to share **sensitive business documents** with Rahul over the internet. He chooses **TLS** because it provides a high standard of security and trust.

- **Step 2: Establishing a Secure Connection**  
  Before sharing any secrets, Ali's phone verifies Rahul's identity through a secure passcode (certificate validation).  
  - This is like the **TLS handshake**. The server presents a certificate to prove its identity.

- **Step 3: Encrypted Communication**  
  Once the connection is established, their conversation is encrypted end-to-end, ensuring no one can intercept it.  
  - TLS uses symmetric encryption for fast and secure communication after the initial handshake.

- **Step 4: Guaranteed Trust and Security**  
  If the certificate is invalid, the call doesn't proceed, protecting Ali from potential fraud.  
  - This ensures **integrity, confidentiality, and authenticity** in TLS.

---

### **Summary of the Story**
- SLS is like a **private conversation in a room**, secure but with fewer formal verification steps.
- TLS is like a **secure phone call with identity verification**, offering stronger, standardized security for internet communications.

---

## A **handshake** is the process where two devices agree on how to communicate securely.

---

### **What is a Handshake?**
- Think of a handshake as **setting up a secret code** before talking.
- Both devices agree on encryption methods and keys during this handshake.
- This ensures the communication remains private and secure.

---

### **What Happens During a TLS Handshake?**
1. **Client Hello**: Ali asks Rahul, "Are you ready to talk securely?" The client sends supported encryption methods.  
2. **Server Hello**: Rahul replies, "Yes, I am. Here’s my certificate." The server responds with its certificate and encryption choice.  
3. **Key Exchange**: They agree on a symmetric key to use for the conversation.
4. **Secure Session**: All communication after this is encrypted with the shared key.

---

### **Analogy: Agreeing on a Secret Code**
- Imagine two friends agreeing on a secret code before discussing sensitive topics.  
- Once the code is set, even if someone listens, they won't understand the conversation.

---

## What other concepts relate to SLS and TLS?

1. **SSL (Secure Sockets Layer)** - TLS evolved from SSL, providing better security.  
2. **HTTPS (HyperText Transfer Protocol Secure)** - Websites use TLS to secure communications.  
3. **Certificates and CAs (Certificate Authorities)** - These validate a server's identity in TLS.

---

