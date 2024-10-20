pipeline{
    //agent any
    agent {docker{ image 'ubuntu:22.04'}}
    stages{
        stage("Build"){
            steps{
                sh 'cat /etc/os-release'
                echo "Build"
            }
        }
        stage("Test"){
            steps{
                echo "Test"
            }
        }
        stage("Integration Test"){
            steps{
                echo "Test Integration"
            }
        }
    }
    post{
        always{
            echo "Im awesome. I run always"
        }
        success{
            echo "I run when you are success"
        }
        failure{
            echo "I run when you  fail"
        }
    }
}