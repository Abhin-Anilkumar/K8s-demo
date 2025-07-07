FROM node:21

# Set working directory in container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy rest of the application code
COPY . .

# Expose the app port
EXPOSE 3000

# Command to run the application
CMD ["node", "app.js"]

