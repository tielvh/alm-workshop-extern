# Use the latest dev-bookworm base image for better performance
FROM mcr.microsoft.com/devcontainers/base:dev-bookworm

# Set environment variables to avoid interactive prompts
ENV DEBIAN_FRONTEND=noninteractive

# Install required dependencies
RUN apt-get update && apt-get install -y \
    curl wget git sudo unzip tar bash-completion \
    ca-certificates lsb-release jq iproute2 docker.io \
    && rm -rf /var/lib/apt/lists/*

# Install OpenShift CLI (oc)
RUN curl -L https://mirror.openshift.com/pub/openshift-v4/clients/ocp/stable/openshift-client-linux.tar.gz | \
    tar -xz -C /usr/local/bin && chmod +x /usr/local/bin/oc

# Install Go (latest stable)
RUN curl -fsSL https://golang.org/dl/go1.23.6.linux-amd64.tar.gz | \
    tar -C /usr/local -xz && ln -s /usr/local/go/bin/go /usr/local/bin/go

# Install Kubernetes CLI tools
RUN curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl" && \
    chmod +x kubectl && mv kubectl /usr/local/bin/

# Install Helm
RUN curl -fsSL https://get.helm.sh/helm-v3.11.2-linux-amd64.tar.gz | tar -xz -C /usr/local/bin --strip-components=1

# Install Minikube
RUN curl -Lo minikube https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64 && \
    chmod +x minikube && mv minikube /usr/local/bin/

# Set up Go path (for VSCode and dev tools)
ENV GOPATH=/go
ENV PATH=$GOPATH/bin:/usr/local/go/bin:$PATH

# Install VSCode extensions in correct directories
RUN code --install-extension redhat.vscode-yaml && \
    code --install-extension eamodio.gitlens && \
    code --install-extension 766b.go-outliner && \
    code --install-extension golang.go && \
    code --install-extension premparihar.gotestexplorer || true

# Default shell
CMD ["bash"]
