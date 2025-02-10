# CodeSync 🚀  
**Real-Time Collaborative Code Editor**  
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A lightweight, open-source alternative to VS Code Live Share. Edit code, share terminals, and voice chat in real-time.

🌟 **Star this repo** to support open-source collaboration tools!

---

## Features ✨
- **Real-Time Code Editing**: Collaborate like Google Docs for code.
- **Shared Terminal**: Run code in isolated Docker containers.
- **Voice Chat**: WebRTC-powered audio communication.
- **GitHub OAuth**: Secure authentication with permissions.
- **Offline-First**: CRDTs for conflict-free sync (coming soon).

---

## Tech Stack 💻  
| **Frontend**       | **Backend**       | **Infra**               |
|--------------------|-------------------|-------------------------|
| Next.js + React    | Go (Gorilla)      | Docker + Docker Compose |
| Monaco Editor      | Redis (pub/sub)   | PostgreSQL              |
| XTerm.js           | WebSocket         | GitHub Actions          |
| WebRTC             | JWT Auth          | Nginx                   |

---

## Getting Started 🛠️  
### Prerequisites  
- Go 1.21+ (Backend)
- Node.js 18+ (Frontend)
- Docker + Docker Compose        
