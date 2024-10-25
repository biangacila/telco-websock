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
        stage('Verify Docker') {
            steps {
                sh 'docker --version'
            }
        }
        stage("Build") {
            steps {
                sh 'mvn --version'
                sh 'docker version'
                echo 'PATH  - $PATH'
                echo 'BUILD_NUMBER  - $env.BUILD_ID'
                echo 'JOB_NAME - $env.JOB_NAME'
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
