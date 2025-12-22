#!/bin/bash

# Telegram Task Manager - Quick Start Script

echo "🚀 Starting Telegram Task Manager..."
echo ""

# Check if .env exists
if [ ! -f .env ]; then
    echo "⚠️  No .env file found. Creating from .env.example..."
    cp .env.example .env
    echo "✅ Created .env file"
    echo ""
    echo "⚠️  IMPORTANT: Edit .env and set your TELEGRAM_BOT_TOKEN before continuing!"
    echo "   Get your bot token from @BotFather on Telegram"
    echo ""
    read -p "Press Enter after you've updated .env, or Ctrl+C to exit..."
fi

# Check if air is installed
if ! command -v air &> /dev/null; then
    echo "⚠️  'air' is not installed. Installing..."
    go install github.com/air-verse/air@latest
    echo "✅ Installed air"
    echo ""
fi

# Install dependencies if needed
if [ ! -d "node_modules" ]; then
    echo "📦 Installing root dependencies..."
    npm install
    echo "✅ Root dependencies installed"
    echo ""
fi

if [ ! -d "mini-app/node_modules" ]; then
    echo "📦 Installing frontend dependencies..."
    cd mini-app && npm install && cd ..
    echo "✅ Frontend dependencies installed"
    echo ""
fi

# Download Go dependencies
echo "📦 Downloading Go dependencies..."
go mod download
echo "✅ Go dependencies ready"
echo ""

# Start development servers
echo "🎯 Starting development servers..."
echo "   - Backend: http://localhost:3001"
echo "   - Frontend: http://localhost:5173"
echo ""
echo "Press Ctrl+C to stop all servers"
echo ""

npm run dev
