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
                echo 'JOB_NAME - $env.JOB_NAME'
                sh 'cat /etc/os-release'
                echo "Build"
            }
        }
        stage("Compile") {
            steps {
                sh "mvn clean compile"
            }
        }
        stage("Test") {
            steps {
                sh "mvn test"
            }
        }
        stage("Integration Test") {
            steps {
                sh "mvn failsafe:integration-test failsafe:verify"
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
