# ---- Build frontend ----
FROM node:20 AS frontend-builder

WORKDIR /app/frontend
COPY frontend/package*.json ./
RUN npm install

COPY frontend .
RUN npm run build


# ---- Build Go backend ----
FROM golang:1.22 AS backend-builder

WORKDIR /app
COPY backend/go.mod backend/go.sum ./backend/
WORKDIR /app/backend
RUN go mod download

COPY backend .
RUN CGO_ENABLED=0 GOOS=linux go build -o server


# ---- Final runtime image ----
FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=backend-builder /app/backend/server .
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist

EXPOSE 8080

CMD ["./server"]