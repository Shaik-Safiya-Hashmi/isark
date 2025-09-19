
# 🌿 Isark – Blockchain-based Ayurvedic Supply Chain  

## 📌 Project Description  
**Isark** is a blockchain-based traceability platform built on **Hyperledger Fabric** to ensure authenticity, transparency, and quality in the Ayurvedic product supply chain.  

The system records every step of an Ayurvedic herb’s journey — from **farm harvest → market aggregation → scientific processing → lab testing → manufacturing → packaging → distribution** — in a tamper-proof blockchain ledger.  

A **QR code** is generated at the end of the process (during packaging), allowing customers to scan and verify the **complete journey** of their product.  

---

## ✨ Features  
- ✅ **Blockchain-secured traceability** (Hyperledger Fabric)  
- ✅ **Smart contract (chaincode)** for recording product journey  
- ✅ **CRUD operations** for products and journey stages  
- ✅ **Frontend** (React/Next.js) for customers to scan QR & view journey  
- ✅ **Backend** (Node.js/Express) for APIs to interact with blockchain  
- ✅ **QR Code integration** for product verification  
- ✅ **Tamper-proof & transparent records**  

---

## 🏗️ Project Structure  

```
Isark/
 ├── fabric-samples/        # Cloned Hyperledger Fabric samples (test-network, etc.)
 │   ├── test-network/      # Network scripts to start Fabric
 │   ├── chaincode/         
 │   │   └── product_journey/   # Custom chaincode (Go)
 │
 ├── backend/               # REST APIs (Node.js/Express) to interact with Fabric
 ├── frontend/              # User interface (React/Next.js) for QR scanning & display
 ├── README.md              # Project documentation
```

---

## ⚡ Tech Stack  
- **Blockchain** → Hyperledger Fabric  
- **Smart Contracts** → Go (Fabric Contract API)  
- **Backend** → Node.js + Express  
- **Frontend** → React/Next.js  
- **Database** → Blockchain ledger (Fabric world state)  
- **QR Codes** → `qrcode` npm package  

---

## 🚀 Setup Instructions  

### 1️⃣ Clone Repository  
```bash
git clone https://github.com/your-username/isark.git
cd isark
```

### 2️⃣ Start Blockchain Network  
```bash
cd fabric-samples/test-network
./network.sh up createChannel -c mychannel -ca
```

### 3️⃣ Deploy Smart Contract  
```bash
./network.sh deployCC -ccn product_journey -ccp ../chaincode/product_journey -ccl go
```

### 4️⃣ Test Chaincode  
Add product:  
```bash
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile $PWD/organizations/ordererOrganizations/example.com/tlsca/tlsca.example.com-cert.pem -C mychannel -n product_journey --peerAddresses localhost:7051 --tlsRootCertFiles $CORE_PEER_TLS_ROOTCERT_FILE --peerAddresses localhost:9051 --tlsRootCertFiles $PWD/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt -c '{"Args":["AddProduct","P001","Ashwagandha","Farm Harvest","Neemuch, MP","2025-09-10","350kg harvested"]}'
```

Query product:  
```bash
peer chaincode query -C mychannel -n product_journey -c '{"Args":["GetProduct","P001"]}'
```

---

## 🔗 Example Blockchain Record  

```json
{
  "id": "P001",
  "name": "Ashwagandha",
  "stage": "Farm Harvest",
  "location": "Neemuch, MP",
  "timestamp": "2025-09-10",
  "details": "350kg harvested"
}
```

---

## 📦 Future Enhancements  
- Integrate **backend APIs** for automation  
- Connect **frontend UI** for QR scanning  
- Add **multi-stage product journey** support  
- Deploy on **cloud infrastructure (Azure/AWS/GCP)**  

---

## 👨‍💻 Team & Contributions  
This project is developed as part of **Smart India Hackathon (SIH)** to promote **trust and transparency in Ayurveda supply chains**.  
