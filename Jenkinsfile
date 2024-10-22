pipeline {
    agent {
        docker {
            image 'docker:20.10-dind'
            args '--privileged' // Needed for Docker in Docker
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
