FROM ne0nd0g/merlin-base:v1.8.0

# Build multi-arch Docker image and push tagged version to Docker Hub
# > sudo docker buildx build --push --platform linux/amd64,linux/arm64 --tag ne0nd0g/merlin-mythic:v1.0.2 --tag ne0nd0g/merlin-mythic:latest .

WORKDIR /Mythic

# Copy in merlin/Payload_Type/merlin (location where this Docker file is) into the Docker image at /Mythic
COPY [".", "."]

# Allow Go to auto-download a newer toolchain if required by go.mod
ENV GOTOOLCHAIN=auto

# Fetch the latest merlin-agent from the fork at build time.
# ADD from a URL is never cached by Docker, so when the branch has new commits
# the subsequent git clone layer is also rebuilt automatically.
ADD https://api.github.com/repos/Vealending/merlin-agent/git/refs/heads/feature/wsl-command /tmp/agent-ref.json
RUN git clone --branch feature/wsl-command --depth 1 https://github.com/Vealending/merlin-agent.git /Mythic/merlin-agent

# Override the go.mod replace directive to use the freshly cloned source
WORKDIR /Mythic/agent
RUN go mod edit -replace github.com/Ne0nd0g/merlin-agent/v2=/Mythic/merlin-agent
RUN go mod tidy && go mod download

WORKDIR /Mythic/container
RUN go mod download

# Build the Merlin container and save in the /usr/local/bin directory because it is in the PATH environment variable
# Do not save it in the Mythic directory because it is temporary while building and will be lost
RUN go build -o /usr/local/bin/MerlinContainer

CMD /usr/local/bin/MerlinContainer
