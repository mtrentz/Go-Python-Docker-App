FROM python:3.8-slim-buster
WORKDIR /app
COPY server ./server
WORKDIR /app/server
RUN pip install --no-cache-dir -r requirements.txt
CMD ["python", "server.py"]
