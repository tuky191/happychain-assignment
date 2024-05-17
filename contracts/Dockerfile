FROM node:18-alpine

# Install dependencies
RUN apk add --no-cache --virtual .gyp python3 make g++ 

# Set the working directory
WORKDIR /app

# Copy package.json 
COPY package.json ./

# Install project dependencies (including Truffle locally)
RUN npm install

# Copy the rest of the project files
COPY . .

# Default command
CMD ["npx", "truffle", "migrate", "--network", "development"]
