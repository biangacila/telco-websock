pipeline {
    agent {
        docker {
            image 'ubuntu:22.04'
            label 'docker' // Optional: specify a node label if needed
        }
    }
    stages {
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
