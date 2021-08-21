FROM python:3.8-slim-buster
WORKDIR /app
COPY server ./server
WORKDIR /app/server
CMD ["python", "server.py"]
