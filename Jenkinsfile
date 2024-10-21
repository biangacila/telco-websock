pipeline {
    agent {
        docker {
            image 'ubuntu:22.04' // Specify your Docker image here
            args '-v /var/run/docker.sock:/var/run/docker.sock' // Pass any arguments you need
        }
    }
    stages {
        stage('Verify Docker') {
            steps {
                sh 'docker --version'
            }
        }
        stage("Build") {
            steps {
                sh 'cat /etc/os-release'
                echo "Build"
            }
        }
        stage("Test") {
            steps {
                echo "Test"
            }
        }
        stage("Integration Test") {
            steps {
                echo "Test Integration"
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
