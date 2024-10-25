pipeline {
    /* agent {
        docker {
            image 'docker:20.10-dind'
            args '--privileged' // Needed for Docker in Docker
        }
    } */
    agent any
    environment {
        dockerHome = tool 'myDocker'
        mavenHome = tool 'myMaven'
        PATH = "$dockerHome/bin:$mavenHome/bin:$PATH"
    }
    stages {
        stage("Checkout") {
            steps {
                sh 'mvn --version'
                sh 'docker version'
                echo 'PATH  - $PATH'
                echo 'BUILD_NUMBER  - $env.BUILD_ID'
                echo 'BUILD_TAG - $env.BUILD_TAG'
                echo 'JOB_NAME - $env.JOB_NAME'
                sh 'cat /etc/os-release'
                echo "Build"
            }
        }
        stage("Build Docker Image") {
            steps {
                //docker build -t 010309/telco-websocket:$env.BUILD_TAG
                sh 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o programfile .'
                script{
                    docker.build("010309/telco-websocket:${env.BUILD_TAG}")
                }
            }
        }
        stage("Push Docker Image") {
            steps {
                sh "mvn test"
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
