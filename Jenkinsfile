pipeline{
    agent any
    stages{
        stage("Build"){
            steps{
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