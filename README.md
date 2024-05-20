# Infra Assignment


## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed [Docker](https://docs.docker.com/get-docker/).

## Getting Started

### 1. Obtain a GitHub Personal Access Token (PAT)

To access private repositories and perform actions on behalf of your account, you'll need a GitHub Personal Access Token.

1. Go to your GitHub account settings.
2. Navigate to **Developer settings** > **Personal access tokens**.
3. Click **Generate new token**. Use Classic token.
4. Give your token a descriptive name.
5. Select the permissions  ```delete:packages, repo, write:packages```
6. Click **Generate token** and copy the token.

### 2. Login to GitHub

Using your Personal Access Token, log in to GitHub from your terminal:

```bash
echo "your_github_pat" | docker login ghcr.io -u your_github_username --password-stdin
```
### 3. Run the Project

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/tuky191/happychain-assignment.git
cd happychain-assignment
docker-compose up -d
```

### 4. Demo

#### Access Blockscout
http://localhost

#### Check output of demo contaier
```bash
docker logs demo -f
```
