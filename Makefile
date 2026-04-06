.PHONY: reset restore redis-up redis-down react-install

# Reset an exercise to blank state, backing up the current solution.
# Usage: make reset EX=00-warmup/ex01-copy-by-hand
reset:
	@if [ -z "$(EX)" ]; then echo "Usage: make reset EX=<exercise-path>"; exit 1; fi
	@if [ ! -f "$(EX)/_original/main.go" ]; then echo "No _original/main.go found in $(EX)"; exit 1; fi
	cp $(EX)/main.go $(EX)/main.go.bak
	cp $(EX)/_original/main.go $(EX)/main.go
	@echo "Reset $(EX) — solution backed up as $(EX)/main.go.bak"

# Restore a backed-up solution after resetting.
# Usage: make restore EX=00-warmup/ex01-copy-by-hand
restore:
	@if [ -z "$(EX)" ]; then echo "Usage: make restore EX=<exercise-path>"; exit 1; fi
	@if [ ! -f "$(EX)/main.go.bak" ]; then echo "No main.go.bak found in $(EX)"; exit 1; fi
	cp $(EX)/main.go.bak $(EX)/main.go
	@echo "Restored solution for $(EX)"

# Start Redis in the background (required for 05-redis exercises).
redis-up:
	docker compose up -d redis

# Stop Redis and remove the container.
redis-down:
	docker compose down

# Install dependencies for the React + TypeScript optional block.
react-install:
	cd optional/react-ts-vite && npm install
