pipeline {
    agent any
    environment {
        dockerHome = tool 'myDocker'
        mavenHome = tool 'myMaven'
        goHome = tool 'myGo' // Changed to 'myGo' as per your Jenkins setup
        PATH = "${goHome}/bin:${dockerHome}/bin:${mavenHome}/bin:$PATH"
    }
    stages {
        stage("Checkout") {
            steps {
                sh 'mvn --version'
                sh 'docker version'
                sh 'go version' // Check Go version
                echo "PATH: $PATH"  // Log the PATH
                echo "PATH  - $PATH"
                echo "BUILD_NUMBER  - $env.BUILD_ID"
                echo "BUILD_TAG - $env.BUILD_TAG"
                echo "JOB_NAME - $env.JOB_NAME"
                sh 'cat /etc/os-release'
                echo "Build"
            }
        }
        stage("Build Docker Image") {
            steps {
                sh 'go mod tidy' // Ensure dependencies are installed
                sh 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o programfile .'
                script {
                    docker.build("010309/telco-websocket:${env.BUILD_TAG}")
                }
            }
        }
        stage("Push Docker Image") {
            steps {
                script {
                    docker.image("010309/telco-websocket:${env.BUILD_TAG}").push()
                }
            }
        }
    }
    post {
        always {
            echo "I'm awesome. I run always"
        }
        success {
            echo "I run when you are successful"
        }
        failure {
            echo "I run when you fail"
        }
    }
}


