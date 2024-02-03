name: Docker Image CI

on:
  push:
    branches: [ main ]
    paths:
      - 'game/**'
      - 'Dockerfile'
  # You can also trigger on pull_request, release, or other events

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
    - name: Check out repository
      uses: actions/checkout@v2

    - name: Build Docker image
      run: docker build -t myapp:latest .

    - name: Push Docker image to Registry
      if: github.ref == 'refs/heads/main' && github.event_name == 'push'
      run: |
        echo "${{ secrets.DOCKER_PASSWORD }}" | docker login -u "${{ secrets.DOCKER_USERNAME }}" --password-stdin
        docker push myapp:latest
