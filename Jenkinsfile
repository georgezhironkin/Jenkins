pipeline {
    agent any
    environment {
        PATH = "/var/jenkins_home/go/bin:${env.PATH}"
    }
    stages {
        stage('Unit Tests') {
            steps {
                sh 'go test -v -run="TestHandler|TestHandler2"'
            }
        }
        stage('Integration Test') {
            steps {
                sh 'go test -v -run="TestIntegration"'
            }
        }
    }
}
