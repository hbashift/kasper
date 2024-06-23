FROM golang:1.21.5
# Install Python
RUN apt-get update && apt-get install -y python3 python3-pip python3-venv
WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
# Create a virtual environment
RUN python3 -m venv venv

# Activate the virtual environment and install Python dependencies
RUN . venv/bin/activate && pip install -r requirements.txt
RUN chmod 777 venv/bin/activate

RUN go build -o ./bin/server ./cmd/kasper/main.go

EXPOSE 8080

CMD ["./bin/server"]